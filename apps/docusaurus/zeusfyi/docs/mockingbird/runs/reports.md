---
sidebar_position: 3
displayed_sidebar: mockingbird
---

# Reports

Shows the outputs of each stage within the workflow. You can click on the expand row button to see the details of the
mockingbird run.

![ddd](https://github.com/zeus-fyi/zeus/assets/17446735/c77d77c4-0023-47f1-8196-66fdd93eeefc)

### Example Run: Eval Results

![ScreenshotPM](https://github.com/zeus-fyi/zeus/assets/17446735/648dbda2-ce22-4eea-b698-5c012844d190)

### Example Run: Analysis->Aggregation Flows

This shows an analysis stage feeding the aggregation stage. Notice how the Completion Choices output from the
analysis stage are then used in the aggregation stage as the prompt body, which is how the aggregation stages work.
By default it will use the last analysis stage in the workflow, relative to the aggregation stage, and
for multiple analysis stages it will aggregate once every analysis stage has completed at least once
and then aggregate all the results from the analysis stages that are within the time window since last aggregation.

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/da52e86d-d24b-4ee0-9f49-37e2c3b6e7a8)

