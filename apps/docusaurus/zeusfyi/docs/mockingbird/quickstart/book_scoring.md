---
sidebar_position: 2
displayed_sidebar: mockingbird
---

# Book Scoring

## Overview

We'll be having the AI create a list of books, and then aggregating the titles into a JSON schema for scoring via an
Auto-Eval.
This pattern is also generally useful for synthetic data generation for ML training, and for creating a scoring system
for any type of data.

## Step 1: Create an AI Analysis Task

### Prompt:

Write a list of 10 book titles, with 5 of them being sci-fi, and 5 of them being any other genre.

```json
[
  {
    "title": "book_name",
    "title": "book_name2"
  }
]
```

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/15d6bb84-f333-4d5d-8e39-88d0a3470cd0)

## Step 2: Create a JSON schema

### JSON Schema:

```json
{
  "book_scores": [
    {
      "score": 9,
      "title": "Dune"
    },
    {
      "score": 10,
      "title": "Foundation"
    },
    {
      "score": 8,
      "title": "Neuromancer"
    }
  ]
}
```

### JSON Schema Builder

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/8b9e951a-a8bb-4e16-b37d-f85995f86222)

#### JSON Schema Fields

##### Schema Name: book_scores

    Field Name: score
    Field Type: number

    Field Description: Score each book from 1-10, with 10 being the
    highest likelihood of being a science-fiction book.

    Field Name: title
    Field Type: string
    Field Description: The title of the book.

After saving, you'll be able to see this schema in the schema list.

![ScreM](https://github.com/zeus-fyi/zeus/assets/17446735/90011dbc-6ed4-44fc-ab90-331de96d7ca0)

## Step 4: Create an Aggregate Task

Since the JSON schema fields are already defined, we can now create an aggregation task to score the books.
We don't need to do any additional steps because the JSON schema field descriptions give the AI enough information to
score the books
based on the schema and field descriptions.

![S](https://github.com/zeus-fyi/zeus/assets/17446735/f021c825-5d8e-4a37-a2f7-260b04db0606)

## Step 5: Create an Eval Fn Task

Since the JSON schema fields are already defined, we can reuse this schema to create an Auto-Eval task to score the
books.
Set the criteria for the Auto-Eval to score passing if the book is rated > 2/5 for likelihood to be a sci-fi book, and
failing if it is not.

![Scr](https://github.com/zeus-fyi/zeus/assets/17446735/49211848-eb63-4bad-a00b-7834b0f2ff52)

## Step 6: Create the Workflow

### Add Tasks to the Workflow

- Add a name and group for the workflow
- Add the analysis task to the workflow.
- Add the aggregation task to the workflow.
- Add the eval fn task to the workflow.

### Connect Analysis -> Aggregation

- Connect the analysis task to the aggregation task.

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/f29003e6-26df-4c7d-8e15-d4c8a0a2cdc5)

### Connect Eval Fn -> Aggregation

- Toggle to the aggregation option

![ScreensPM](https://github.com/zeus-fyi/zeus/assets/17446735/a7810283-3321-492c-ac57-8f326c8a8ff0)

- Connect the aggregation task to the Auto-Eval task.
  ![ScM](https://github.com/zeus-fyi/zeus/assets/17446735/9830bd5e-fd75-4b06-809e-f6baf67df700)

### Review and Save

- Review the workflow to ensure all tasks are connected properly.
  ![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/838e4429-8493-4fa4-874e-ecf3f4c177cd)

Save the workflow.

### Run the Workflow

![ScreeM](https://github.com/zeus-fyi/zeus/assets/17446735/a63ddba1-5a70-4ecf-ab58-55ce27fdc539)

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/57bae935-7fc1-45d3-ba00-9ba50c4bd97f)

### Review the Results

![Scree](https://github.com/zeus-fyi/zeus/assets/17446735/5c75e4a8-1ba5-4cfb-ad48-135f4ec6f648)

![Screensh](https://github.com/zeus-fyi/zeus/assets/17446735/0f4ccb25-7624-4461-8ff1-fe8dcb908453)


