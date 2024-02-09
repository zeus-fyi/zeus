---
sidebar_position: 4
displayed_sidebar: zK8s
---

# Cluster Access Dashboard

## Overview

![SM](https://github.com/zeus-fyi/zeus/assets/17446735/cb41a325-1003-44f4-9483-2ffe07632fda)

You can update your alias names, and environment tags here as well as disable context access.

## AWS - EKS Cluster Creation

After accessing the Cluster Access Dashboard, you'll see the service role added to the aws-auth ConfigMap.
This will allow the EKS cluster to authenticate with the IAM role and allow the role to assume the `system:masters` role
in the cluster.

![ScreenshAM](https://github.com/zeus-fyi/zeus/assets/17446735/00519ec6-83f7-4203-b10b-d34fc7c44d86)
