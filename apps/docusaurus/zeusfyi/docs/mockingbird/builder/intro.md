---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Intro

The builder section allows you to build tasks (analysis, aggregation), retrievals, evals, and more from the UI and
then mix and match them to quickly create new workflows, which is ideal for rapid prototyping and experimentation.
We tried to keep things as simple as possible to use via the UI to make it as accessible as possible. Let us know if
anything is confusing or you have any questions.

## Key Components of the Workflow Builder

### Retrievals

Purpose: This stage is responsible for the initial data gathering. Users define the retrieval by naming it and
specifying the group and platform it associates with.
Functionality: Set up multiple retrieval stages to collect data from different sources or services like databases, APIs,
or social media platforms.

### Analysis

Function: Once data is retrieved, the analysis stage interprets and processes the information using selected AI models.
Configuration: Assign a name, group, and the AI model to be used for analysis, along with setting the analysis cycle
count, determining how frequently the analysis should occur.

### Aggregation

Aggregation stages combine the results from multiple analysis cycles into a cohesive summary.
Execution Logic: Define aggregation cycles relative to analysis cycles, ensuring aggregation happens after all necessary
analyses are completed.
Automated Evals
Purpose: Automated evals are used to evaluate the results of analyses or aggregations against a set of criteria to
validate or trigger further actions.
Cycles: Configure eval cycles in tandem with analysis cycles to ensure evaluations are timely and relevant.

### Fundamental Time Period

Definition: This is the base time period against which all cycles are referenced, determining the workflow's execution
intervals.
Application: Set the fundamental time step and unit to establish how often the workflow runs, which can range from
minutes to hours, depending on the needs of the task.

## Workflow Builder View

![Screen](https://github.com/zeus-fyi/zeus/assets/17446735/15ad98ff-ed0c-482c-9ced-9f4c08f222f7)

## Creating a Workflow

Define Workflow Name and Group: Start by giving your workflow a distinctive name and group for organizational purposes.

### Add Stages

Sequentially add retrieval, analysis, aggregation, and eval stages, specifying each stage's parameters and settings.

### Connect Stages

Link retrieval stages to analysis, ensuring data flows correctly through the workflow.

### Set Time Periods

Adjust the fundamental time period to control the overall pacing of the workflow execution.

### Save Workflow

Once all stages are correctly configured, save the workflow to activate it within the system.

## Example Usage

Consider a scenario where you are monitoring social media engagement. You would set up retrieval stages to pull data
from social media APIs, analyze the sentiment and relevance using an AI model, aggregate the analysis results to track
engagement trends over time, and evaluate these trends to decide on marketing strategies.