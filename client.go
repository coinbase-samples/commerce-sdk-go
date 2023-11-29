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

import "net/http"

var BaseUrl = "https://api.commerce.coinbase.com"

type Client struct {
	HttpClient  http.Client
	Credentials *Credentials
	HttpBaseUrl string
}

func (c *Client) BaseUrl(u string) *Client {
	c.HttpBaseUrl = u
	return c
}

func NewClient(credentials *Credentials, httpClient http.Client) *Client {

	return &Client{
		HttpClient:  httpClient,
		Credentials: credentials,
		HttpBaseUrl: BaseUrl,
	}

}