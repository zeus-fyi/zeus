---
sidebar_position: 8
displayed_sidebar: mockingbird
---

# JSON Schema Builder

## Intro

The JSON Schema Builder and Integrator is a comprehensive tool designed to streamline the process of creating,
managing, and utilizing JSON schemas within AI-driven workflows. This tool bridges the gap between generating AI
JSON outputs and integrating them with various systems through API calls, while incorporating automated
evaluations and human approval processes.

![Screensho](https://github.com/zeus-fyi/zeus/assets/17446735/dd039918-acb7-4783-86a5-3aff043c043a)

## Features

### JSON Schema Definition and Storage

Define JSON Schemas: Users can easily define JSON schemas for specific tasks, ensuring that the AI-generated output
adheres to the required structure.
Store Schemas: Once defined, schemas are stored and can be reused across different AI tasks, providing consistency and
saving time.

### AI Task Connection

Attach to AI Tasks: Schemas can be linked to AI JSON output tasks, which use the defined schema to structure the output
data correctly.

### Automated Evaluations

Connect schemas to automated evaluations to ensure that the AI output meets quality and relevance standards before being
passed along the workflow.

### Temporal Workflow Integration

API Workflow Integration: Seamlessly integrate the JSON payloads into Temporal workflows that can make API calls to
external systems, allowing for robust process automation.

### Human Approval Triggers

Triggered Approvals: Set up conditions under which human approvals are required, adding an extra layer of verification
to the automated process

## Getting Started

### Step 1: Defining a JSON Schema

Access the Schema Builder: Navigate to the Schema section of the tool.
Create a New Schema: Enter a schema name, description, and select the response type (single object or array of objects).
Add Fields: Specify field names and data types, and provide descriptions for each.

### Step 2: Connecting to AI Tasks

Select AI Task: Choose the AI task you want to connect with your schema.
Attach Schema: Link the AI task with the appropriate schema to ensure structured output.

### Step 3: Setting up Automated Evaluations

Define Evaluation Criteria: Establish rules and conditions for the automated evaluation of AI outputs.
Attach Evaluations: Link these evaluations to the AI tasks to be automatically triggered upon task completion.

### Step 4: Integrating with Temporal Workflows

Use the API approval trigger, and connect your API, which will be sent the payload from the triggered output if
approved.
These all run within a Temporal workflow, that feature durable execution, so you can chain them together
with other tasks, and run them on a schedule, or manually.

### Step 5: Implementing Human Approval Processes

Define Approval Conditions: Determine the conditions under which human intervention is necessary.
Set up Notifications: Configure notifications to alert the appropriate personnel for approvals.
Approval Workflow: Integrate the approval process within the workflow before any external API calls.
