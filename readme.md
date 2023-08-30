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




 


