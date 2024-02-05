---
sidebar_position: 2
displayed_sidebar: mockingbird
---

# Twitter

## How to Integrate Twitter API Keys

Start by establishing authorization and a routing table for Twitter API calls.
Use the “Automated Twitter Auth & Routing Table Setup” to create a routing group named twitter-{YOUR_TWITTER_HANDLE}.
This will also generate a bearer token and save it in the platform's secret manager.
This section is in the indexer tab on the main AI page. It will redirect to Twitter for you to authorize your account
for usage.

![Scre](https://github.com/zeus-fyi/zeus/assets/17446735/13e70734-97f5-4c41-890a-5a0fb103eb8a)

## Oauth Flow

It will redirect to Twitter for you to authorize your account for usage.

![Scre](https://github.com/zeus-fyi/zeus/assets/17446735/4e364665-e028-4b27-b6d6-ff5cda1c7992)

## Generated Table and Route

You should see a routing group named twitter-{YOUR_TWITTER_HANDLE} in the routing table.

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/1d6d50c5-ce70-4e33-a859-909d2c0274a9)

## OAuth2 Token Refresh

We currently do not refresh the token automatically, but you can refresh it manually by pressing connect twitter again,
and repeating all the prior steps. We are working on a solution to automate this process.

## Verify Twitter API Calls

In the search tab on the main AI page, you can verify the API calls by using the following route path:

Example Route: `api/v1/me`

Twitter API Documentation: https://developer.twitter.com/en/docs/twitter-api

![Sc](https://github.com/zeus-fyi/zeus/assets/17446735/5839a764-8aec-4cf5-9313-1615a3ebfe5a)
