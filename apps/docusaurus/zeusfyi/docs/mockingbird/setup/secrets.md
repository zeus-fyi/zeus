---
sidebar_position: 1
displayed_sidebar: mockingbird
---

# Secrets

### Integrating Secrets for Data Retrieval and Analysis

Integrating secrets into our platform allows you to seamlessly connect with various data sources by using
platform-specific API keys.
This ensures secure access to the data you need for retrieval and analysis.
The process involves setting secret keys and values that correspond to the API credentials required for integration.

### Setting Up Secrets

To integrate your platform API keys for data retrieval and to use your OpenAI API key, follow these steps:

### Navigate to the Secrets Section

This is where you will manage all your API keys and tokens securely.

![Screenshot](https://github.com/zeus-fyi/zeus/assets/17446735/4e927018-b089-44b9-a129-969c0d5ff7c6)

### Enter the Secret Name:

For each platform, you will need to use the exact secret name as listed below:

    openai-api-key
    reddit-oauth2-secret
    reddit-password
    reddit-username
    discord-api-key
    reddit-oauth2-public
    twitter-oauth2-public
    twitter-oauth2-secret

### Enter the Key Name

Set the key name to 'mockingbird' which is the identifier for the integration within the platform.

### Enter the Value

This is where you input the actual API key or token that you have obtained from the respective platform. The value
should be kept confidential.

### Save Secret

After entering the secret name, key, and value, save the secret to secure it within the system.

    Secret Name: openai-api-key
    Key Name: mockingbird
    Value: [Your OpenAI API Key Here]
    
    Secret Name: reddit-oauth2-secret
    Key Name: mockingbird
    Value: [Your Reddit OAuth2 Secret Here]

## API Keys Integration for Routing Groups

![image](https://github.com/zeus-fyi/zeus/assets/17446735/6f476ee0-df38-4bc9-9a9c-5a417ab5a342)

### Load Balancing and Bearer Tokens

For efficient web data retrieval, it's recommended to use a load balancer group. To append a bearer token to your
requests:

1. **Secret Name**: Use the format `api-{YOUR_ROUTING_GROUP}` for the secret name.
2. **Key Name**: Set the key name to `mockingbird`.
3. **Value**: The value should be the bearer token provided by the API service.

#### Example

If your load balancer routing group is `test`, the platform secret should be set as:

- **Secret Name**: `api-test`
- **Key Name**: `mockingbird`
- **Value**: `[Your Bearer Token]`

This configuration ensures that the bearer token is correctly appended to all outbound API requests for the `test`
routing group.

![Screensh](https://github.com/zeus-fyi/zeus/assets/17446735/e8598c39-56e9-427b-ac40-b676e165970a)

## Platform Secrets View

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/5f14b4cf-2285-40b9-a60f-a42f8060d7c1)