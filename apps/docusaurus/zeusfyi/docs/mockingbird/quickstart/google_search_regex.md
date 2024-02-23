---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Google Search Regex

## Overview

The Google Search Regex Transformation allows for the dynamic modification of search query parameters in the
customsearch/v1 endpoint. It uses regular expressions (regex) to match and transform input query strings (q) into
URL-encoded strings that the Google Custom Search JSON API can interpret.

### Prerequisites

Before you begin, ensure you have:

- A Programmable Search Engine created and configured. (https://developers.google.com/custom-search/v1/overview)
- Located your Search Engine ID in the Programmable Search Engine control panel.
- Obtained an API key for the Custom Search JSON API.

## Implementation

URL Encoding:

Use the regex to transform the input q into a URL-encoded string.
Regex Application:

Apply regex options to the encoded q to refine the search query.
API Endpoint:

Construct the API endpoint by inserting the transformed q parameter.

```text
GET https://www.googleapis.com/customsearch/v1?q={q}&cx={SEARCH_KEY}&key=YOUR_API_KEY
```

Any `{param}` bracket enclosed value that gets substituted from your auto-eval JSON output parameters that match the
name of the bracket enclosed value.

## Example : Google Search Regex

* JSON output ML -> q parameter -> Eval Trigger -> Transformation -> Search API request -> Regex -> Next ML Stage

### Add Google Search Endpoint in Load Balancer

![Screensh](https://github.com/zeus-fyi/zeus/assets/17446735/78c6245d-ef25-45c2-a317-26848c479481)

### Add Google Search API Retrieval + Set Regex

![ScrM](https://github.com/zeus-fyi/zeus/assets/17446735/72ff01d4-f629-4b8a-82e4-50f34a32fe59)

### Set JSON Output Params, and Eval Criteria

This could be that it needs a certain keyword in the query, or length, or other factors.
![ScreenshoPM](https://github.com/zeus-fyi/zeus/assets/17446735/d5ee0d9f-a45a-4273-944d-9e9a5a608848)

### Create Trigger Google Search API + Regex

This will automatically trigger after passing the eval criteria.
![ScreeM](https://github.com/zeus-fyi/zeus/assets/17446735/ed8bab6c-5d74-4d7e-8ae6-6699bcdb55a7)

### Build Workflow

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/30151fde-a79b-473e-a00c-d9fb094ec391)
![Scr](https://github.com/zeus-fyi/zeus/assets/17446735/494e66f4-6976-4807-952f-c968186f6e59)

#### Sample JSON Schema

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/27de9fab-2cd9-49b7-935a-17c1ad96d365)

#### Sample Analysis Task

This adds the prompt injection to mock a retrieval for it to create a query for google search. It uses the JSON schema
we defined in the previous step.
![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/b01530d6-83fa-4614-908a-08961cee9c0f)

#### Sample Aggregation Task

![Scre](https://github.com/zeus-fyi/zeus/assets/17446735/80496fa4-68c8-41f0-985c-5324c31d2837)

#### Regexes Used

```regexp
r`https?://[^\s<>"]+|www\.[^\s<>"]+`
r`og:([^":]+)":\s*"([^"]+)`
```

## Additional Information

For more details on the Custom Search JSON API, refer to the official documentation.

https://developers.google.com/custom-search/v1/overview