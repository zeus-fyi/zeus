---
sidebar_position: 2
displayed_sidebar: mockingbird
---

# API Approvals

## Overview

The Actions tab is an essential feature for overseeing and managing API trigger calls within your workflow.
It serves as a control panel where users can approve or reject pending actions, as well as view the historical
results and responses of past actions. This ensures that all external communications via API calls are vetted by human
oversight.

## Approving and Rejecting API Triggers

### Action Approval Process

#### Review Request Summary:

Examine the details of the API call request, including the Workflow Result ID, Approval ID, and the Request Summary.

#### Approve or Reject:

Make an informed decision to approve or reject the trigger action. Approving an action will execute the API call, while
rejecting it will prevent the call from being made.

### Update Status:

Once an action is approved or rejected, its status is updated to reflect the new state.

## Action Buttons

### Approve:

Click this button to authorize the API call.

### Reject:

Click this button to decline the API call.

## Historical Actions

### Viewing Past Actions

### Access Historical Data:

The lower section of the Actions tab displays a history of all API trigger actions that have been taken.

Analyse Past Decisions: Review the Approval ID, Request Summary, Final State, Updated At, and the full list of requests
and responses to understand the outcome of past decisions.

### Historical Insights

Gain insights into the decision-making process over time.
Assess the consistency and accuracy of past approvals or rejections.

### Example of the Actions Tab

    Trigger Name: social-media
    Trigger Group: social-media-approvals  
    Trigger Env: social-media-engagement

Action Details:

- Workflow Result ID: 170665375241213000
- Approval ID: 1706653756335363475
- Approval State: Pending
- Request Summary: Requesting approval for trigger action [...]

Historical Actions:

- Approval ID: 170664805335778953
- Request Summary: Finished approval for trigger action
- Final State: Finished
- Updated At: 1/30/2024, 10:00:05 PM

![Scre](https://github.com/zeus-fyi/zeus/assets/17446735/c1451545-4ea7-4c32-9415-01d079b43b39)
