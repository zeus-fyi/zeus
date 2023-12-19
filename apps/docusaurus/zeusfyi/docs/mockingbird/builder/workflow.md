---
sidebar_position: 2
displayed_sidebar: mockingbird
---

# Workflow

This allows you to write natural language instructions to chain to your search queries. Add a name for your workflow,
and then write instructions for the AI to follow, and it will save the workflow for you. You can
attach various tasks, retrievals, evals, and more to your workflow, and then run it on a schedule, or manually.

![zzzzzz](https://github.com/zeus-fyi/zeus/assets/17446735/195c966f-ef5b-4a5a-a09e-ad49d4e880f0)

## Retrievals

Press the Add Retrieval button to add a retrieval to your workflow. You can add as many retrievals as you want.
It will switch the below table to the retrievals you have saved, and you can select and add them to your workflow.

![ssss](https://github.com/zeus-fyi/zeus/assets/17446735/9fa6dc5e-e645-4421-89dd-7d5d03ea4a3e)

## Analysis

### Instructions

Similar to the retrievals instructions, you can add analysis instructions to your workflow.

![zz](https://github.com/zeus-fyi/zeus/assets/17446735/c47732cc-4470-474a-a7f7-5059090fc8dd)

### Retrieval-Analysis Relationships

Add retrievals to feed your analysis stages by selecting the appropriate relationships from the dropdowns and adding
sources.

![A](https://github.com/zeus-fyi/zeus/assets/17446735/b37e6266-0ac6-491a-983a-f558e0632f15)

### Analysis Time Cycles

One analysis cycle is equal to one fundamental time period.

## Aggregation

### Instructions

Similar to the analysis instructions, you can add aggregation instructions to your workflow. You must connect
at least one analysis stage to your aggregation stage.

### Analysis-Aggregation Dependencies

After you add the relationships, you'll see the analysis stages you've connected to your aggregation stage in Analysis->
Aggregation Dependencies.

![AAA](https://github.com/zeus-fyi/zeus/assets/17446735/23c40e45-8217-4838-8901-433ba1fdca77)

### Analysis-Aggregation Time Cycles

One aggregation cycle is equal to the longest of any dependent analysis cycles. If you have an analysis stage that
occurs every 2 time cycles,
and set the aggregation cycle count to 2, it will run on time cycle 4 after the analysis stage completes.

## Fundamental Time Period

Think of this as your AI Cpu Time. This is the minimum time interval between AI actions per cycle of workflow execution.
So if you have a 5 minute fundamental time period, and you have 3 analysis stages in your workflow per one cycle of run
time,
it will run 3 stages each 5 minutes.

## Run Overrides

Using the run scheduler, you can change your fundamental time period, and the run scheduler will automatically adjust
the run times,
this also changes the search window sizes for analysis, aggregation, etc stages.

![AAAA](https://github.com/zeus-fyi/zeus/assets/17446735/bc8099bc-b8b2-4c4a-bbb8-053c4d2f4c22)

## Workflows Table

You can view your saved workflows in the table below. You can also delete them, and preview the workflow instructions.
We do not have workflow editing setup via UI for public use yet, so you'll have to delete and recreate your workflow if
you want to change it.

![adfs](https://github.com/zeus-fyi/zeus/assets/17446735/77dadfd2-8f5d-4031-9024-1582745f0c96)
