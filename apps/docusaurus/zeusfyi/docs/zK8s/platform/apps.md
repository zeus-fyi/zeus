---
sidebar_position: 2
displayed_sidebar: zK8s
---

# Apps

## Intro

This tab shows you all your registered zK8s apps, and open source apps that are available to deploy on our platform.
Many other apps are on our GitHub page under the cookbook section, and will eventually be added into this UI page.

![Screenshot 2023-10-01 at 4 44 50 PM](https://github.com/zeus-fyi/zeus/assets/17446735/df285804-e268-454d-8103-a17ea3ce387c)

## How to Build a Workload via the UI

You can build base workloads via the UI. Add at least one cluster base, and at least one workload base, and then you can
select which components you want to include on the next page

![builder](https://github.com/zeus-fyi/zeus/assets/17446735/20b8be59-9438-481d-822d-aefafe469038)

Now you can add components to your workload like in this example, where we add a StatefulSet for Redis

![Screenshot 2023-10-01 at 4 57 25 PM](https://github.com/zeus-fyi/zeus/assets/17446735/c4a0c1ab-ab0d-4d85-91fe-5c2b29ef2b5b)

When you add a Service, it will automatically add your ports from the container spec, and you can add more ports if you
need to. You can also add a ConfigMap, and an Ingress, and then you'll preview your workload before saving it.

![Screenshot 2023-10-01 at 5 02 20 PM](https://github.com/zeus-fyi/zeus/assets/17446735/f9f3f2da-326b-4409-ac99-033d77d6d287)

When you select an Ingress, you'll select which port to expose as an API from the container ports you've specified and
then it will automatically configure Nginx, your web certificate for SSL, and your DNS record.

Additionally, your API is optionally configurable to come with authentication using your existing API key.

![Screenshot 2023-10-01 at 5.04.24 PM.png](..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2Fvar%2Ffolders%2Fcl%2Fypxrt_zd5dl62x74lxt969xh0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_fYVEa8%2FScreenshot%202023-10-01%20at%205.04.24%20PM.png)

Once you've configured your app, you'll need to generate a preview. This will show you the full kubernetes yaml
configuration for your app, and you can copy it, or save it to your account.

![Screenshot 2023-10-01 at 5 08 49 PM](https://github.com/zeus-fyi/zeus/assets/17446735/20c5c1b5-5f0e-4d69-b374-7dd8b4000088)
