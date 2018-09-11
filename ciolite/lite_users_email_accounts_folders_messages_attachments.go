package ciolite

// Api functions that support: users/email_accounts/folders/messages/attachments

import (
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageAttachmentsResponse data struct
type GetUserEmailAccountsFolderMessageAttachmentsResponse struct {
	Type               string `json:"type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	BodySection        string `json:"body_section,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	MessageID          string `json:"message_id,omitempty"`
	XAttachmentID      string `json:"x_attachment_id,omitempty"`

	Size         int `json:"size,omitempty"`
	AttachmentID int `json:"attachment_id,omitempty"`
}

// GetUserEmailAccountsFolderMessageAttachments gets listings of email attachments.
// queryValues may optionally contain Delimiter
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachments(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) ([]GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountsFolderMessageAttachment retrieves an email attachment.
// queryValues may optionally contain Delimiter
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID), attachmentID),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
