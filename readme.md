# bdbtc - BTC example with BlockDaemon

* Create source and dest address using address/gen.go
* Fund source using the faucet, e.g. `curl -X POST localhost:3000/faucet --data '{"address":"n2adeqJdu1SZUNieehDx4wbsLyG4a2tE9e"}'`
* View the faucet transaction using esplora and grab the pubkey script, plug into tx.go
* Run tx.go to generate the signed transaction
* Set the RAW_TX environment variable, optionally decode it, send it
* Note the transaction id, wait for the block to me mined.

Note - the above assumes we're running nigiri locally

Use the bitcoin cli via ngiri

```
$ alias bcli='docker exec -it bitcoin bitcoin-cli -datadir=config'
$ bcli getblockchaininfo
{
  "chain": "regtest",
  "blocks": 103,
  "headers": 103,
  "bestblockhash": "08f7d39ff04b6cfb9071b44c69be9580ba6a75687eac0bd1d32c35479950733d",
  "difficulty": 4.656542373906925e-10,
  "time": 1693256173,
  "mediantime": 1675717538,
  "verificationprogress": 1,
  "initialblockdownload": false,
  "chainwork": "00000000000000000000000000000000000000000000000000000000000000d0",
  "size_on_disk": 31585,
  "pruned": false,
  "warnings": ""
}
```


https://developer.bitcoin.org/examples/transactions.html

```
$ go run gen.go 
Enter key id
Key ID: uZQ4rHjR0HeM8wUbHbDNqByZ1sT4
Source address:  mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP
Dest address:  moPFejHJc4LZ1hgNLLKxYAWCgCe64qQRP3

curl -X POST localhost:3000/faucet --data '{"address":"mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP"}'
{"txId":"f517a44469936d300c6c14ec92f7eab65a689d0ab0c01dbf0403e806abfed3f1"}
```



```
$ go run tx.go 
Signed Tx: 0100000001f1d3feab06e80304bf1dc0b00a9d685ab6eaf792ec146c0c306d936944a417f5000000006b4830450221009ea6702b4fc9151285d944d3afe0849510f012ab928706adb89632a0a9c8808602200311ff1e4bd061019760cdddf4e21d6e54398ccf44ad65c5e2fa79f387e999610121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000
```

```
export RAW_TX="0100000001f1d3feab06e80304bf1dc0b00a9d685ab6eaf792ec146c0c306d936944a417f5000000006b4830450221009ea6702b4fc9151285d944d3afe0849510f012ab928706adb89632a0a9c8808602200311ff1e4bd061019760cdddf4e21d6e54398ccf44ad65c5e2fa79f387e999610121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000"

$ bcli decoderawtransaction $RAW_TX
{
  "txid": "40d976db9952143e1bea9189732090bba61cd2adf77981a1cfd7e467f0916136",
  "hash": "40d976db9952143e1bea9189732090bba61cd2adf77981a1cfd7e467f0916136",
  "version": 1,
  "size": 192,
  "vsize": 192,
  "weight": 768,
  "locktime": 0,
  "vin": [
    {
      "txid": "f517a44469936d300c6c14ec92f7eab65a689d0ab0c01dbf0403e806abfed3f1",
      "vout": 0,
      "scriptSig": {
        "asm": "30450221009ea6702b4fc9151285d944d3afe0849510f012ab928706adb89632a0a9c8808602200311ff1e4bd061019760cdddf4e21d6e54398ccf44ad65c5e2fa79f387e99961[ALL] 023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8",
        "hex": "4830450221009ea6702b4fc9151285d944d3afe0849510f012ab928706adb89632a0a9c8808602200311ff1e4bd061019760cdddf4e21d6e54398ccf44ad65c5e2fa79f387e999610121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8"
      },
      "sequence": 4294967295
    }
  ],
  "vout": [
    {
      "value": 0.99998000,
      "n": 0,
      "scriptPubKey": {
        "asm": "OP_DUP OP_HASH160 564c2ac2ccd4eaeaa50b1d6bebec864077aa6065 OP_EQUALVERIFY OP_CHECKSIG",
        "desc": "addr(moPFejHJc4LZ1hgNLLKxYAWCgCe64qQRP3)#f309pdtk",
        "hex": "76a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac",
        "address": "moPFejHJc4LZ1hgNLLKxYAWCgCe64qQRP3",
        "type": "pubkeyhash"
      }
    }
  ]
}


$ bcli -regtest sendrawtransaction $RAW_TX
40d976db9952143e1bea9189732090bba61cd2adf77981a1cfd7e467f0916136
```


