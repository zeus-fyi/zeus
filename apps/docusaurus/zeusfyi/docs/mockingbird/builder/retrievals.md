---
sidebar_position: 5
displayed_sidebar: mockingbird
---

# Inputs/Outputs

For generic retrievals we recommend you setup an API that can be called from the load balancer group for quick starting.

## APIs via Load Balancer

It will get data from each route in the load balancer group you reference here and then aggregate the results into a
single response.

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/f8011006-e92c-413a-9267-ea69165b91e9)

See the RPC Load Balancer Documentation for adding routes, and creating routing groups that you can use here.

## Verify Retrieval API Calls

In the search tab on the main AI page, you can verify the API calls using the API platform option and setting the route
path.

![Sc](https://github.com/zeus-fyi/zeus/assets/17446735/5839a764-8aec-4cf5-9313-1615a3ebfe5a)
