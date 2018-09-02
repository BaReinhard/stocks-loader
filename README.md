# Stock Loader

### Requirements

1. [AlphaVantage API Key](https://www.alphavantage.co/support/#api-key)
2. Google Cloud Service Account JSON with Datastore Scope
3. ENV VARS:
   - PROJECT_ID: YOUR_PROJECT_ID
   - GOOGLE_APPLICATION_CREDENTIALS: PATH TO YOUR Service Account JSON File
   - API_TOKEN: SET TO YOUR API KEY FROM STEP 1

**Description**:
This can easily be used to get a list of stocks and store them in datastore for further processing at a later time.
