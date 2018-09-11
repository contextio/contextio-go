package ciolite

// Api functions that support: users/email_accounts/folders/messages/flags

import (
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageFlagsResponse data struct
type GetUserEmailAccountsFolderMessageFlagsResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Flags UserEmailAccountsFolderMessageFlags `json:"flags,omitempty"`
}

// UserEmailAccountsFolderMessageFlags embedded data struct within GetUserEmailAccountsFolderMessageFlagsResponse
type UserEmailAccountsFolderMessageFlags struct {
	Read     bool `json:"read,omitempty"`
	Answered bool `json:"answered,omitempty"`
	Flagged  bool `json:"flagged,omitempty"`
	Draft    bool `json:"draft,omitempty"`
}

// GetUserEmailAccountsFolderMessageFlags returns the message flags.
// queryValues may optionally contain Delimiter
func (cioLite CioLite) GetUserEmailAccountsFolderMessageFlags(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageFlagsResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/flags", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageFlagsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
