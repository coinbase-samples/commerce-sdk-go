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

package commerce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateCharge(t *testing.T) {
	pricing_type := "fixed_price"
	currency := "USD"
	chargeAmount := "1.00"

	creds, err := ReadEnvCredentials("COMMERCE_API_KEY")
	if err != nil {
		fmt.Printf("Error retireving creds: %s ", err)
	}

	c := NewClient(creds, http.Client{})
	req := &ChargeRequest{
		PricingType: pricing_type,
		LocalPrice: &LocalPrice{
			Amount:   chargeAmount,
			Currency: currency,
		},
	}

	resp, err := c.CreateCharge(req)
	if err != nil {
		t.Fatalf("Error creating charge: %s", err)
	}

	formattedResponse, err := json.MarshalIndent(resp, " ", " ")
	if err != nil {
		t.Fatalf("Error formatting charge: %b", err)
	}

	fmt.Print(string(formattedResponse))
}
