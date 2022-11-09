# go-nowpayments

This is a Go library for the [crypto NOWPayments API](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a) version 1.

Note that the current implementation focuses on the *payments API* only for now:

Part/Endpoint|Implemented
---|:---:
[Recurring payments](https://documenter.getpostman.com/view/7907941/S1a32n38#689df54e-9f43-42b3-bfe8-9bcca0444a6a)|No
[Billing (sub-partner)](https://documenter.getpostman.com/view/7907941/S1a32n38#a523b89b-40b7-4afe-b940-043d434a6c80)|No
[Payments](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a)|Yes
`GET /estimate`|:heavy_check_mark:
`GET /min-amount`|:heavy_check_mark:
`GET /payment`|:heavy_check_mark:
[Currencies](https://documenter.getpostman.com/view/7907941/S1a32n38#cb80ccdc-8f7c-426c-89df-1ed2241954a5)|Yes
`GET /currencies`|:heavy_check_mark:
[Payouts](https://documenter.getpostman.com/view/7907941/S1a32n38#138ee72b-4c4f-40d0-a565-4a1e907f4d94)|No
[API status](https://documenter.getpostman.com/view/7907941/S1a32n38#9998079f-dcc8-4e07-9ac7-3d52f0fd733a)|Yes
`GET /status`|:heavy_check_mark:
[Authentication](https://documenter.getpostman.com/view/7907941/S1a32n38#174cd8c5-5973-4be7-9213-05567f8adf27)|No
`POST /auth`|:heavy_check_mark:

## Installation

```bash
$ go get github.com/matm/go-nowpayments
```

## CLI Tool

A `np` tool is available to quickly play with the payments API from the command line.
```bash
$ go install github.com/matm/go-nowpayments/cmd/np@latest
```

