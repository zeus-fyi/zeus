---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Tweetflow

## Overview

This guide is crafted to walk you through the process of developing a system that streamlines your social media
interactions, specifically tailored for Twitter using Zeusfyi’s Mockingbird AI system.

Our goal is to enable you to efficiently analyze and respond to your latest saved bookmarks on Twitter. The workflow
we’re constructing will leverage cutting-edge AI to curate context-aware tweet replies automatically. Here’s an overview
of what this guide will cover:

- Automated Analysis: The workflow begins by analyzing your most recent Twitter bookmark, extracting key themes and
  sentiments to inform the response generation.
- Reply Generation: Using sophisticated AI, the system will then compose a reply that is relevant and engaging, ensuring
  the conversation flows naturally.
- Auto-Evaluation: Before any tweet reaches you, it undergoes an automated evaluation process. This step assesses the
  reply against pre-established quality criteria and metrics to ensure appropriateness and alignment with your digital
  voice.
- Approval Mechanism: Once the tweet passes the auto-evaluation, it is forwarded to you for final review. This guide
  will explain how to approve or reject the suggested content, giving you complete control over what gets posted.
- Execution: Upon your approval, the tweet will be automatically posted, maintaining timely and relevant engagement with
  your audience.

- By following this tutorial, you will set up a workflow that not only speeds up your response time on Twitter but also
  ensures that each interaction is thoughtful and brand-consistent. Let’s embark on this journey to enhance your Twitter
  presence with automation and precision.

## Step 1: Automated Twitter Auth & Routing Table Setup

Connect Twitter: Start by establishing authorization and a routing table for Twitter API calls.
Use the “Automated Twitter Auth & Routing Table Setup” to create a routing group named twitter-{YOUR_TWITTER_HANDLE}.
This will also generate a bearer token and save it in the platform's secret manager.
This section is in the indexer tab on the main AI page. It will redirect to Twitter for you to authorize your account
for usage.

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/25300931-cfd3-466d-b898-15bce758c50f)

## Step 2: Inputs/Outputs Configuration

API via Load Balancer: Configure an API endpoint that can be used by a load balancer to aggregate data from different
routes.
This is for setting up general retrieval processes that can quickly start. In the below example. You’ll need to setup
the below two stages.

### Input Retrieval: Getting Your Latest Tweet Bookmark

- Route: `users/:id/bookmarks?max_results=1`

  ![Screenshot 2024-02-04 at 6 59 39 PM](https://github.com/zeus-fyi/zeus/assets/17446735/eba87817-1f9d-4260-869a-43dbf4f638cb)

### Reply via API: Sending an API request to Twitter to post a Tweet.

![Screenshot 2024-02-04 at 6 59 39 PM](https://github.com/zeus-fyi/zeus/assets/17446735/eba87817-1f9d-4260-869a-43dbf4f638cb)

## Step 3: Defining a JSON Schema

Define the Twitter post request schema
Add Fields: Add field names, types, and descriptions necessary for the Twitter API (e.g., text, in_reply_to_tweet_id).

Note: Actual API request needed for creating tweet replies

```json
{
  "text": "tweet body text",
  "reply": {
    "in_reply_to_tweet_id": "tweet_id_value"
  }
}
```

Simplified schema tweet API request we’re building

```json
{
  "text": "tweet body text",
  "in_reply_to_tweet_id": "in_reply_to_tweet_id"
}
```

We can skip the nested object field for this specific request since we provide an embedded transformer which
converts to the above format when API requests are sent to the twitter API from a human-in-the-loop approved tweet.

![Screenshot 2024-02-04 at 7 01 20 PM](https://github.com/zeus-fyi/zeus/assets/17446735/df3548a6-35b8-445c-a155-0376856f0fc9)

## Step 4: Creating the AI Task for Tweet Response

Add an analysis task to generate a tweet response based on the latest bookmarked tweet.

![Screenshot 2024-02-M](https://github.com/zeus-fyi/zeus/assets/17446735/64a741bc-72fe-4eb4-90a1-0d89be0afa15)

## Step 5: Creating the Trigger Tweet Response API Call

You attach your previous API Input/Output stage for it to route your API request after it’s been approved. This also
only triggers an action for approval if all your eval info stages pass.

![Screenshot 2024-02-04 at 7 02 ](https://github.com/zeus-fyi/zeus/assets/17446735/8d39c309-581c-4a0e-9c4d-a1ae0dc72215)

![Screenshot 2024-02-04 at 7 02 56PM](https://github.com/zeus-fyi/zeus/assets/17446735/9f3f07d0-dffb-448c-ab2e-87867b87b6f4)

## Step 6: Creating the Eval Task for Tweet Response

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/a93e1279-6d4e-4a30-86ae-3706cda56e5f)

You’ll attach the JSON schema we created earlier and then set the eval criteria on the fields. In this example it
verifies that the text body meets the 280 character limit for twitter and that the tweet _id being replied to is not
empty. You then attach the trigger to execute on Eval completion.

## Step 7: Creating the Workflow

- Attach bookmark retrieval
- Attach tweet response analysis task
- Attach eval to your tweet response analysis task
- Save workflow

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/c1944550-3fad-4f79-afc5-55db9887c9a2)

![ScreM](https://github.com/zeus-fyi/zeus/assets/17446735/07e041fa-e344-434a-9afe-c2b385463758)

## Step 8: Run Execute Workflow

Since we just want to execute by cycles. Just select 1 cycle of runtime and then press start.
Temporal Workflow Integration: Use the durable execution feature of Temporal workflows to chain tasks together, running
them on a schedule or manually.

![Screensh](https://github.com/zeus-fyi/zeus/assets/17446735/4142dfa7-57f1-4a9d-8d5d-d65c49239072)

## Run Inspection

You can view the results of your eval stages and tasks in the Run section. If your eval passed, you’ll be able to see
the trigger action in the Actions tab.

![Sc P](https://github.com/zeus-fyi/zeus/assets/17446735/c20e47b1-d415-441c-b872-55a4b3097413)

## Step 9: Inspect and Approve AI Driven API Call

Now you can select Approve or Reject for the AI generated Tweet. If you like it, press approve and it will execute the
POST tweet request using the JSON body output from your Auto-Eval stage.

![Screenshot 2024-02-04 at 7 05 54 PM](https://github.com/zeus-fyi/zeus/assets/17446735/f8c2911a-6e47-42c9-b95c-a9d374017fa9)