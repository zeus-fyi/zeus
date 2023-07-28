The cookbook example shown here is how you could run it via Zeus & creating an ENV var trigger on deploy

To run it without the hook or off Zeus. Either set this value directly in the cm-anvil

```shell
data:
  start.sh: |-
    #!/bin/sh 
    exec anvil --fork-url ${RPC_URL} --host 0.0.0.0
```

Or manually create the ENV secret in your namespace

```yaml
            - name: RPC_URL
              valueFrom:
                secretKeyRef:
                  key: rpc
                  name: hardhat
```
You can also send anvil an api request to set the rpc url among other things.

More details found here: https://book.getfoundry.sh/reference/anvil/

We also maintain a web3 golang library that natively supports more anvil commands interchangeably with hardhat ones.
https://github.com/zeus-fyi/web3/blob/master/client/client.go

We also have a load balancer & serverless style deployments that manages fork states via session locks (see our Iris package) that is built for scale.
Our first Iris product is a few weeks from production. We're expecting to release a new product for rpc-loadbalancing in early August. 
Ask for details if you want to learn more: 

support@zeus.fyi