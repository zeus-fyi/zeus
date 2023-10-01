---
sidebar_position: 3
displayed_sidebar: zK8s
---

# Redis + KeyDB #

### Redis ###

#### ```cookbooks/redis ```

Contains full Kubernetes infra setup for open source BSD 3-Clause version of Redis with master-replica and cluster
configurations. We also include our Redis t-digest integration in our default Docker image.
You can find the pre-built bundle on our Docker repo zeusfyi/redis:latest

### KeyDB ###

#### ```cookbooks/keydb ```

Contains full Kubernetes infra setup for keyDB, which is a forked multi-threaded version of Redis that is maintained by
Snap Inc.
