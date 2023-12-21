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
	"fmt"
	"time"
)

type Credentials struct {
	ApiKey string `json:"X-CC-Api-Key"`
}

type ChargeRequest struct {
	BuyerLocale string                  `json:"Buyer_locale,omitempty"`
	PricingType string                  `json:"pricing_type"`
	LocalPrice  *LocalPrice             `json:"local_price"`
	Metadata    *map[string]interface{} `json:"metadata,omitempty"`
	RedirectUrl string                  `json:"redirect_url,omitempty"`
	CancelUrl   string                  `json:"cancel_url,omitempty"`
}

type LocalPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type ChargeResponse struct {
	Data ChargeData `json:"data"`
}

type ChargeData struct {
	BrandColor       string     `json:"brand_color"`
	BrandLogoUrl     string     `json:"brand_logo_url"`
	ChargeKind       string     `json:"charge_kind"`
	Code             string     `json:"code"`
	ConfirmedAt      string     `json:"confirmed_at"`
	CreatedAt        string     `json:"created_at"`
	ExpiresAt        string     `json:"expires_at"`
	HostedUrl        string     `json:"hosted_url"`
	Id               string     `json:"id"`
	OrganizationName string     `json:"organization_name"`
	Pricing          Pricing    `json:"pricing"`
	PricingType      string     `json:"pricing_type"`
	Redirects        Redirects  `json:"redirects"`
	SupportEmail     string     `json:"support_email"`
	Timeline         []Timeline `json:"timeline"`
	Web3Data         Web3Data   `json:"web3_data"`
}

type Pricing struct {
	Local      Price `json:"local"`
	Settlement Price `json:"settlement"`
}

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Redirects struct {
	CancelUrl                string `json:"cancel_url"`
	SuccessUrl               string `json:"success_url"`
	WillRedirectAfterSuccess bool   `json:"will_redirect_after_success"`
}

type Timeline struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

type Web3Data struct {
	TransferIntent    TransferIntent    `json:"transfer_intent"`
	SuccessEvents     []Event           `json:"success_events"`
	FailureEvents     []interface{}     `json:"failure_events"`
	ContractAddresses map[string]string `json:"contract_addresses"`
}

type TransferIntent struct {
	CallData         CallData         `json:"call_data"`
	ResponseMetadata ResponseMetadata `json:"metadata"`
}

type CallData struct {
	Deadline          string `json:"deadline"`
	FeeAmount         string `json:"fee_amount"`
	Id                string `json:"id"`
	Operator          string `json:"operator"`
	Prefix            string `json:"prefix"`
	Recipient         string `json:"recipient"`
	RecipientAmount   string `json:"recipient_amount"`
	RecipientCurrency string `json:"recipient_currency"`
	RefundDestination string `json:"refund_destination"`
	Signature         string `json:"signature"`
}

type ResponseMetadata struct {
	ChainId         int    `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	Sender          string `json:"sender"`
}

type Event struct {
	Finalized         bool   `json:"finalized"`
	InputTokenAddress string `json:"input_token_address"`
	InputTokenAmount  string `json:"input_token_amount"`
	NetworkFeePaid    string `json:"network_fee_paid"`
	Recipient         string `json:"recipient"`
	Sender            string `json:"sender"`
	Timestamp         string `json:"timestamp"`
	TxHsh             string `json:"tx_hsh"`
}

type ErrorMessage struct {
	Value string `json:"message"`
}

// Events

type EventResponse struct {
	Pagination Pagination  `json:"pagination"`
	Data       []EventData `json:"data"`
}

type Pagination struct {
	Order         string     `json:"order"`
	StartingAfter *string    `json:"starting_after"`
	EndingBefore  *string    `json:"ending_before"`
	Total         int32      `json:"total"`
	Limit         int32      `json:"limit"`
	Yielded       int32      `json:"yielded"`
	PreviousUri   string     `json:"previous_uri"`
	NextUri       string     `json:"next_uri"`
	Data          ChargeData `json:"data"`
	CursorRange   []string   `json:"cursor_range"`
}

type EventData struct {
	ApiVersion string       `json:"api_version"`
	CreatedAt  time.Time    `json:"created_at"`
	Data       DetailedData `json:"data"`
	ID         string       `json:"id"`
	Resource   string       `json:"resource"`
	Type       string       `json:"type"`
}

type DetailedData struct {
	Id           string                  `json:"id"`
	Code         string                  `json:"code"`
	Pricing      Pricing                 `json:"pricing"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty"`
	Timeline     []Timeline              `json:"timeline"`
	Redirects    Redirects               `json:"redirects"`
	Web3Data     Web3Data                `json:"web3_data"`
	CreatedAt    time.Time               `json:"created_at"`
	ExpiresAt    time.Time               `json:"expires_at"`
	HostedUrl    string                  `json:"hosted_url"`
	BrandColor   string                  `json:"brand_color"`
	ChargeKind   string                  `json:"charge_kind"`
	PricingType  string                  `json:"pricing_type"`
	SupportEmail string                  `json:"support_email"`
	BrandLogoUrl string                  `json:"brand_logo_url"`
	OrgName      string                  `json:"organization_name"`
}

type ApiErrorDetail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type ChargeError struct {
	Status   int            `json:"status"`
	Error    ApiErrorDetail `json:"error"`
	Warnings []string       `json:"warnings"`
}

type CommerceError struct {
	ApiError *ChargeError
	Err      error
}

func (e CommerceError) Error() string {
	if e.ApiError != nil {
		return fmt.Sprintf("Commerce API error: %v, warnings: %v", e.ApiError.Error, e.ApiError.Warnings)
	}
	return e.Err.Error()
}
