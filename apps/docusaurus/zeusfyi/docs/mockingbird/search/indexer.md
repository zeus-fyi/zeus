---
sidebar_position: 2
displayed_sidebar: mockingbird
---

# Indexer

The indexers will use your search parameters and accounts to create a time series of the fetched data every few minutes.
It only requests the latest data since the last time it was run, and should not fetch the same data twice within the
same search group. The indexer will also store the data in the database for archiving and retrieval.

## Twitter

![Screen](https://github.com/zeus-fyi/zeus/assets/17446735/2b32bdaa-216d-4379-8006-679cc26bd2c6)

Set the Twitter API credentials in the platform secrets to start indexing tweets,
then set the platform search query using the v2 Twitter search API format.

### Platform Search Query

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/067a6fc7-8b0c-4687-8c79-747c9f09ce58)

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/d862eaac-2462-4ce2-b0ca-09f97df714fc)

![ScrM](https://github.com/zeus-fyi/zeus/assets/17446735/a411ed5b-9981-4b72-a819-50f8f25adf2b)

![Screensh2PM](https://github.com/zeus-fyi/zeus/assets/17446735/87528c34-8b1d-4813-9af8-c062eabbbc6b)

![ScrM](https://github.com/zeus-fyi/zeus/assets/17446735/6a11c27e-a28b-417a-a34a-8a677df07b27)

Full
Documentation: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-recent

## Reddit

![SScreM](https://github.com/zeus-fyi/zeus/assets/17446735/7026105b-f2bc-462e-8e7b-9c9ddd99d062)

Set the Reddit API OAuth2 credentials, user, and password, in the platform secrets to start indexing subreddits.
Also set the routing table as well for reddit retrievals/api calls. If you don't want to use your primary reddit
account, that's fine, you can create a new account and use that for the API credentials as well.

![S](https://github.com/zeus-fyi/zeus/assets/17446735/d2ca5fe5-b191-42fc-a769-ee4f9ca4aaf3)

## Activate / Deactivate Indexing

![ScrM](https://github.com/zeus-fyi/zeus/assets/17446735/b9bb2c6e-58b3-467e-b45a-c59a7bf067e2)

Select the relevant search query row, and click the Activate/Deactivate button to start/stop indexing data.