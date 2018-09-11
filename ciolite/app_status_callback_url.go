package ciolite

// Api functions that support: status_callback_url

// GetStatusCallbackURLResponse data struct
type GetStatusCallbackURLResponse struct {
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
	ResourceURL       string `json:"resource_url,omitempty"`
}

// CreateStatusCallbackURLParams form values data struct.
// Requires: StatusCallbackURL
type CreateStatusCallbackURLParams struct {
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
}

// CreateDeleteStatusCallbackURLResponse data struct
type CreateDeleteStatusCallbackURLResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetStatusCallbackURL gets a list of app status callback url's.
func (cioLite CioLite) GetStatusCallbackURL() (GetStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/app/status_callback_url",
	}

	// Make response
	var response GetStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateStatusCallbackURL create an app status callback url.
// Requires: StatusCallbackURL
func (cioLite CioLite) CreateStatusCallbackURL(formValues CreateStatusCallbackURLParams) (CreateDeleteStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/app/status_callback_url",
		FormValues: formValues,
	}

	// Make response
	var response CreateDeleteStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteStatusCallbackURL removes an app status callback url.
func (cioLite CioLite) DeleteStatusCallbackURL() (CreateDeleteStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   "/app/status_callback_url",
	}

	// Make response
	var response CreateDeleteStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
