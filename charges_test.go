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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCreateCharge(t *testing.T) {
	pricing_type := "fixed_price"
	currency := "USD"
	chargeAmount := "1.00"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds, err := ReadEnvCredentials("COMMERCE_API_KEY")
	if err != nil {
		t.Fatalf("error retireving credentials: %s ", err)
	}

	c := NewClient(creds, http.Client{})
	req := &ChargeRequest{
		PricingType: pricing_type,
		LocalPrice: &LocalPrice{
			Amount:   chargeAmount,
			Currency: currency,
		},
	}

	chargeResponse, err := c.CreateCharge(ctx, req)
	if err != nil {
		if commErr, ok := err.(*CommerceError); ok {
			t.Fatalf("system error creating charge : %v", commErr)
		} else {
			t.Fatalf("error creating charge: %s", err)
		}
	}

	formattedResponse, err := json.MarshalIndent(chargeResponse, " ", " ")
	if err != nil {
		t.Fatalf("error formatting charge: %b", err)
	}

	fmt.Print(string(formattedResponse))
}
