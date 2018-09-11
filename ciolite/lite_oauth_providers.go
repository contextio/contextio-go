package ciolite

// Api functions that support: connect_tokens

import (
	"fmt"
)

// GetOAuthProvidersResponse data struct
type GetOAuthProvidersResponse struct {
	Type                   string `json:"type,omitempty"`
	ProviderConsumerKey    string `json:"provider_consumer_key,omitempty"`
	ProviderConsumerSecret string `json:"provider_consumer_secret,omitempty"`
	ResourceURL            string `json:"resource_url,omitempty"`
}

// CreateOAuthProviderParams form values data struct.
// Requires Type, ProviderConsumerKey, ProviderConsumerSecret.
type CreateOAuthProviderParams struct {
	// Requires:
	Type                   string `json:"type"`
	ProviderConsumerKey    string `json:"provider_consumer_key"`
	ProviderConsumerSecret string `json:"provider_consumer_secret"`
}

// CreateOAuthProviderResponse data struct
type CreateOAuthProviderResponse struct {
	Success             bool   `json:"success,omitempty"`
	ProviderConsumerKey string `json:"provider_consumer_key,omitempty"`
	ResourceURL         string `json:"resource_url,omitempty"`
}

// DeleteOAuthProviderResponse data struct
type DeleteOAuthProviderResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetOAuthProviders get the list of OAuth providers configured.
func (cioLite CioLite) GetOAuthProviders() ([]GetOAuthProvidersResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/lite/oauth_providers",
	}

	// Make response
	var response []GetOAuthProvidersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetOAuthProvider gets information about a given OAuth provider.
func (cioLite CioLite) GetOAuthProvider(key string) (GetOAuthProvidersResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/oauth_providers/%s", key),
	}

	// Make response
	var response GetOAuthProvidersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateOAuthProvider adds a new OAuth2 provider.
// formValues requires Type, ProviderConsumerKey, and ProviderConsumerSecret
func (cioLite CioLite) CreateOAuthProvider(formValues CreateOAuthProviderParams) (CreateOAuthProviderResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/lite/oauth_providers",
		FormValues: formValues,
	}

	// Make response
	var response CreateOAuthProviderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteOAuthProvider removes a given OAuth provider.
func (cioLite CioLite) DeleteOAuthProvider(key string) (DeleteOAuthProviderResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/oauth_providers/%s", key),
	}

	// Make response
	var response DeleteOAuthProviderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
