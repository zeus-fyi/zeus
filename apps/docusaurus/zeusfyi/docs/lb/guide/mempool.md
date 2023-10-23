---
sidebar_position: 6
---

# Mempool Streaming

During beta testing we're offering unlimited free access to our mempool service at a level which is comparable to:

- Blox - Enterprise/Enterprise-Elite ($1250-5000/mo)
- Blocknative - Growth1-Growth2 ($1250-5000/mo)

To use the service, just create a websocket subscription like the below, and it'll stream mempool transactions
to you in real time. Don't forget to add your bearer token to the request header.

`wss://iris.zeus.fyi/v1/mempool`

Go code example:

```go
addr := flag.String("addr", "iris.zeus.fyi", "ws service address")
u := url.URL{Scheme: "wss", Host: *addr, Path: "/v1/mempool"}

requestHeader := http.Header{}
requestHeader.Add("Authorization", "Bearer "+t.BearerToken)
ws, _, werr := websocket.DefaultDialer.Dial(u.String(), requestHeader)
```

We're streaming txs from the following sources:

```go
// Unnamed contract addresses are Uniswap v2/3, Multicall, UniversalRouter (both versions)
var TxSources = map[string]bool{
"0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D": true,
"0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f": true,
"0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B": true,
"0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD": true,
"0xf164fC0Ec4E93095b804a4795bBe1e041497b92a": true,
"0xE592427A0AEce92De3Edee1F18E0157C05861564": true,
"0x1111111254EEB25477B68fb85Ed929f73A960582": true, // 1inch v5: Aggregation Router	 https://etherscan.io/address/0x1111111254eeb25477b68fb85ed929f73a960582#code
"0x1111111254fb6c44bAC0beD2854e76F90643097d": true, // 1inch v4: Aggregation Router https://etherscan.io/address/0x1111111254fb6c44bac0bed2854e76f90643097d
"0x111111125434b319222CdBf8C261674aDB56F3ae": true, // 1inch v2: Aggregation Router
"0x881D40237659C251811CEC9c364ef91dC08D300C": true, // metamask swap https://etherscan.io/address/0x881d40237659c251811cec9c364ef91dc08d300c
"0xDef1C0ded9bec7F1a1670819833240f027b25EfF": true, // 0x https://etherscan.io/address/0xdef1c0ded9bec7f1a1670819833240f027b25eff
"0xe66B31678d6C16E9ebf358268a790B763C133750": true, // 0x coinbase wallet proxy https://etherscan.io/address/0xe66b31678d6c16e9ebf358268a790b763c133750#code
"0x2a0c0DBEcC7E4D658f48E01e3fA353F44050c208": true, // idex https://etherscan.io/address/0x2a0c0dbecc7e4d658f48e01e3fa353f44050c208
"0x6131B5fae19EA4f9D964eAc0408E4408b66337b5": true, // KyberSwap: Meta Aggregation Router v2 https://etherscan.io/address/0x6131b5fae19ea4f9d964eac0408e4408b66337b5
"0x9008D19f58AAbD9eD0D60971565AA8510560ab41": true, // CoW settlement https://etherscan.io/address/0x9008d19f58aabd9ed0d60971565aa8510560ab41
}
```

# Don't wait. We'll be out of beta testing within the next couple of weeks, and free access will be gone forever.