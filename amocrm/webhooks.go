package amocrm

import (
	"context"
	"fmt"
)

// Webhook represents an AmoCRM webhook
type Webhook struct {
	ID          string   `json:"id,omitempty"`
	Destination string   `json:"destination"`
	Settings    []string `json:"settings"`
	Disabled    bool     `json:"disabled,omitempty"`
}

// WebhooksService handles communication with webhook-related methods
type WebhooksService struct {
	client *Client
}

// WebhooksResponse represents the API response for webhooks list
type WebhooksResponse struct {
	Embedded struct {
		Webhooks []Webhook `json:"webhooks"`
	} `json:"_embedded"`
	Links Links `json:"_links"`
}

// List retrieves a list of webhooks
func (s *WebhooksService) List(ctx context.Context) ([]Webhook, error) {
	var resp WebhooksResponse
	if err := s.client.GetJSON(ctx, "/webhooks", &resp); err != nil {
		return nil, err
	}

	return resp.Embedded.Webhooks, nil
}

// Subscribe creates a new webhook subscription
func (s *WebhooksService) Subscribe(ctx context.Context, webhook *Webhook) error {
	type request struct {
		Webhooks []Webhook `json:"webhooks"`
	}

	req := request{
		Webhooks: []Webhook{*webhook},
	}

	var resp WebhooksResponse
	if err := s.client.PostJSON(ctx, "/webhooks", req, &resp); err != nil {
		return err
	}

	return nil
}

// Unsubscribe deletes a webhook subscription
func (s *WebhooksService) Unsubscribe(ctx context.Context, webhookID string) error {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	return s.client.DeleteJSON(ctx, path)
}
