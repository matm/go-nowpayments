# go-nowpayments

This is a Go library for the [crypto NOWPayments API](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a) version 1.

Note that the current implementation mostly focuses on the payments API for now:

Topic|Endpoint|Implemented
---|:---|:---:
[Recurring payments](https://documenter.getpostman.com/view/7907941/S1a32n38#689df54e-9f43-42b3-bfe8-9bcca0444a6a)||No
[Billing (sub-partner)](https://documenter.getpostman.com/view/7907941/S1a32n38#a523b89b-40b7-4afe-b940-043d434a6c80)||No
[Payments](https://documenter.getpostman.com/view/7907941/S1a32n38#84c51632-01ad-49c0-96f8-fb4b5ad2b24a)||Yes
||Get estimated price|:heavy_check_mark:
||Get the minimum payment amount|:heavy_check_mark:
||Get payment status|:heavy_check_mark:
||Get list of payments|:heavy_check_mark:
||Get/Update payment estimate|:heavy_check_mark:
||Create invoice|:heavy_check_mark:
||Create payment|:heavy_check_mark:
||Create payment from invoice|:heavy_check_mark:
[Currencies](https://documenter.getpostman.com/view/7907941/S1a32n38#cb80ccdc-8f7c-426c-89df-1ed2241954a5)||Yes
||Get available currencies|:heavy_check_mark:
||Get available checked currencies|:heavy_check_mark:
[Payouts](https://documenter.getpostman.com/view/7907941/S1a32n38#138ee72b-4c4f-40d0-a565-4a1e907f4d94)||No
[API status](https://documenter.getpostman.com/view/7907941/S1a32n38#9998079f-dcc8-4e07-9ac7-3d52f0fd733a)||Yes
||Get API status|:heavy_check_mark:
[Authentication](https://documenter.getpostman.com/view/7907941/S1a32n38#174cd8c5-5973-4be7-9213-05567f8adf27)||Yes
||Authentication|:heavy_check_mark:

## Installation

```bash
$ go get github.com/matm/go-nowpayments
```

## CLI Tool

A `np` tool is available to easily play with the payments API from the command line. Please make sure to target the sandbox API server in this case.

Can be installed with:
```bash
$ go install github.com/matm/go-nowpayments/cmd/np@latest
```

The following commands are available:
```
Usage of np:
  -a float
        pay amount for new payment/invoice (default 2)
  -c    show list of selected currencies
  -d    turn debugging on
  -i    new invoice
  -l    list all payments
  -n    new payment
  -p string
        status of payment ID
  -pi string
        new payment from invoice ID
```

In order to work, `np` expects a JSON config file provided as an argument, like
```
$ np -f conf.json -c
```
to list all crypto currencies available for payments.

The JSON config file looks like
```json
{
  "server": "https://api-sandbox.nowpayments.io/v1",
  "login": "your_email_adresse",
  "password": "some_password",
  "apiKey": "the API key to use"
}
```

- `server` is the path to the API server, i.e. one of
  - sandbox: `https://api-sandbox.nowpayments.io/v1`
  - production: `https://api.nowpayments.io/v1`
- `login` and `password` are your NOWPayments credentials
- `apiKey` is one API key generated in your admin account

