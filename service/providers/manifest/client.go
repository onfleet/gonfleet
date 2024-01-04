package manifest

import (
	"fmt"
	"net/http"

	onfleet "github.com/onfleet/gonfleet"
	"github.com/onfleet/gonfleet/netwrk"
)

type Client struct {
	apiKey       string
	rlHttpClient *netwrk.RlHttpClient
	url          string
	call         netwrk.Caller
}

func Plug(apiKey string, rlHttpClient *netwrk.RlHttpClient, url string, call netwrk.Caller) *Client {
	return &Client{
		apiKey:       apiKey,
		rlHttpClient: rlHttpClient,
		url:          url,
		call:         call,
	}
}

// Reference https://docs.onfleet.com/reference/delivery-manifest
func (c *Client) Generate(params *onfleet.ManifestGenerateParams, googleAPIKey string) (onfleet.DeliveryManifest, error) {
	deliveryManifest := onfleet.DeliveryManifest{}
	hubId := params.HubId
	workerId := params.WorkerId
	body := map[string]string{
		"path":   fmt.Sprintf("providers/manifest/generate?hubId=%s&workerId=%s", hubId, workerId),
		"method": "GET",
	}
	additionalAPIKeys := netwrk.AdditionalAPIKey{}
	if googleAPIKey != "" {
		additionalAPIKeys = netwrk.AdditionalAPIKey{
			Key:   "X-API-Key",
			Value: fmt.Sprintf("Google %s", googleAPIKey),
		}
	}

	err := c.call(
		c.apiKey,
		c.rlHttpClient,
		http.MethodPost,
		c.url,
		nil,
		nil,
		body,
		&deliveryManifest,
		additionalAPIKeys,
	)

	return deliveryManifest, err
}
