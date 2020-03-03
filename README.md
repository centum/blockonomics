# Description

This package is golang port [auto_invoice](https://github.com/blockonomics/auto_invoice) with additional capability.

# About

* Get BTC payment address
* Create blockonomics bitcoin invoice.
* No need for approvals/API key/documentation. Use your own bitcoin address to receive funds.
* Invoices are peer to peer encrypted. Bitcoin amount adjusts dynamically according to current price. Blockonomics 
supports almost all fiat currencies.

# Example

* Get address
```shell script
blockonomics-cli addr_new --token your_token --account another_account --reset
```
* Make invoice
```shell script
blockonomics-cli invoice --addr 1dicVTdHxTBFaNRAe54p6Zre3vtjJzqK8C --amount 10.00 --currency USD --description "Description invoice"
```

Related links
* https://www.blockonomics.co/views/api.html
* https://github.com/blockonomics/auto_invoice