==================

curl -sSL "https://mempool.space/testnet/api/address/mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP"


 {"address":"mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP","chain_stats":{"funded_txo_count":0,"funded_txo_sum":0,"spent_txo_count":0,"spent_txo_sum":0,"tx_count":0},"mempool_stats":{"funded_txo_count":0,"funded_txo_sum":0,"spent_txo_count":0,"spent_txo_sum":0,"tx_count":0}}


Fund via https://coinfaucet.eu/en/btc-testnet/

curl -sSL "https://mempool.space/testnet/api/address/mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP"
{"address":"mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP","chain_stats":{"funded_txo_count":0,"funded_txo_sum":0,"spent_txo_count":0,"spent_txo_sum":0,"tx_count":0},"mempool_stats":{"funded_txo_count":1,"funded_txo_sum":2495874,"spent_txo_count":0,"spent_txo_sum":0,"tx_count":1}}

curl -sSL "https://mempool.space/testnet/api/address/mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP/txs"

[
   {
      "txid":"36fb8a4907cbe5e906af4623585f49d9fca6e3317b9781ffe02178487fb2a288",
      "version":2,
      "locktime":2475611,
      "vin":[
         {
            "txid":"a3308a240c0e6683ebb84d17305ac039f2c98b48b580d0c496b5c6e3c6c1d21d",
            "vout":0,
            "prevout":{
               "scriptpubkey":"0014a8c16e6cd25df523ce15028f624e622323331c58",
               "scriptpubkey_asm":"OP_0 OP_PUSHBYTES_20 a8c16e6cd25df523ce15028f624e622323331c58",
               "scriptpubkey_type":"v0_p2wpkh",
               "scriptpubkey_address":"tb1q4rqkumxjth6j8ns4q28kynnzyv3nx8zcg4pl2x",
               "value":15570579207
            },
            "scriptsig":"",
            "scriptsig_asm":"",
            "witness":[
               "304402203334b42b397a00c18601c56ab7ca13475345c05c2e89b50a14f5a9793b1d5241022001e8213be09e7750320e6170c89ef9d13b20c3d6516fde9fdc5ccf8fd72fde8a01",
               "02e8b143e62f2744b949ff4fa97da0fbd67b5d4c8debf7303e0957adae3ff593ef"
            ],
            "is_coinbase":false,
            "sequence":4294967293
         }
      ],
      "vout":[
         {
            "scriptpubkey":"76a914cf8c27c50a962ce9d5750121cbef451d0b06112a88ac",
            "scriptpubkey_asm":"OP_DUP OP_HASH160 OP_PUSHBYTES_20 cf8c27c50a962ce9d5750121cbef451d0b06112a OP_EQUALVERIFY OP_CHECKSIG",
            "scriptpubkey_type":"p2pkh",
            "scriptpubkey_address":"mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP",
            "value":2495874
         },
         {
            "scriptpubkey":"76a9149744064648d2a6b7b157836e6f615b1d02fd1da588ac",
            "scriptpubkey_asm":"OP_DUP OP_HASH160 OP_PUSHBYTES_20 9744064648d2a6b7b157836e6f615b1d02fd1da5 OP_EQUALVERIFY OP_CHECKSIG",
            "scriptpubkey_type":"p2pkh",
            "scriptpubkey_address":"muJmpSqMhGbwTBpKdgW44UZJPj9gvZaLH1",
            "value":15568068633
         }
      ],
      "size":228,
      "weight":585,
      "fee":14700,
      "status":{
         "confirmed":true,
         "block_height":2475612,
         "block_hash":"0000000000000015e0fbbdb69983f4851e4a560694109c3ac5cd80be6051b048",
         "block_time":1693417900
      }
   }
]


go run tx.go 
Signed Tx: 01000000011dd2c1c6e3c6b596c4d080b5488bc9f239c05a30174db8eb83660e0c248a30a3000000006b483045022100c4ef45c30c17d46618b82a56072a323945537ce168e8a3fb9ae0d8254789dc8f02204efd7b210cb4f70c1e2e7c9b111d4f13831e52ea6d5b5f06750b68fa3d9926720121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000



https://www.blockcypher.com/dev/bitcoin/?shell#customizing-transaction-requests


xxxxxxxxxxxxxxxxxxxx


curl https://api.blockcypher.com/v1/btc/test3/addrs/mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP

transaction ref - 36fb8a4907cbe5e906af4623585f49d9fca6e3317b9781ffe02178487fb2a288

