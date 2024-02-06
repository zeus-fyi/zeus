---
sidebar_position: 2
displayed_sidebar: zK8s
---

# External Clusters

## AWS Platform Secrets

Format for secret key names

    KEY={CLOUD_PROVIDER}-{SERVICE}-{REGION}

Example

    aws-eks-us-east-2

Format for platform service account secret names

    zeus-{KEY}-service-account-access-key
    zeus-{KEY}-service-account-secret-key

Example

    zeus-aws-eks-us-east-2-service-account-access-key
    zeus-aws-eks-us-east-2-service-account-secret-key

Full Example

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/45f3253a-031b-41ea-b1c4-1f8c5ad1de5b)
