---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Search

This section allows you to preview and mock API retrievals, and search indexer results, that you can filter
and sort by date, and filter by negative and positive keywords.

### Search Response Format

    Timestamp | Platform | Body

### Global Indexer Search

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/bf809410-2fcd-4462-b73a-ac7688b12189)

Leaving platform and groups blank will search all platforms and groups you are indexing over the given time period.

## Twitter

![ScM](https://github.com/zeus-fyi/zeus/assets/17446735/f5d6c195-d1a4-49a9-aae3-6f6b234f1cf6)

### Search Response Format

    Timestamp | Platform | Tweet Body

## Reddit

![Scrd](https://github.com/zeus-fyi/zeus/assets/17446735/9ee60b2f-d9b9-4ab7-b3b1-311ad2956108)

### Search Response Format

    Timestamp | Platform | Subreddit | Post ID | Author | URL | Title | Body

Though it may contain additional field headers for now as we evaluate which should be the defaults,
the search response format will always include the timestamp, platform, and the body of the post.

## API Retrieval Mocking

Here you can verify the API calls using the API platform option and setting the route path. It will
use your authentication credentials to make the API call, if any are set in the platform secrets for the table,
and return the response.

![Sc](https://github.com/zeus-fyi/zeus/assets/17446735/5839a764-8aec-4cf5-9313-1615a3ebfe5a)
    