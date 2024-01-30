---
sidebar_position: 3
displayed_sidebar: mockingbird
---

# Analysis

Token overflow strategy will determine how the AI will handle requests that are projected to exceed the maximum token
length for the model you select, or has returned a result with that error. Deduce will chunk your analysis into smaller
pieces and aggregate them into a final analysis result. Truncate will simply truncate the request to the maximum token
length it can support. If you set the max tokens field greater than 0, it becomes the maximum number of tokens to spend
per task request.

![zzz](https://github.com/zeus-fyi/zeus/assets/17446735/8340c438-5b19-4684-9b19-1cb0c973a5fe)

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/da52e86d-d24b-4ee0-9f49-37e2c3b6e7a8)

### Example Run: Analysis->Aggregation Flows

This shows an analysis stage feeding the aggregation stage. Notice how the Completion Choices output from the
analysis stage are then used in the aggregation stage as the prompt body, which is how the aggregation stages work.
By default it will use the last analysis stage in the workflow, relative to the aggregation stage, and
for multiple analysis stages it will aggregate once every analysis stage has completed at least once
and then aggregate all the results from the analysis stages that are within the time window since last aggregation.

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/da52e86d-d24b-4ee0-9f49-37e2c3b6e7a8)
