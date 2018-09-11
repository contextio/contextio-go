package ciolite

// Api functions that support: users/email_accounts/folders/messages/attachments

import (
	"bytes"
	"encoding/json"
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

	AttachmentLink string
}

// UnmarshalJSON is here because the `as_link` response is a raw response with an http link
func (m *GetUserEmailAccountsFolderMessageAttachmentsResponse) UnmarshalJSON(b []byte) error {

	// json response  should start with an open object
	if !bytes.HasPrefix(b, []byte("{")) {
		*m = GetUserEmailAccountsFolderMessageAttachmentsResponse{}
		m.AttachmentLink = string(b[:])
		return nil
	}

	// avoid recursion
	type temp GetUserEmailAccountsFolderMessageAttachmentsResponse
	var tmp temp

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	*m = GetUserEmailAccountsFolderMessageAttachmentsResponse(tmp)
	return nil
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

// GetUserEmailAccountsFolderMessageAttachmentParam query values data struct.
type GetUserEmailAccountsFolderMessageAttachmentParam struct {
	Delimiter string `json:"delimiter,omitempty"`
	AsLink    bool   `json:"as_link,omitempty"`
}

// GetUserEmailAccountsFolderMessageAttachment retrieves an email attachment.
// queryValues may optionally contain Delimiter and AsLink
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues GetUserEmailAccountsFolderMessageAttachmentParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

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
