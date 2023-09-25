### NVMe Setup Instructions:

See cloud provider specific setup instructions for NVMe setup:<br>

[zeus/cluster_resources/nvme](https://github.com/zeus-fyi/zeus/tree/main/zeus/cluster_resources/nvmee)

Kubernetes Local Persistent Volume Setup Notes:

- You must set a PersistentVolume nodeAffinity when using local volumes.

[https://kubernetes.io/docs/concepts/storage/volumes/#local](https://kubernetes.io/docs/concepts/storage/volumes/#local)

### AWS

storageClass: fast-disks

### Sui Recommended Specs:

CPUs: 8 physical cores / 16 vCPUs<br>
RAM: 128 GB<br>
Storage (SSD): 4 TB NVMe drive<br>

Additional Info: [https://docs.sui.io/build/fullnode](https://docs.sui.io/build/fullnode)

Snapshots: [https://docs.sui.io/build/snapshot](https://docs.sui.io/build/snapshot)

Testnet Seed Nodes:
```yaml
p2p-config:
  seed-peers:
    - address: /dns/ewr-tnt-ssfn-00.testnet.sui.io/udp/8084
      peer-id: df8a8d128051c249e224f95fcc463f518a0ebed8986bbdcc11ed751181fecd38
    - address: /dns/lax-tnt-ssfn-00.testnet.sui.io/udp/8084
      peer-id: f9a72a0a6c17eed09c27898eab389add704777c03e135846da2428f516a0c11d
    - address: /dns/lhr-tnt-ssfn-00.testnet.sui.io/udp/8084
      peer-id: 9393d6056bb9c9d8475a3cf3525c747257f17c6a698a7062cbbd1875bc6ef71e
    - address: /dns/mel-tnt-ssfn-00.testnet.sui.io/udp/8084
      peer-id: c88742f46e66a11cb8c84aca488065661401ef66f726cb9afeb8a5786d83456e
```

Mainnet Seed Nodes:
```yaml
p2p-config:
  seed-peers:
    - address: /dns/icn-00.mainnet.sui.io/udp/8084
      peer-id: 303f1f35afc9a6f82f8d21724f44e1245f4d8eca0806713a07c525dadda95a66
    - address: /dns/icn-01.mainnet.sui.io/udp/8084
      peer-id: cb7ce193cf7a41e9cc2f99e65dd1487b6314a57c74be42cc8c9225b203301812
    - address: /dns/mel-00.mainnet.sui.io/udp/8084
      peer-id: d32b55bdf1737ec415df8c88b3bf91e194b59ee3127e3f38ea46fd88ba2e7849
    - address: /dns/mel-01.mainnet.sui.io/udp/8084
      peer-id: bbf3be337fc16614a1953da83db729abfdc40596e197f36fe408574f7c9b780e
    - address: /dns/ewr-00.mainnet.sui.io/udp/8084
      peer-id: c7bf6cb93ca8fdda655c47ebb85ace28e6931464564332bf63e27e90199c50ee
    - address: /dns/ewr-01.mainnet.sui.io/udp/8084
      peer-id: 3227f8a05f0faa1a197c075d31135a366a1c6f3d4872cb8af66c14dea3e0eb66
    - address: /dns/sjc-00.mainnet.sui.io/udp/8084
      peer-id: 6f0b25087cd6b2fd2e4329bcf308ac95a37c49277dd7286b72470c124809db5b
    - address: /dns/sjc-01.mainnet.sui.io/udp/8084
      peer-id: af1d5d8468b3612ac2b6ff3ca91e99a71390dbe5b83dea9f6ae2da708d689227
    - address: /dns/lhr-00.mainnet.sui.io/udp/8084
      peer-id: c619a5e0f8f36eac45118c1f8bda28f0f508e2839042781f1d4a9818043f732c
    - address: /dns/lhr-01.mainnet.sui.io/udp/8084
      peer-id: 53dcedf250f73b1ec83250614498947db00d17c0181020fcdb7b6db12afbc175
```