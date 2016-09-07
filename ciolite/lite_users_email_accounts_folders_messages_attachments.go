package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/attachments

import (
	"fmt"
	"net/url"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFolderMessageAttachmentsResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
type GetUserEmailAccountsFolderMessageAttachmentsResponse struct {
	Type               string `json:"type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	BodySection        string `json:"body_section,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	EmailMessageID     string `json:"email_message_id,omitempty"`
	XAttachmentID      string `json:"x_attachment_id,omitempty"`

	Size         int `json:"size,omitempty"`
	AttachmentID int `json:"attachment_id,omitempty"`
}

// GetUserEmailAccountsFolderMessageAttachments gets listings of email attachments.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachments(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) ([]GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountsFolderMessageAttachment retrieves an email attachment.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID), attachmentID),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
