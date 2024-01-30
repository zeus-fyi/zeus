---
sidebar_position: 4
displayed_sidebar: mockingbird
---

# Aggregations

Use this to tell the AI how to aggregate the results of your analysis chunks into a rolling aggregation window. If
aggregating on a single analysis, the aggregation cycle count sets how many base analysis cycles to aggregate on. If
aggregating on multiple analysis, it will aggregate whenever the the underlying analysis is run.

#### Diagram: Analysis->Aggregation Flow

```mermaid
flowchart TD
    A[Analysis Stage] -->|Completion Choices| B[Aggregation Stage]
    B --> C{Aggregation Decision}
    C -->|Single Analysis| D[Aggregate Once]
    C -->|Multiple Analyses| E[Wait for All Analyses]
    E --> F[All Completed]
    F --> G[Time Window Check]
    G -->|Within Window| H[Aggregate Results]
```

## Text Aggregations

![gg](https://github.com/zeus-fyi/zeus/assets/17446735/7fa3a1f2-cf92-4442-8b7f-50841dca7f02)

## JSON Aggregations

![ScreenshotPM](https://github.com/zeus-fyi/zeus/assets/17446735/e7d816a2-439e-40e9-b732-4dacfe27e09d)

### Example Run: Analysis->Aggregation Flows

This shows an analysis stage feeding the aggregation stage. Notice how the Completion Choices output from the
analysis stage are then used in the aggregation stage as the prompt body, which is how the aggregation stages work.
By default it will use the last analysis stage in the workflow, relative to the aggregation stage, and
for multiple analysis stages it will aggregate once every analysis stage has completed at least once
and then aggregate all the results from the analysis stages that are within the time window since last aggregation.

### Start::End SQL Search Window

The start and end fields in the report are the generated SQL window query that the AI uses to aggregate the results
when fetching from a time series indexed datastore.

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/da52e86d-d24b-4ee0-9f49-37e2c3b6e7a8)
