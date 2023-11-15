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
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) ListEvents() (*EventResponse, error) {
	url := c.HttpBaseUrl + "/events"

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("X-CC-Api-Key", c.Credentials.ApiKey)
	httpReq.Header.Set("X-CC-Version", "2018-03-22")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var eventResponse EventResponse
	err = json.NewDecoder(resp.Body).Decode(&eventResponse)
	if err != nil {
		return nil, err
	}

	return &eventResponse, nil

}

func (c *Client) ShowEvent(eventId string) (*EventResponse, error) {

	if eventId == "" {
		return nil, errors.New("Please enter an eventId")
	}

	url := fmt.Sprintf("%s/events/%s", c.HttpBaseUrl, eventId)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("X-CC-Api-Key", c.Credentials.ApiKey)
	httpReq.Header.Set("X-CC-Version", "2018-03-22")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var eventResponse EventResponse
	err = json.NewDecoder(resp.Body).Decode(&eventResponse)
	if err != nil {
		return nil, err
	}

	return &eventResponse, nil
}
