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

Under the "Trust relationships" tab

![ScreeM](https://github.com/zeus-fyi/zeus/assets/17446735/8538ad25-021b-41bc-9d73-fffd62673c8d)

You'll need to set ec2 and eks service principals

### Trusted entities

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "eks.amazonaws.com",
          "ec2.amazonaws.com"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```

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

#### AWS UI Navigation: IAM>Policies>Create policy

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/9085dbfd-d773-4828-9580-177d55dc6682)

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/2edfe144-c40b-4b61-bf29-203903d38e35)

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/b83a508e-5152-4e78-986f-58e8caa6ea09)

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
        "arn:aws:iam::{USER_NUMBER}:role/AWS-EKS-Role",
        "arn:aws:iam::{USER_NUMBER}:role/zeusEksClusterRole"
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
      "Resource": "*"
    }
  ]
}
```

### Review permissions and create user

Alternatively, you can create the user and then attach the policies afterwards.

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/e91b0054-e39a-4536-81b0-c63a8cdec7ea)

## Create Access and Secret Keys

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/150b3991-1f12-4917-9e90-afb3a529fddc)

    [x] Third-party service
    [x] I understand the above recommendation and want to proceed to create an access key.

![Screen](https://github.com/zeus-fyi/zeus/assets/17446735/99189052-050e-45b7-8ed3-e82f237aacf2)

You'll need to add the access key and secret key to your platform secrets in the next step.

    Access key: {ACCESS_KEY}
    Secret access key: {SECRET_KEY}