---
sidebar_position: 3
displayed_sidebar: zK8s
---

# AWS - EKS Account Secrets

## Connect Service Account Authentication

Required secret name formatting for AWS services

    NAME_PREFIX=zeus
    SERVICE_KEY={CLOUD_PROVIDER}-{SERVICE}-{REGION}
    
    ACCESS_KEY_SUFFIX=service-account-access-key
    SECRET_KEY_SUFFIX=service-account-secret-key
    
    SECRET_NAME_ACCESS_KEY={NAME_PREFIX}-{SERVICE_KEY}-{ACCESS_KEY_SUFFIX}
    SECRET_NAME_SECRET_KEY={NAME_PREFIX}-{SERVICE_KEY}-{SECRET_KEY_SUFFIX}

Example

    CLOUD_PROVIDER=aws
    SERVICE=eks
    REGION=us-east-2

    SERVICE_KEY=aws-eks-us-east-2

Full example format for platform service account secret names

    SECRET_NAME_ACCESS_KEY=zeus-aws-eks-us-east-2-service-account-access-key
    SECRET_NAME_SECRET_KEY=zeus-aws-eks-us-east-2-service-account-secret-key

Key name for secret reference is your EKS cluster name

    KEY={YOUR_EKS_CLUSTER_NAME}

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/58a69aed-1188-45e6-9a1e-717c848ef90c)

Example

    KEY=zeus-eks-us-east-2

Full Example

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/e2b36677-bd2c-43b6-8dd9-571ba6b3cb8f)