package ciolite

// Api functions that support: discovery

// GetDiscoveryParams query values data struct.
// Requires Email.
type GetDiscoveryParams struct {
	// Required:
	Email string `json:"email"`
}

// GetDiscoveryResponse data struct
type GetDiscoveryResponse struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`

	// Value only appears if there is an error message
	Value string `json:"value,omitempty"`

	Documentation []interface{} `json:"documentation,omitempty"`

	Found bool `json:"found,omitempty"`

	IMAP GetDiscoveryIMAPResponse `json:"imap,omitempty"`
}

// GetDiscoveryIMAPResponse embedded data struct
type GetDiscoveryIMAPResponse struct {
	Server   string `json:"server,omitempty"`
	Username string `json:"username,omitempty"`

	UseSSL bool `json:"use_ssl,omitempty"`
	OAuth  bool `json:"oauth,omitempty"`

	Port int `json:"port,omitempty"`
}

// GetDiscovery attempts to discover connection settings for a given email address.
// queryValues requires Email to be set.
func (cioLite CioLite) GetDiscovery(queryValues GetDiscoveryParams) (GetDiscoveryResponse, error) {

	// Make request
	request := clientRequest{
		Method:      "GET",
		Path:        "/lite/discovery",
		QueryValues: queryValues,
	}

	// Make response
	var response GetDiscoveryResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
