package ciolite

// Api functions that support: users/email_accounts/messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// GetUserEmailAccountsMessageParams query values data struct.
// Optional: Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags,
// and (for GetUserEmailAccountsMessages only) Limit, Offset.
type GetUserEmailAccountsMessageParams struct {
	// Optional:
	Delimiter    string `json:"delimiter,omitempty"`
	BodyType     string `json:"body_type,omitempty"`
	IncludeBody  bool   `json:"include_body,omitempty"`
	IncludeFlags bool   `json:"include_flags,omitempty"`

	// IncludeHeaders can be "0", "1", or "raw"
	IncludeHeaders string `json:"include_headers,omitempty"`

	// Optional for GetUserEmailAccountsMessages (not used by GetUserEmailAccountMessage):
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// GetUsersEmailAccountMessagesResponse data struct
type GetUsersEmailAccountMessagesResponse struct {
	MessageID   string `json:"message_id,omitempty"`
	Subject     string `json:"subject,omitempty"`
	InReplyTo   string `json:"in_reply_to,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`

	Folders         []string `json:"folders,omitempty"`
	References      []string `json:"references,omitempty"`
	ReceivedHeaders []string `json:"received_headers,omitempty"`

	ListHeaders ListHeaders `json:"list_headers,omitempty"`

	Addresses GetUsersEmailAccountMessageAddresses `json:"addresses,omitempty"`

	PersonInfo PersonInfo `json:"person_info,omitempty"`

	Attachments []UsersEmailAccountMessageAttachment `json:"attachments,omitempty"`

	Bodies []UsersEmailAccountMessageBody `json:"bodies,omitempty"`

	SentAt     int `json:"sent_at,omitempty"`
	ReceivedAt int `json:"received_at,omitempty"`
}

// UsersEmailAccountMessageAttachment embedded data struct within GetUsersEmailAccountMessagesResponse
type UsersEmailAccountMessageAttachment struct {
	Type               string `json:"type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	BodySection        string `json:"body_section,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	MessageID          string `json:"message_id,omitempty"`
	//XAttachmentID      string `json:"x_attachment_id,omitempty"` // Format is sometimes a string, sometimes an array of strings?

	Size         int `json:"size,omitempty"`
	AttachmentID int `json:"attachment_id,omitempty"`
}

// UsersEmailAccountMessageBody embedded data struct within GetUsersEmailAccountMessagesResponse
type UsersEmailAccountMessageBody struct {
	BodySection string `json:"body_section,omitempty"`
	Type        string `json:"type,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Content     string `json:"string,omitempty"`

	Size int `json:"size,omitempty"`
}

// GetUsersEmailAccountMessageAddresses data struct within GetUsersEmailAccountMessagesResponse
type GetUsersEmailAccountMessageAddresses struct {
	From    []Address `json:"from,omitempty"`
	To      []Address `json:"to,omitempty"`
	Cc      []Address `json:"cc,omitempty"`
	Bcc     []Address `json:"bcc,omitempty"`
	Sender  []Address `json:"sender,omitempty"`
	ReplyTo []Address `json:"reply_to,omitempty"`
}

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *GetUsersEmailAccountMessageAddresses) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty struct
		*m = GetUsersEmailAccountMessageAddresses{}
		return nil
	}
	// avoid recursion
	type getUsersEmailAccountMessageAddressesTemp GetUsersEmailAccountMessageAddresses
	var tmp getUsersEmailAccountMessageAddressesTemp

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*m = GetUsersEmailAccountMessageAddresses(tmp)
	return nil
}

// GetUserEmailAccountsMessages gets listings of email messages for a user.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType,
// IncludeHeaders, IncludeFlags, Limit, Offset
func (cioLite CioLite) GetUserEmailAccountsMessages(userID string, label string, queryValues GetUserEmailAccountsMessageParams) ([]GetUsersEmailAccountMessagesResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/messages", userID, label),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response []GetUsersEmailAccountMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountMessage gets file, contact and other information about a given email message.
// queryValues may optionally contain Delimiter, IncludeBody, BodyType, IncludeHeaders, IncludeFlags
func (cioLite CioLite) GetUserEmailAccountMessage(userID string, label string, messageID string, queryValues GetUserEmailAccountsMessageParams) (GetUsersEmailAccountMessagesResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s/messages/%s", userID, label, url.QueryEscape(messageID)),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUsersEmailAccountMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
