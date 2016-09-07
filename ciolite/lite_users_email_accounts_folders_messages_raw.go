package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/raw

import (
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageRawResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
type GetUserEmailAccountsFolderMessageRawResponse string

// GetUserEmailAccountsFolderMessageRaw fetches the raw RFC-822 message text of a given email.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/raw#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageRaw(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageRawResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/raw", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageRawResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
