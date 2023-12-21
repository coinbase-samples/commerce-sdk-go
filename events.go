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
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListEvents(ctx context.Context) (*EventResponse, error) {
	url := c.HttpBaseUrl + "/events"

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

	if err := handleErrorResponse(resp); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	eventResponse := &EventResponse{}
	if err := json.Unmarshal(body, eventResponse); err != nil {
		return nil, err
	}

	return eventResponse, nil
}

func (c *Client) ShowEvent(ctx context.Context, eventId string) (*EventResponse, error) {

	if eventId == "" {
		return nil, errors.New("Please enter an eventId")
	}

	url := fmt.Sprintf("%s/events/%s", c.HttpBaseUrl, eventId)

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

	if err := handleErrorResponse(resp); err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	eventResponse := &EventResponse{}
	if err = json.Unmarshal(body, eventResponse); err != nil {
		return nil, err
	}

	return eventResponse, nil
}
