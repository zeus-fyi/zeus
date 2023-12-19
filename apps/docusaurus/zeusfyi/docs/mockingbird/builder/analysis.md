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