/**
 * Copyright 2023-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coinbase-samples/commerce-sdk-go"
)

func main() {
	creds, err := commerce.ReadEnvCredentials("COMMERCE_API_KEY")
	if err != nil {
		fmt.Printf("Error reading environmental variable: %s", err)
	}

	client := commerce.NewClient(creds, http.Client{})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	charge, err := client.CreateCharge(ctx, &commerce.ChargeRequest{
		PricingType: "fixed_price",
		LocalPrice: &commerce.LocalPrice{
			Amount:   "1.00",
			Currency: "USD",
		},
	})

	if err != nil {
		if comErr, ok := err.(*commerce.CommerceError); ok {
			log.Fatalf("api error creating charge: %v", comErr)
		} else {
			log.Fatalf("error: %s\n", err)
		}
	}

	fmt.Printf("charge created \n Id: %v\n hosted_url: %s", charge.Data.Id, charge.Data.HostedUrl)
}
