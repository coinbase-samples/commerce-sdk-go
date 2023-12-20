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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	chargesEndpoint = "/charges"
)

type CommerceError struct {
	ApiError *ChargeError
	Err      error
}

func (e *CommerceError) Error() string {
	if e.ApiError != nil {
		return fmt.Sprintf("API error: %v, warnings: %v", e.ApiError.Error, e.ApiError.Warnings)
	}
	return e.Err.Error()
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("X-CC-Api-Key", c.Credentials.ApiKey)
	req.Header.Set("X-CC-Version", "2018-03-22")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func (c *Client) CreateCharge(ctx context.Context, req *ChargeRequest) (*ChargeResponse, error) {

	if req.LocalPrice == nil {
		return nil, errors.New("LocalPrice is required for ChargeRequest")
	}

	if req.PricingType == "" {
		return nil, errors.New("PricingType is required for ChargeRequest")
	}

	url := fmt.Sprintf("%s%s", c.HttpBaseUrl, chargesEndpoint)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	c.setHeaders(httpReq)

	resp, err := c.HttpClient.Do(httpReq)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		chargeErr, err := handleErrorResponse(resp)
		if err != nil {
			return nil, &CommerceError{Err: err}
		}
		return nil, &CommerceError{ApiError: chargeErr}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var chargeResponse *ChargeResponse
	if err = json.Unmarshal(body, &chargeResponse); err != nil {
		return nil, err
	}

	return chargeResponse, nil
}

func (c *Client) GetCharge(ctx context.Context, chargeId string) (*ChargeResponse, error) {

	url := fmt.Sprintf("%s%s/%s", c.HttpBaseUrl, chargesEndpoint, chargeId)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	c.setHeaders(httpReq)

	resp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		chargeErr, err := handleErrorResponse(resp)
		if err != nil {
			return nil, &CommerceError{Err: err}
		}
		return nil, &CommerceError{ApiError: chargeErr}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var chargeResponse *ChargeResponse
	if err := json.Unmarshal(body, &chargeResponse); err != nil {
		return nil, err
	}

	return chargeResponse, nil

}

func handleErrorResponse(response *http.Response) (*ChargeError, error) {
	var chargeErr *ChargeError
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading body: %s", err)
		return nil, err
	}
	if err := json.Unmarshal(body, &chargeErr); err != nil {
		return nil, err
	}
	return chargeErr, nil
}
