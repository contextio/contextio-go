package ciolite

// Api functions that support: users/email_accounts/folders/messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// GetUserEmailAccountsFolderMessageParams query values data struct.
// Optional: Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags,
// and (for GetUserEmailAccountsFolderMessages only) Limit, Offset.
type GetUserEmailAccountsFolderMessageParams struct {
	// Optional:
	Delimiter    string `json:"delimiter,omitempty"`
	BodyType     string `json:"body_type,omitempty"`
	IncludeBody  bool   `json:"include_body,omitempty"`
	IncludeFlags bool   `json:"include_flags,omitempty"`

	// IncludeHeaders can be "0", "1", or "raw"
	IncludeHeaders string `json:"include_headers,omitempty"`

	// Optional for GetUserEmailAccountsFolderMessages (not used by GetUserEmailAccountFolderMessage):
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// GetUsersEmailAccountFolderMessagesResponse data struct
type GetUsersEmailAccountFolderMessagesResponse struct {
	MessageID   string `json:"message_id,omitempty"`
	Subject     string `json:"subject,omitempty"`
	InReplyTo   string `json:"in_reply_to,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`

	Folders         []string `json:"folders,omitempty"`
	References      []string `json:"references,omitempty"`
	ReceivedHeaders []string `json:"received_headers,omitempty"`

	ListHeaders ListHeaders `json:"list_headers,omitempty"`

	Addresses GetUsersEmailAccountFolderMessageAddresses `json:"addresses,omitempty"`

	PersonInfo PersonInfo `json:"person_info,omitempty"`

	Attachments []UsersEmailAccountFolderMessageAttachment `json:"attachments,omitempty"`

	Bodies []UsersEmailAccountFolderMessageBody `json:"bodies,omitempty"`

	SentAt     int `json:"sent_at,omitempty"`
	ReceivedAt int `json:"received_at,omitempty"`
}

// UsersEmailAccountFolderMessageAttachment embedded data struct within GetUsersEmailAccountFolderMessagesResponse
type UsersEmailAccountFolderMessageAttachment struct {
	Type               string `json:"type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	BodySection        string `json:"body_section,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	MessageID          string `json:"message_id,omitempty"`
	//XAttachmentID      string `json:"x_attachment_id,omitempty"` // Format is sometimes a string, sometimes an array of strings?

	Size         int `json:"size,omitempty"`
	AttachmentID int `json:"attachment_id,omitempty"`
}

// UsersEmailAccountFolderMessageBody embedded data struct within GetUsersEmailAccountFolderMessagesResponse
type UsersEmailAccountFolderMessageBody struct {
	BodySection string `json:"body_section,omitempty"`
	Type        string `json:"type,omitempty"`
	Encoding    string `json:"encoding,omitempty"`

	Size int `json:"size,omitempty"`
}

// ListHeaders embedded data struct within GetUsersEmailAccountFolderMessagesResponse
type ListHeaders map[string]string

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *ListHeaders) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty map
		*m = make(map[string]string)
		return nil
	}
	mp := make(map[string]string)
	err := json.Unmarshal(b, &mp)
	if err != nil {
		return err
	}
	*m = mp
	return nil
}

// PersonInfo embedded data struct within GetUsersEmailAccountFolderMessagesResponse and WebhookMessageData
type PersonInfo map[string]map[string]string

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *PersonInfo) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty map
		*m = make(map[string]map[string]string)
		return nil
	}
	mp := make(map[string]map[string]string)
	err := json.Unmarshal(b, &mp)
	if err != nil {
		return err
	}
	*m = mp
	return nil
}

// Address embedded data struct within GetUsersEmailAccountFolderMessageAddresses and WebhookMessageDataAddresses
type Address struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *Address) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty struct
		*m = Address{}
		return nil
	}
	// avoid recursion
	type addressTemp Address
	var tmp addressTemp

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*m = Address(tmp)
	return nil
}

// GetUsersEmailAccountFolderMessageAddresses data struct within GetUsersEmailAccountFolderMessagesResponse
type GetUsersEmailAccountFolderMessageAddresses struct {
	From    []Address `json:"from,omitempty"`
	To      []Address `json:"to,omitempty"`
	Cc      []Address `json:"cc,omitempty"`
	Bcc     []Address `json:"bcc,omitempty"`
	Sender  []Address `json:"sender,omitempty"`
	ReplyTo []Address `json:"reply_to,omitempty"`
}

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *GetUsersEmailAccountFolderMessageAddresses) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty struct
		*m = GetUsersEmailAccountFolderMessageAddresses{}
		return nil
	}
	// avoid recursion
	type getUsersEmailAccountFolderMessageAddressesTemp GetUsersEmailAccountFolderMessageAddresses
	var tmp getUsersEmailAccountFolderMessageAddressesTemp

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*m = GetUsersEmailAccountFolderMessageAddresses(tmp)
	return nil
}

// MoveUserEmailAccountFolderMessageParams form values data struct.
// Requires: NewFolderID, and may optionally contain Delimiter.
type MoveUserEmailAccountFolderMessageParams struct {
	// Required:
	NewFolderID string `json:"new_folder_id"`
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
}

// MoveUserEmailAccountFolderMessageResponse data struct
type MoveUserEmailAccountFolderMessageResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolderMessages gets listings of email messages for a user.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType,
// IncludeHeaders, IncludeFlags, Limit, Offset
func (cioLite CioLite) GetUserEmailAccountsFolderMessages(userID string, label string, folder string, queryValues GetUserEmailAccountsFolderMessageParams) ([]GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages", userID, label, url.QueryEscape(folder)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response []GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolderMessage gets file, contact and other information about a given email message.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags
func (cioLite CioLite) GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageParams) (GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// MoveUserEmailAccountFolderMessage moves a message.
// formValues requires NewFolderID, and may optionally contain Delimiter
func (cioLite CioLite) MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues MoveUserEmailAccountFolderMessageParams) (MoveUserEmailAccountFolderMessageResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "PUT",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response MoveUserEmailAccountFolderMessageResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
