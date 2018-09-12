package ciolite

// Api functions that support: users/webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/mail"
)

// GetUsersWebhooksResponse data struct
type GetUsersWebhooksResponse struct {
	CallbackURL        string `json:"callback_url,omitempty"`
	WebhookID          string `json:"webhook_id,omitempty"`
	FilterTo           string `json:"filter_to,omitempty"`
	FilterFrom         string `json:"filter_from,omitempty"`
	FilterCc           string `json:"filter_cc,omitempty"`
	FilterSubject      string `json:"filter_subject,omitempty"`
	FilterThread       string `json:"filter_thread,omitempty"`
	FilterNewImportant string `json:"filter_new_important,omitempty"`
	FilterFileName     string `json:"filter_file_name,omitempty"`
	FilterFolderAdded  string `json:"filter_folder_added,omitempty"`
	FilterToDomain     string `json:"filter_to_domain,omitempty"`
	FilterFromDomain   string `json:"filter_from_domain,omitempty"`
	BodyType           string `json:"body_type,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`

	Active            bool `json:"active,omitempty"`
	Failure           bool `json:"failure,omitempty"`
	IncludeBody       bool `json:"include_body,omitempty"`
	IncludeHeader     bool `json:"include_header,omitempty"`
	ReceiveDrafts     bool `json:"receive_drafts,omitempty"`
	ReceiveAllChanges bool `json:"receive_all_changes,omitempty"`
	ReceiveHistorical bool `json:"receive_historical,omitempty"`
}

// CreateUserWebhookParams form values data struct.
// Requires: CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
type CreateUserWebhookParams struct {
	// Requires:
	CallbackURL string `json:"callback_url"`

	// Optional:
	FilterTo           string `json:"filter_to,omitempty"`
	FilterFrom         string `json:"filter_from,omitempty"`
	FilterCC           string `json:"filter_cc,omitempty"`
	FilterSubject      string `json:"filter_subject,omitempty"`
	FilterThread       string `json:"filter_thread,omitempty"`
	FilterNewImportant string `json:"filter_new_important,omitempty"`
	FilterFileName     string `json:"filter_file_name,omitempty"`
	FilterFolderAdded  string `json:"filter_folder_added,omitempty"`
	FilterToDomain     string `json:"filter_to_domain,omitempty"`
	FilterFromDomain   string `json:"filter_from_domain,omitempty"`
	BodyType           string `json:"body_type,omitempty"`
	IncludeBody        bool   `json:"include_body,omitempty"`
	IncludeHeader      bool   `json:"include_header,omitempty"`
	ReceiveDrafts      bool   `json:"receive_drafts,omitempty"`
	ReceiveAllChanges  bool   `json:"receive_all_changes,omitempty"`
	ReceiveHistorical  bool   `json:"receive_historical,omitempty"`
}

// CreateUserWebhookResponse data struct
type CreateUserWebhookResponse struct {
	WebhookID   string `json:"webhook_id,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// ModifyUserWebhookParams form values data struct.
// formValues requires Active
type ModifyUserWebhookParams struct {
	// Required:
	Active bool `json:"active"`
}

// ModifyWebhookResponse data struct
type ModifyWebhookResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// DeleteWebhookResponse data struct
type DeleteWebhookResponse struct {
	Success bool `json:"success,omitempty"`
}

// WebhookCallback data struct that will be received from CIO
type WebhookCallback struct {
	AccountID string `json:"account_id,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
	Token     string `json:"token,omitempty" valid:"required"`
	Signature string `json:"signature,omitempty" valid:"required"`

	Timestamp int `json:"timestamp,omitempty" valid:"required"`

	// Data is an error message that gives more information about the cause of failure
	Data string `json:"data,omitempty"`

	MessageData WebhookMessageData `json:"message_data,omitempty"`
}

// WebhookMessageData data struct within WebhookCallback
type WebhookMessageData struct {
	MessageID string `json:"message_id,omitempty"`
	Subject   string `json:"subject,omitempty"`

	References []string `json:"references,omitempty"`
	Folders    []string `json:"folders,omitempty"`

	Date         int `json:"date,omitempty"`
	DateReceived int `json:"date_received,omitempty"`

	Addresses WebhookMessageDataAddresses `json:"addresses,omitempty"`

	PersonInfo PersonInfo `json:"person_info,omitempty"`

	Flags WebhookMessageDataFlags `json:"flags,omitempty"`

	Sources []WebhookMessageDataAccount `json:"sources,omitempty"`

	EmailAccounts []WebhookMessageDataAccount `json:"email_accounts,omitempty"`

	Files []WebhookMessageDataFile `json:"files,omitempty"`

	Bodies []WebhookBody `json:"bodies,omitempty"`

	Headers mail.Header `json:"headers,omitempty"`
}

