###Please type your API-key from etherscan.io in a config.toml.Otherwise, all results will be 0

#
Service returns total transactions and sum amounts of block <br/>

For example run: 
#### go run main.go
and go to 
#### http://127.0.0.1:8080/api/block/109789/total
response:
```json5
{
"transactions": 1, // total transactions in the block
"amount": "4.99877009" // sum of transaction in block
}
```

