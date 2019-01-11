package ciolite

// Api functions that support: users/email_accounts

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// GetUserEmailAccountsParams query values data struct.
// Optional: Status, StatusOK
type GetUserEmailAccountsParams struct {
	// Optional:
	Status   string `json:"status,omitempty"`
	StatusOK string `json:"status_ok,omitempty"`
}

// GetUsersEmailAccountsResponse data struct
type GetUsersEmailAccountsResponse struct {
	Status               string `json:"status,omitempty"`
	ResourceURL          string `json:"resource_url,omitempty"`
	Type                 string `json:"type,omitempty"`
	AuthenticationType   string `json:"authentication_type,omitempty"`
	Server               string `json:"server,omitempty"`
	Label                string `json:"label,omitempty"`
	Username             string `json:"username,omitempty"`
	MailServiceAccountID string `json:"mailservice_account_id,omitempty"`

	UseSSL bool `json:"use_ssl,omitempty"`

	Port int `json:"port,omitempty"`
}

// CreateEmailAccountResponse data struct
type CreateEmailAccountResponse struct {
	Status               string `json:"status,omitempty"`
	Label                string `json:"label,omitempty"`
	ResourceURL          string `json:"resource_url,omitempty"`
	MailServiceAccountID string `json:"mailservice_account_id,omitempty"`
}

// ModifyUserEmailAccountParams form values data struct.
// formValues optionally may contain Status, ForceStatusCheck, Password,
// ProviderRefreshToken, ProviderConsumerKey, StatusCallbackURL
type ModifyUserEmailAccountParams struct {
	// Optional:
	Status               string `json:"status,omitempty"`
	Password             string `json:"password,omitempty"`
	ProviderRefreshToken string `json:"provider_refresh_token,omitempty"`
	ProviderConsumerKey  string `json:"provider_consumer_key,omitempty"`
	StatusCallbackURL    string `json:"status_callback_url,omitempty"`
	ForceStatusCheck     bool   `json:"force_status_check,omitempty"`
}

// ModifyEmailAccountResponse data struct
type ModifyEmailAccountResponse struct {
	Success       bool   `json:"success,omitempty"`
	ResourceURL   string `json:"resource_url,omitempty"`
	FeedbackCode  string `json:"feedback_code,omitempty"`
	ConnectionLog string `json:"connection_log,omitempty"`
}

// DeleteEmailAccountResponse data struct
type DeleteEmailAccountResponse struct {
	Success      bool   `json:"success,omitempty"`
	ResourceURL  string `json:"resource_url,omitempty"`
	FeedbackCode string `json:"feedback_code,omitempty"`
}

// StatusCallback data struct that will be received from CIO when a user's status changes
type StatusCallback struct {
	AccountID      string `json:"account_id,omitempty"`
	UserID         string `json:"user_id,omitempty"`
	ServerLabel    string `json:"server_label,omitempty"`
	EmailAccount   string `json:"email_account,omitempty"`
	Failure        string `json:"failure,omitempty"`
	FailureMessage string `json:"failure_message,omitempty"`

	Token     string `json:"token,omitempty" valid:"required"`
	Signature string `json:"signature,omitempty" valid:"required"`

	Timestamp int `json:"timestamp,omitempty" valid:"required"`
}

// GetUserEmailAccounts gets a list of email accounts assigned to a user.
// queryValues may optionally contain Status, StatusOK
func (cioLite CioLite) GetUserEmailAccounts(userID string, queryValues GetUserEmailAccountsParams) ([]GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/lite/users/%s/email_accounts", userID),
		QueryValues: queryValues,
		UserID:      userID,
	}

	// Make response
	var response []GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccount gets the parameters and status for an email account.
// Status can be one of: OK, CONNECTION_IMPOSSIBLE, INVALID_CREDENTIALS, TEMP_DISABLED, DISABLED
func (cioLite CioLite) GetUserEmailAccount(userID string, label string) (GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s", userID, label),
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccount adds a mailbox to a given user.
// formValues requires Email, Server, Username, UseSSL, Port, Type,
// and (if OAUTH) ProviderRefreshToken and ProviderConsumerKey,
// and (if not OAUTH) Password,
// and may optionally contain StatusCallbackURL
func (cioLite CioLite) CreateUserEmailAccount(userID string, formValues CreateUserParams) (CreateEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/users/%s/email_accounts", userID),
		FormValues: formValues,
		UserID:     userID,
	}

	// Make response
	var response CreateEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUserEmailAccount modifies an email account on a given user.
// formValues optionally may contain Status, ForceStatusCheck, Password,
// ProviderRefreshToken, ProviderConsumerKey, StatusCallbackURL
func (cioLite CioLite) ModifyUserEmailAccount(userID string, label string, formValues ModifyUserEmailAccountParams) (ModifyEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "POST",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s", userID, label),
		FormValues:   formValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response ModifyEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserEmailAccount deletes an email account of a user.
func (cioLite CioLite) DeleteUserEmailAccount(userID string, label string) (DeleteEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "DELETE",
		Path:         fmt.Sprintf("/lite/users/%s/email_accounts/%s", userID, label),
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response DeleteEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// FindEmailAccountMatching searches an array of GetUsersEmailAccountsResponse's
// for the one that matches the provided email address, and returns it.
func FindEmailAccountMatching(emailAccounts []GetUsersEmailAccountsResponse, email string) (GetUsersEmailAccountsResponse, error) {

	if emailAccounts != nil {

		// Try to match against the username
		for _, emailAccount := range emailAccounts {
			if strings.ToLower(email) == strings.ToLower(emailAccount.Username) {
				return emailAccount, nil
			}
		}

		// Try to match the local part against the username or label
		localPart := strings.ToLower(upToSeparator(email, "@"))

		for _, emailAccount := range emailAccounts {

			if localPart == strings.ToLower(upToSeparator(emailAccount.Username, "@")) ||
				localPart == strings.ToLower(upToSeparator(emailAccount.Label, ":")) {

				return emailAccount, nil
			}
		}
	}
	return GetUsersEmailAccountsResponse{}, errors.Errorf("No email accounts match %s in %v", email, emailAccounts)
}

// upToSeparator returns a string up to the separator, or the whole string
// if the separator is not contained in the string
func upToSeparator(s string, sep string) string {
	idx := strings.Index(s, sep)
	if idx >= 0 {
		return s[:idx]
	}
	return s
}