// WebhookBody embedded data struct within WebhookMessageData
type WebhookBody struct {
	Type        string `json:"type,omitempty"`
	Charset     string `json:"charset,omitempty"`
	BodySection string `json:"body_section,omitempty"`
	Content     string `json:"content,omitempty"`
}

// WebhookMessageDataFlags embedded data struct within WebhookMessageData
type WebhookMessageDataFlags struct {
	Flagged  bool `json:"flagged,omitempty"`
	Answered bool `json:"answered,omitempty"`
	Draft    bool `json:"draft,omitempty"`
	Seen     bool `json:"seen,omitempty"`
}

// WebhookMessageDataAccount embedded data struct within WebhookMessageData
type WebhookMessageDataAccount struct {
	Label       string `json:"label,omitempty"`
	Folder      string `json:"folder,omitempty"`
	UID         int    `json:"uid,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// WebhookMessageDataFile embedded data struct within WebhookMessageData
type WebhookMessageDataFile struct {
	ContentID          string `json:"content_id,omitempty"`
	Type               string `json:"type,omitempty"`
	FileName           string `json:"file_name,omitempty"`
	BodySection        string `json:"body_section,omitempty"`
	ContentDisposition string `json:"content_disposition,omitempty"`
	MainFileName       string `json:"main_file_name,omitempty"`

	XAttachmentID interface{} `json:"x_attachment_id,omitempty"`

	FileNameStructure [][]string `json:"file_name_structure,omitempty"`

	AttachmentID int `json:"attachment_id,omitempty"`
	Size         int `json:"size,omitempty"`

	IsEmbedded bool `json:"is_embedded,omitempty"`
}

// WebhookMessageDataAddresses struct within WebhookMessageData
type WebhookMessageDataAddresses struct {
	From       Address   `json:"from,omitempty"`
	To         []Address `json:"to,omitempty"`
	Cc         []Address `json:"cc,omitempty"`
	Bcc        []Address `json:"bcc,omitempty"`
	Sender     []Address `json:"sender,omitempty"`
	ReplyTo    []Address `json:"reply_to,omitempty"`
	ReturnPath []Address `json:"return_path,omitempty"`
}

// UnmarshalJSON is here because the empty state is an array in the json, and is a object/map when populated
func (m *WebhookMessageDataAddresses) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`[]`), b) {
		// its the empty array, set an empty struct
		*m = WebhookMessageDataAddresses{}
		return nil
	}
	// avoid recursion
	type webhookMessageDataAddressesTemp WebhookMessageDataAddresses
	var tmp webhookMessageDataAddressesTemp

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*m = WebhookMessageDataAddresses(tmp)
	return nil
}

// GetUserWebhooks gets listings of Webhooks configured for a user.
func (cioLite CioLite) GetUserWebhooks(userID string) ([]GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/webhooks", userID),
		UserID: userID,
	}

	// Make response
	var response []GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserWebhook gets the properties of a given Webhook.
func (cioLite CioLite) GetUserWebhook(userID string, webhookID string) (GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/webhooks/%s", userID, webhookID),
		UserID: userID,
	}

	// Make response
	var response GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserWebhook creates a new Webhook on a user.
// formValues requires CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
func (cioLite CioLite) CreateUserWebhook(userID string, formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/users/%s/webhooks", userID),
		FormValues: formValues,
		UserID:     userID,
	}

	// Make response
	var response CreateUserWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUserWebhook changes the properties of a given Webhook.
// formValues requires Active
func (cioLite CioLite) ModifyUserWebhook(userID string, webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/users/%s/webhooks/%s", userID, webhookID),
		FormValues: formValues,
		UserID:     userID,
	}

	// Make response
	var response ModifyWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserWebhookAccount cancels a Webhook.
func (cioLite CioLite) DeleteUserWebhookAccount(userID string, webhookID string) (DeleteWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/users/%s/webhooks/%s", userID, webhookID),
		UserID: userID,
	}

	// Make response
	var response DeleteWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
