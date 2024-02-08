---
sidebar_position: 1
displayed_sidebar: zK8s
---

# AWS - EKS Cluster Creation

### Overview

1. IAM: Authorization Setups
2. VPC Networking Setup
3. EKS Cluster Creation

## Create IAM service account EKS cluster role

### AWS UI Navigation: IAM>Dashboard

1. Create role
2. Select trusted entity

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/0c1ff8d0-4e38-4d56-b8ce-0dd90e9ec69a)

1. Select AWS service
2. Select EKS service
3. Select EKS-Cluster

### Add Permissions

![ScreenM](https://github.com/zeus-fyi/zeus/assets/17446735/3ee2e0cd-649c-4b50-9bf1-b34c82314c61)

Permissions policies:

- AmazonEKSClusterPolicy

### Name, review, and create

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/30ccc6f6-7e2f-4b4d-9d30-548a6cb30c07)

    Role name: zeusEksClusterRole
    Description: Allows Zeus access to other AWS service resources that
                 are required to operate clusters managed by EKS and Zeusfyi, Inc.

### Create role

![Screens](https://github.com/zeus-fyi/zeus/assets/17446735/4a3a8e82-d89a-4cb4-aa51-292d97e92d9e)

## Create VPC and Subnets for EKS

### AWS UI Navigation: VPC>Your VPCs>Create VPC

![ScreenshoPM](https://github.com/zeus-fyi/zeus/assets/17446735/8bd32bdc-a1c6-4e67-9ccd-e521ef6fb074)
![Scree](https://github.com/zeus-fyi/zeus/assets/17446735/1bd4fac6-be37-47a8-a052-d6ae70d9bf16)

![ScreenshM](https://github.com/zeus-fyi/zeus/assets/17446735/7d5e56fd-ec73-4fe2-a81b-953afccad87a)

#### VPC settings

    [X] VPC and more
    [X] Auto-generate
    Name: zeus-eks-us-east-2

Review and once everything looks good Create VPC

### Note Your VPC values need for EKS Setup.

![Scre](https://github.com/zeus-fyi/zeus/assets/17446735/400d4048-2b88-4e6d-a8fa-d95b59e8d173)

You'll need your VPC ID, and Security group ID for EKS Setup

## Create EKS Cluster

### AWS UI Navigation: EKS>Clusters>Create EKS cluster

We'll run through these steps.

![ScreensM](https://github.com/zeus-fyi/zeus/assets/17446735/0b1a5445-eff3-458a-a78e-d688bc45bdb0)

### Step 1: Configure Cluster

![Screenshot](https://github.com/zeus-fyi/zeus/assets/17446735/b05a4ed6-9bfa-485c-99a7-df5f910b0a9d)

    Name: zeus-eks-us-east-2
    Cluster service role: zeusEksClusterRole

using the cluster service role we setup previously

![ScreenshoM](https://github.com/zeus-fyi/zeus/assets/17446735/ca602581-bf8d-4aca-b3bc-81e1d134987c)

#### Bootstrap cluster admin access:

1. Allow cluster administrator access
2. Allow cluster administrator access for your IAM principal.

#### Cluster authentication mode:

1. EKS API and ConfigMap

The cluster will source authenticated IAM principals from both EKS access entry APIs and the aws-auth ConﬁgMap.

### Step 2: Specify Networking

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/d8ad6401-aea5-4ef8-9cb4-470f5b3e10e0)

1. Select a public subnet in each availability zone
2. Remove the private subnets (if any populated)
3. Add the security group from the VPC setup.

![ScreenshotM](https://github.com/zeus-fyi/zeus/assets/17446735/877cfcdd-3070-49fc-9f39-4e203e04d7c0)

1. Set cluster endpoint access to public

[X] Public-The cluster endpoint is accessible from outside of your VPC. Worker node traffic will leave your VPC to
connect to the endpoint.

### Step 3: Configure observability

Click next. We'll setup our own Prometheus installation instead of the AWS managed version.

### Step 4: Select add-ons

![Screensh](https://github.com/zeus-fyi/zeus/assets/17446735/06c4a6e8-27a6-4630-9411-4b95d38482e2)

Use the defaults, click next.

### Step 5: Configure selected add-ons settings

![ScreenshoPM](https://github.com/zeus-fyi/zeus/assets/17446735/509a02ad-d718-4746-830b-ce849db7d269)

Use the defaults, click next.

### Step 6: Review and create

Review the final configuration, then press Create.

Your EKS cluster may take up to an hour to finish creating.
Though it usually completes initial setup within 20–30 mins.
