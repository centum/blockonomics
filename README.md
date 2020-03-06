# Description

This package is golang port [auto_invoice](https://github.com/blockonomics/auto_invoice) with additional capability.

# About

* Get BTC payment address
* Create blockonomics bitcoin invoice.
* No need for approvals/API key/documentation. Use your own bitcoin address to receive funds.
* Invoices are peer to peer encrypted. Bitcoin amount adjusts dynamically according to current price. Blockonomics 
supports almost all fiat currencies.

# Example

* Balance
```shell script
> blockonomics-cli balance --addr "1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp 1dice97ECuByXAvqXpaYzSaQuPVvrtmz6"
[
  {
    "addr": "1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp",
    "confirmed": 11509579,
    "unconfirmed": 0
  },
  {
    "addr": "1dice97ECuByXAvqXpaYzSaQuPVvrtmz6",
    "confirmed": 2347960,
    "unconfirmed": 0
  }
]
```
* Price
```shell script
> blockonomics-cli price
  9114.66
```
* Get address
```shell script
> blockonomics-cli addr_new --token your_token --account another_account --reset
```
* Make invoice
```shell script
> blockonomics-cli invoice --addr 1dicVTdHxTBFaNR2e57p6Zre3vtjJzqK8C --amount 10.00 --currency USD --description "Description invoice"
  "https://www.blockonomics.co/invoice/12761/#/?key=SECRETKEY"
```
* Transaction detail
```shell script
> blockonomics-cli tx_detail --txid=c4978bfc9b4cd632fb37eb5f69c7c686ae364d9cb1b32ec01c0f8bae72530a4e
{
  "vin": [
    {
      "address": "1AEgdWjJrEbroURgWmPrXkFdzxGxdF7c4G",
      "value": 5724464598
    }
  ],
  "vout": [
    {
      "address": "1FnQjXQc8F6jyjF8L92yLpnMhSWpw8t8jo",
      "value": 10000
    },
    {
      "address": "1AEgdWjJrEbroURgWmPrXkFdzxGxdF7c4G",
      "value": 4824404598
    },
    {
      "address": "33wBKF7y471qK9zuWQDHbesnGX8JL5YCbW",
      "value": 899950000
    }
  ],
  "status": "Confirmed",
  "fee": 100000,
  "time": 1577166893,
  "size": 257
}
```
* History
```shell script
> blockonomics-cli history --addr "1JJ5taVeiHcD6DXNkLLfocbHcE9Nzio1qV 13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM" 
{
  "pending": [],
  "history": [
    {
      "txid": "c74d59ef2de2029bc5f45d74673f05e743ceab463c1685ae72e33bd0527b9d80",
      "addr": [
        "13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM"
      ],
      "value": 588,
      "time": 1551328313
    },
    {
      "txid": "98120f3d9834dc61839339123001717218428397ea8ab48412e53aa2bb8fbd64",
      "addr": [
        "13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM"
      ],
      "value": 4532403,
      "time": 1495517482
    },
    {
      "txid": "ff6e8e30537908e60fc41552cf639df1332f74539cd666c7f429ec2ecd983510",
      "addr": [
        "1JJ5taVeiHcD6DXNkLLfocbHcE9Nzio1qV"
      ],
      "value": -497100500,
      "time": 1443475345
    },
    {
      "txid": "5e4e03748327a22288623b02dab1721ac9f8082c7294aaa7f9581be49dced2c5",
      "addr": [
        "1JJ5taVeiHcD6DXNkLLfocbHcE9Nzio1qV"
      ],
      "value": 497100500,
      "time": 1443423780
    },
    {
      "txid": "2d05f0c9c3e1c226e63b5fac240137687544cf631cd616fd34fd188fc9020866",
      "addr": [
        "13A1W4jLPP75pzvn2qJ5KyyqG3qPSpb9jM"
      ],
      "value": 5000000000,
      "time": 1231660825
    }
  ]
}
```
* Add address for monitoring
```shell script
blockonomics-cli addr_mon_add --addr "1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp" --tag "mining" --token your_token
```
* Remove address from monitoring
```shell script
blockonomics-cli addr_mon_del --addr "1dice8EMZmqKvrGE4Qc9bUFf9PX3xaYDp" --tag "mining" --token your_token
```
* List address from monitoring
```shell script
blockonomics-cli addr_mon_list --token your_token
```

Related links
* https://www.blockonomics.co/views/api.html
* https://github.com/blockonomics/auto_invoice
