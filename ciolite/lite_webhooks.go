package ciolite

// Api functions that support: webhooks

import (
	"fmt"
)

// GetWebhooks gets listings of Webhooks configured for the application.
func (cioLite CioLite) GetWebhooks() ([]GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/lite/webhooks",
	}

	// Make response
	var response []GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetWebhook gets the properties of a given Webhook.
func (cioLite CioLite) GetWebhook(webhookID string) (GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/webhooks/%s", webhookID),
	}

	// Make response
	var response GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateWebhook creates a new Webhook for the application.
// formValues requires CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
func (cioLite CioLite) CreateWebhook(formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/lite/webhooks",
		FormValues: formValues,
	}

	// Make response
	var response CreateUserWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyWebhook changes the properties of a given Webhook.
// formValues requires Active
func (cioLite CioLite) ModifyWebhook(webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/webhooks/%s", webhookID),
		FormValues: formValues,
	}

	// Make response
	var response ModifyWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteWebhookAccount cancels a Webhook.
func (cioLite CioLite) DeleteWebhookAccount(webhookID string) (DeleteWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/webhooks/%s", webhookID),
	}

	// Make response
	var response DeleteWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
