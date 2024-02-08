---
sidebar_position: 2
displayed_sidebar: zK8s
---

# AWS - EKS Service Account

## Overview

This guide will walk you through the process of creating a service account in AWS for use with EKS.

## Add Additional Policy to EKS Cluster Role

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/7944efc1-2ce8-4694-97c5-2573118aa4c7)

- AmazonEKS_CNI_Policy (if not already attached)
- AmazonEKSClusterPolicy (should be attached during EKS cluster creation setup)

## Create AWS-EKS-Role

![ScreenshPM](https://github.com/zeus-fyi/zeus/assets/17446735/09f22345-de95-4f27-8328-91c890837ee1)

Amazon Managed Policies

- AmazonEC2ContainerRegistryReadOnly
- AmazonEKS_CNI_Policy
- AmazonEKSClusterPolicy
- AmazonEKSServicePolicy
- AmazonEKSWorkerNodePolicy

## Create IAM Service User

#### AWS UI Navigation: IAM>Dashboard

### Specify user details

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/565e4103-b797-4f92-87fe-238b7deba0dd)

### Permission options

Amazon Managed Policies

- AmazonEC2ContainerRegistryReadOnly
- AmazonEKS_CNI_Policy
- AmazonEKSWorkerNodePolicy
- AmazonSSMManagedInstanceCore
- AWSPriceListServiceFullAccess

![aws](https://github.com/zeus-fyi/zeus/assets/17446735/c8d72d5f-b31c-43e1-a8df-e790f2b636c6)

Customer Managed Policies

- ZeusEksServiceAccountPolicy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "VisualEditor0",
      "Effect": "Allow",
      "Action": [
        "ec2:*",
        "eks:*",
        "iam:GetRole"
      ],
      "Resource": "*"
    }
  ]
}
```

If you're using the same role name from the EKS cluster creation:

roleName: zeusEksClusterRole

- ZeusEksGetIamPolicy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "iam:GetRole",
        "iam:ListAttachedRolePolicies"
      ],
      "Resource": [
        "arn:aws:iam::{ACCOUNT_NUMBER}:role/AWS-EKS-Role",
        "arn:aws:iam::{ACCOUNT_NUMBER}:role/zeusEksClusterRole"
      ]
    }
  ]
}
```

- ZeusEksIamPolicy

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "iam:PassRole",
      "Resource": "arn:aws:iam::{ACCOUNT_NUMBER}:role/AWS-EKS-Role"
    }
  ]
}
```

### Review permissions and create user

Alternatively, you can create the user and then attach the policies afterwards.

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/e91b0054-e39a-4536-81b0-c63a8cdec7ea)