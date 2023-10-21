---
sidebar_position: 3
displayed_sidebar: zK8s
---

# Clusters

## Default View

This shows all the clusters you have deployed and have access to, and their current status.

![Screenshot 2023-10-01 at 4 42 07 PM](https://github.com/zeus-fyi/zeus/assets/17446735/99bebef1-5a3b-45d5-9605-ac68658b6cac)

## App View

You can select the cluster name to show only the clusters for any specific app. Additionally, you can upgrade your fleet
across every cloud provider and region with a single click to the latest version of your app.

![Screenshot 2023-10-01 at 4 36 17 PM](https://github.com/zeus-fyi/zeus/assets/17446735/a62b00ce-5b81-48e5-b201-a3251872ebf2)

## App Deployment via UI

Once you've created your app, you can then visit the app page and review the app's details and best cloud provider and
region
option that makes best sense for your app. It uses your workload config to filter out servers that don't meet your
requirements. It also shows you the cost of the block storage, and other add-ons when applicable.

Also note, this will provision servers as well for you, as well as any additional add-ons like block storage.
When you deploy via the UI, it will also taint your app to the server by default. If you already have enough
server capacity you can select 0 and it will schedule on your pre-existing servers.

This process will create a namespace for you, and you can see your live workload details on the cluster page.

![Screenshot 2023-10-01 at 5 17 30 PM](https://github.com/zeus-fyi/zeus/assets/17446735/107ace70-1a9e-4208-a184-6d2d70d713a4)

## App Edits in UI

Sometimes it makes sense to want to make quick changes to your app, such as when in development, or when the change is
simple.
So we added this section where you review your current app configuration and update your workloads in-line.

![Screenshot 2023-10-01 at 5 21 18 PM](https://github.com/zeus-fyi/zeus/assets/17446735/fd4c1f56-1cde-45b1-b21f-075cf73533f3)
