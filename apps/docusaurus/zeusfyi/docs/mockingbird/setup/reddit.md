---
sidebar_position: 3
displayed_sidebar: mockingbird
---

# Reddit

## How to Integrate Reddit API Keys

You'll need to setup a Reddit app to get the API keys. Here's how you can do it:
https://www.reddit.com/wiki/api/

## Token Setup

You'll need to manually add the reddit secrets from your API keys you've obtained from Reddit.
They need to match the secret names listed below.

![Screen](https://github.com/zeus-fyi/zeus/assets/17446735/50a1170a-13cb-475b-80ec-11347b9cdf2a)

## Generated Reddit Routing Table for API Calls

You should see a routing group named reddit-{YOUR_REDDIT_USERNAME} in the routing table.

## Verify Reddit API Calls

In the search tab on the main AI page, you can verify the API calls by using the following route path:

Currently only PUT, POST, and GET requests are supported.

Example Route: `api/v1/me`

Reddit API Documentation: https://developer.twitter.com/en/docs/twitter-api

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/6d1e83ea-8f23-47d2-b29a-ff3811e6deeb)
