package ciolite

// Api functions that support: users/email_accounts/folders/messages/read

import (
	"fmt"
	"net/url"
)

// UserEmailAccountsFolderMessageReadResponse data struct
type UserEmailAccountsFolderMessageReadResponse struct {
	Success bool `json:"success,omitempty"`
}

// MarkUserEmailAccountsFolderMessageRead marks the message as read.
// formValues may optionally contain Delimiter
func (cioLite CioLite) MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "POST",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		FormValues:   formValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// MarkUserEmailAccountsFolderMessageUnRead marks the message as unread.
// formValues may optionally contain Delimiter
func (cioLite CioLite) MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		FormValues:   formValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
