package ciolite

// Api functions that support: users/email_accounts/folders/messages/body

import (
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageBodyParams query values data struct.
// Optional: Delimiter, Type.
type GetUserEmailAccountsFolderMessageBodyParams struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
	Type      string `json:"type,omitempty"`
}

// GetUserEmailAccountsFolderMessageBodyResponse data struct
type GetUserEmailAccountsFolderMessageBodyResponse struct {
	Type        string `json:"type,omitempty"`
	Charset     string `json:"charset,omitempty"`
	Content     string `json:"content,omitempty"`
	BodySection string `json:"body_section,omitempty"`
}

// GetUserEmailAccountsFolderMessageBody fetches the message body of a given email.
// queryValues may optionally contain Delimiter, Type
func (cioLite CioLite) GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageBodyParams) ([]GetUserEmailAccountsFolderMessageBodyResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/body", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageBodyResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
