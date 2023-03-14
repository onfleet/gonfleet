package onfleet

type Webhook struct {
	Count     int64   `json:"count"`
	ID        string  `json:"id"`
	IsEnabled bool    `json:"isEnabled"`
	Name      string  `json:"name"`
	Threshold float64 `json:"threshold,omitempty"`
	Trigger   int     `json:"trigger"`
	Url       string  `json:"url"`
}

type WebhookCreateParams struct {
	Name      string  `json:"name"`
	Threshold float64 `json:"threshold,omitempty"`
	Trigger   int     `json:"trigger"`
	Url       string  `json:"url"`
}
