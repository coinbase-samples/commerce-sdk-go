## Overview

The _Commerce Go SDK_ is a sample library that demonstrates the structure of a [Coinbase Commerce](https://commerce.coinbase.com/) driver for the [REST APIs](https://docs.cloud.coinbase.com/commerce/reference).

## License

The _Commerce Go SDK_ sample library is free and open source and released under the [Apache License, Version 2.0](https://github.com/coinbase-samples/commerce-sdk-go/blob/main/LICENSE).

The application and code are only available for demonstration purposes.

## Usage

To use the _Commerce Go SDK_, initialize the Credentials struct and create a new client. The Credentials struct is JSON enabled. Ensure that Commerce API credentials are stored in a secure manner.

### Example

This code snippet reads the value of the environment variable `COMMERCE_API_KEY` to initiate a new Commerce client.

**Copy Commerce SDK repository**

```
git clone git@github.com:coinbase-samples/commerce-sdk-go.git
```

**Set the API key as an environment variables in a terminal application**

```bash
export COMMERCE_API_KEY=<YOUR-COMMERCE-API-KEY>
```

**Initialize the Commerce client**

```go
creds, err := commerce.ReadEnvCredentials("COMMERCE_API_KEY")
if err != nil{
fmt.Printf("Error reading environmental variable: %s", err)
 }

client := commerce.NewClient(creds, http.Client{},)
```

Once a client is initialized, you may call any of the functions. For example, to create a charge,

```go
charge, err := client.CreateCharge(&commerce.ChargeRequest{
	PricingType: "fixed_price",
	LocalPrice: commerce.LocalPrice{
	Amount: "1.00",
	Currency: "USD",
	},
 })

if err != nil {
	fmt.Printf("Error: %s", err)
}

fmt.Print(charge.Data.ID)
```

### Quickstart example

The ["example"](https://github.com/coinbase-samples/commerce-sdk-go/tree/main/example) folder will contain the logic to create a charge in the amount of $1.00. To see this in action:

1. Change into the 'example' directory by running: `cd example`
2. Run `go run example.go`

Expected output:

```
created a charge ID: 64e05b38-a938-4620-a9ca-e3b806b3757b
got a hosted url: https://commerce.coinbase.com/charges/EGTHQJXZ
```

### Obtaining Coinbase Commerce credentials

Coinbase Commerce API keys can be created in the Commerce UI under [Settings --> Security.](https://beta.commerce.coinbase.com/settings/security)
