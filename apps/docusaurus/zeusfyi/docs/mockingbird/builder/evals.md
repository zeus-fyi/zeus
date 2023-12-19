---
sidebar_position: 6
displayed_sidebar: mockingbird
---

# Evals

### Pre-Release

You can only save and view evals at the moment. We are doing QA and some final development still on the public evals
feature
integration and will be releasing it in stages this week.

This allows you to setup scoring rules for AI system outputs that set metrics for the AI to use in its decision
making process. For metric array types, the comparison value returns the true/false if every array element item passes
the comparison eval test.

![ScreM](https://github.com/poga/redis-percentile/assets/17446735/31d06b5f-367e-4b8b-a9f2-4c0a0154c471)

## Model Scored Evals

![S](https://github.com/poga/redis-percentile/assets/17446735/37e1fdc6-0f3c-4a46-9740-e91578ea9b69)

### States

This allows you to use states to set up different rules for things like when to stop a workflow, or to re-adjust the
cycle times.
Use info for general purpose information, and error for when you want to stop a workflow if any of the evals fail.

![sd](https://github.com/poga/redis-percentile/assets/17446735/877a353a-b405-45cd-9c05-f0d826e012eb)

### Model Scored JSON Output Evals

Setting the type to Model Scored will allow you to use the model to score the eval to your written instructions on a per
metric basis.
When you select this mode, the model will review the inputs and generate a JSON output with the metrics and their scores
and then
check it against your user inserted evals. If the evals pass, it will return true, otherwise false.

### Model Scored JSON Output Data Types

Select the JSON schema data-type that you want to use for the eval. The model will generate the JSON output based on the
data-type you select and
inject the JSON instructions into the model to generate the output into well formatted JSON responses.

![Scre](https://github.com/poga/redis-percentile/assets/17446735/267b7d7d-fbd1-4250-8053-85c735c7f3a2)

### Eval Scoring

You can select results types from: [Pass, Fail, Reject]

### Reject

Reject will discard the result and not use it in the analysis. This is useful for filtering out results that you
don't want to use in some analysis.

![Scree](https://github.com/poga/redis-percentile/assets/17446735/745f4044-c44f-4554-9fe9-7ac9e2b29071)

## API Evals

For this we will forward the response from the model stage directly to your own api endpoint. You can then score
using whatever method you want and return eval compatible JSON responses back. The model will then check the responses
against your evals tabulate the results.

![ScreenshoM](https://github.com/poga/redis-percentile/assets/17446735/1cf688cb-3902-4a72-85aa-6aeb28b2f0dd)