curl https://api.blockcypher.com/v1/btc/test3/txs/36fb8a4907cbe5e906af4623585f49d9fca6e3317b9781ffe02178487fb2a288


curl https://api.blockcypher.com/v1/btc/test3/txs/36fb8a4907cbe5e906af4623585f49d9fca6e3317b9781ffe02178487fb2a288
{
  "block_hash": "0000000000000015e0fbbdb69983f4851e4a560694109c3ac5cd80be6051b048",
  "block_height": 2475612,
  "block_index": 11,
  "hash": "36fb8a4907cbe5e906af4623585f49d9fca6e3317b9781ffe02178487fb2a288",
  "addresses": [
    "muJmpSqMhGbwTBpKdgW44UZJPj9gvZaLH1",
    "mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP",
    "tb1q4rqkumxjth6j8ns4q28kynnzyv3nx8zcg4pl2x"
  ],
  "total": 15570564507,
  "fees": 14700,
  "size": 228,
  "vsize": 147,
  "preference": "high",
  "confirmed": "2023-08-30T17:51:40Z",
  "received": "2023-08-30T17:51:40Z",
  "ver": 2,
  "lock_time": 2475611,
  "double_spend": false,
  "vin_sz": 1,
  "vout_sz": 2,
  "opt_in_rbf": true,
  "confirmations": 9,
  "confidence": 1,
  "inputs": [
    {
      "prev_hash": "a3308a240c0e6683ebb84d17305ac039f2c98b48b580d0c496b5c6e3c6c1d21d",
      "output_index": 0,
      "output_value": 15570579207,
      "sequence": 4294967293,
      "addresses": [
        "tb1q4rqkumxjth6j8ns4q28kynnzyv3nx8zcg4pl2x"
      ],
      "script_type": "pay-to-witness-pubkey-hash",
      "age": 2475612,
      "witness": [
        "304402203334b42b397a00c18601c56ab7ca13475345c05c2e89b50a14f5a9793b1d5241022001e8213be09e7750320e6170c89ef9d13b20c3d6516fde9fdc5ccf8fd72fde8a01",
        "02e8b143e62f2744b949ff4fa97da0fbd67b5d4c8debf7303e0957adae3ff593ef"
      ]
    }
  ],
  "outputs": [
    {
      "value": 2495874,
      "script": "76a914cf8c27c50a962ce9d5750121cbef451d0b06112a88ac",
      "addresses": [
        "mzSN4vjy8Pc2UXYQfFc8KjAs8UmoPWbYUP"
      ],
      "script_type": "pay-to-pubkey-hash"
    },
    {
      "value": 15568068633,
      "script": "76a9149744064648d2a6b7b157836e6f615b1d02fd1da588ac",
      "spent_by": "495c6b546cef9cf65e689e5b48f4233b109c0b6fa88660f55396686cabe5cb6a",
      "addresses": [
        "muJmpSqMhGbwTBpKdgW44UZJPj9gvZaLH1"
      ],
      "script_type": "pay-to-pubkey-hash"
    }
  ]
}

01000000011dd2c1c6e3c6b596c4d080b5488bc9f239c05a30174db8eb83660e0c248a30a3000000006b483045022100c46e4c69bfc9881510eeb2d1de7884647a758f45d243c0a8f656e85c2a641f9b022014ab4dfe3605292fdd2f374e4dff9ebc2d38ae8fdece3ced3f920312a51e8d340121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000

curl -d '{"tx":"01000000011dd2c1c6e3c6b596c4d080b5488bc9f239c05a30174db8eb83660e0c248a30a3000000006b483045022100c46e4c69bfc9881510eeb2d1de7884647a758f45d243c0a8f656e85c2a641f9b022014ab4dfe3605292fdd2f374e4dff9ebc2d38ae8fdece3ced3f920312a51e8d340121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000"}' https://api.blockcypher.com/v1/bcy/test/txs/decode?token=$TOKEN



curl -d '{"tx":"01000000011dd2c1c6e3c6b596c4d080b5488bc9f239c05a30174db8eb83660e0c248a30a3000000006b483045022100c46e4c69bfc9881510eeb2d1de7884647a758f45d243c0a8f656e85c2a641f9b022014ab4dfe3605292fdd2f374e4dff9ebc2d38ae8fdece3ced3f920312a51e8d340121023eb874189980351d2db5f478c56339bab9aa67bd20700c13668b20abf0ff56d8ffffffff0130d9f505000000001976a914564c2ac2ccd4eaeaa50b1d6bebec864077aa606588ac00000000"}' https://api.blockcypher.com/v1/btc/test3/txs/push?token=$TOKEN