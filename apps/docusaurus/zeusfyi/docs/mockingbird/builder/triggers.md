---
sidebar_position: 7
displayed_sidebar: mockingbird
---

# Triggers and Actions

## Intro

The Triggers and Actions system is a critical component of our automation framework.
It links evaluation functions to API triggers to streamline the execution of AI-driven workflows.
Upon evaluation, this system generates a request payload for user approval or rejection.
If approved, the corresponding API call is executed. Rejected requests are archived and not executed.
This mechanism ensures that automated decisions are overseen and validated by human judgment when necessary.

![Scr](https://github.com/zeus-fyi/zeus/assets/17446735/893cb682-b82d-4cbc-9cdf-742e41f7142b)

## Trigger Action Procedures

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/f8011006-e92c-413a-9267-ea69165b91e9)

### Setting Up Triggers

- **Define a New Trigger**: Assign a name and group to a new trigger, identify the trigger source, and specify the
  trigger action.
- **Link to Eval Function**: Connect the trigger to an eval function that will analyze AI outputs to decide on API call
  initiation.

### API Trigger Settings

- **Retry Mechanisms**: Configure the maximum retry attempts for API calls and set up backoff coefficients for retry
  intervals.
- **Save Trigger**: Save the configured trigger so it can start listening for outputs from the linked eval function.

## Evaluation Functions and API Calls

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/f8011006-e92c-413a-9267-ea69165b91e9)

### Connecting Eval Functions

- **Select the Eval Function**: Pick the appropriate evaluation function that will provide criteria for API call
  triggers.

### Automated Payload Creation

- **Payload Generation**: Once the eval function criteria are met, the system automatically generates an API call
  request payload.

## Actions Menu

![Scree](https://github.com/zeus-fyi/zeus/assets/17446735/8f340b39-5c42-4813-9ca2-7e94334cafdf)

### Review and Decision

- **Approve/Reject Payloads**: Use the actions menu to review generated payloads and make decisions to approve or reject
  them.

### Historical Tracking

- **Action Records**: The system maintains a log of all actions taken, which can be accessed for review or auditing
  purposes.

## Conclusion

This triggers and actions framework empowers users to have final control over automated processes, ensuring that every
action taken by the AI is deliberate and in line with organizational standards.