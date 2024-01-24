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
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/coinbase-samples/commerce-sdk-go"
)

func TestShowEvent(t *testing.T) {

	eventId := "2c63ac0e-24a5-4a63-a28a-affbc92ade75"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	creds, err := commerce.ReadEnvCredentials("COMMERCE_API_KEY")
	if err != nil {
		t.Fatalf("error retireving credentials: %s ", err)
	}

	c := commerce.NewClient(creds, http.Client{})

	eventData, err := c.ShowEvent(ctx, eventId)
	if reflect.ValueOf(eventData).IsZero() {
		t.Fatalf("no event returned")
	}

	eventJson, err := json.MarshalIndent(eventData, "", " ")
	if err != nil {
		t.Fatalf("error converting event data to JSON: %s", err)
	}

	fmt.Printf("event successfully retrieved \n %s", string(eventJson[:]))

}
