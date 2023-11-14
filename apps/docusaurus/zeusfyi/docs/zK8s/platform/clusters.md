---
sidebar_position: 3
displayed_sidebar: zK8s
---

# Clusters

## Default View

This shows all the clusters you have deployed and have access to, and their current status.

![Screenshot 2023-10-01 at 4 42 07 PM](https://github.com/zeus-fyi/zeus/assets/17446735/99bebef1-5a3b-45d5-9605-ac68658b6cac)

#### It'll help you answer questions like these quickly:

- How many clusters do I have?
- Where are they located?
- What cloud provider are they on?
- What region are they in?
- What Kubernetes cluster & namespace are they in?

## App View

You can select the cluster name to show only the clusters for any specific app. Additionally, you can upgrade your fleet
across every cloud provider and region with a single click to the latest version of your app. You can also destroy your
cluster and namespace from here.

### Fleet Upgrades & Rollout Restarts

You can also upgrade your fleet across every cloud provider and region with a single click to the latest version of your
app.

![Screenshot 2023-11-10 at 10 57 34PM](https://github.com/zeus-fyi/zeus/assets/17446735/f9552076-d05b-4713-a2d0-cbc7a0b3c1d5)

### App Deployment via UI

Once you've created your app, you can then visit the app page and review the app's details and best cloud provider and
region option that makes best sense for your app. It uses your workload config to filter out servers that don't meet
your requirements. It also shows you the cost of the block storage, and other add-ons when applicable.

Also note, this will provision servers as well for you, as well as any additional add-ons like block storage.
When you deploy via the UI, it will also taint your app to the server by default. If you already have enough
server capacity you can select 0, and it will schedule on your pre-existing servers.

This process will create a namespace for you, and you can see your live workload details on the cluster page.

![Screenshot 2023-10-01 at 5 17 30 PM](https://github.com/zeus-fyi/zeus/assets/17446735/107ace70-1a9e-4208-a184-6d2d70d713a4)

