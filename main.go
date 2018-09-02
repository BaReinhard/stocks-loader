package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/bareinhard/alpha"
)

func main() {
	symbols := []string{"SNAP", "MSFT"}
	for _, val := range symbols {
		opts := alpha.Options{
			Function:   "TIME_SERIES_DAILY",
			Symbol:     val,
			OutputSize: "compact",

			APIKey: os.Getenv("API_TOKEN"),
		}
		client := alpha.NewClient(opts, &http.Client{})
		api, err := client.Get()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		ctx := context.Background()
		for _, v := range api.TimeSeries {
			fmt.Printf("Symbol: %v\n", val)
			fmt.Printf("Date: %v\n", v.Date)
			fmt.Printf("High: %v\n", v.High)
			fmt.Printf("Low: %v\n", v.Low)
			fmt.Printf("Close: %v\n", v.Close)
			fmt.Printf("Open: %v\n", v.Open)
			fmt.Printf("Volume: %v\n\n", v.Volume)
			addData(ctx, v, val)
		}
	}
}

func addData(ctx context.Context, pload alpha.TimeSeriesDaily, symbol string) error {
	client, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	key := datastore.NameKey("Stocks", fmt.Sprintf("%s-%s", symbol, pload.Date), nil)
	exists := checkData(ctx, key, client)
	if !exists {
		_, err = client.Put(ctx, key, &pload)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Already Exists in Datastore")
	}
	return nil
}
func checkData(ctx context.Context, key *datastore.Key, client *datastore.Client) bool {
	var p alpha.TimeSeriesDaily
	err := client.Get(ctx, key, &p)
	if err != nil && !strings.Contains(err.Error(), "no such entity") {
		fmt.Printf("Error Checking Data: %v", err)
		return false
	} else if err != nil && strings.Contains(err.Error(), "no such entity") {
		return false
	}
	if p.Date == "" {
		return false
	}
	return true
}
