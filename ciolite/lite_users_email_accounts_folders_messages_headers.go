package ciolite

// Api functions that support: users/email_accounts/folders/messages/headers

import (
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageHeadersParams query values data struct.
// Optional: Delimiter, Raw.
type GetUserEmailAccountsFolderMessageHeadersParams struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
	Raw       bool   `json:"raw,omitempty"`
}

// GetUserEmailAccountsFolderMessageHeadersResponse data struct
type GetUserEmailAccountsFolderMessageHeadersResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Headers map[string][]string `json:"headers,omitempty"`
}

// GetUserEmailAccountsFolderMessageHeaders gets the complete headers of a given email message.
// queryValues may optionally contain Delimiter, Raw
func (cioLite CioLite) GetUserEmailAccountsFolderMessageHeaders(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageHeadersParams) (GetUserEmailAccountsFolderMessageHeadersResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/headers", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageHeadersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
