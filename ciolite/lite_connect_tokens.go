package ciolite

// Api functions that support: connect_tokens

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// GetConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#get
// 	https://context.io/docs/lite/connect_tokens#id-get
type GetConnectTokenResponse struct {
	Token              string `json:"token,omitempty"`
	Email              string `json:"email,omitempty"`
	EmailAccountID     string `json:"email_account_id,omitempty"`
	CallbackURL        string `json:"callback_url,omitempty"`
	StatusCallbackURL  string `json:"status_callback_url,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"` // TODO: reconfirm which fields the response includes
	ServerLabel        string `json:"server_label,omitempty"`

	AccountLite bool `json:"account_lite,omitempty"`

	Created int `json:"created,omitempty"`
	Used    int `json:"used,omitempty"`

	Expires ExpiresMixed `json:"expires,omitempty"`

	User GetConnectTokenUserResponse `json:"user,omitempty"`
}

// GetConnectTokenUserResponse data struct within GetConnectTokenResponse
// 	https://context.io/docs/lite/connect_tokens#get
// 	https://context.io/docs/lite/connect_tokens#id-get
type GetConnectTokenUserResponse struct {
	ID             string   `json:"id,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	Created        int      `json:"created,omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`
}

// CreateConnectTokenParams form values data struct.
// Requires CallbackURL, and optionally may have
// Email, FirstName, LastName, StatusCallbackURL.
// 	https://context.io/docs/lite/connect_tokens#post
// 	https://context.io/docs/lite/users/connect_tokens#post
type CreateConnectTokenParams struct {
	// Required:
	CallbackURL string `json:"callback_url"`

	// Optional:
	Email             string `json:"email,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
}

// CreateConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#post
type CreateConnectTokenResponse struct {
	Success            bool   `json:"success,omitempty"`
	Token              string `json:"token,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"`
	AccessToken        string `json:"access_token,omitempty"`
	AccessTokenSecret  string `json:"access_token_secret,omitempty"`
}

// DeleteConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#id-delete
type DeleteConnectTokenResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetConnectTokens get a list of connect tokens created with your API key.
// 	https://context.io/docs/lite/connect_tokens#get
func (cioLite CioLite) GetConnectTokens() ([]GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/lite/connect_tokens",
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetConnectToken gets information about a given connect token.
// 	https://context.io/docs/lite/connect_tokens#id-get
func (cioLite CioLite) GetConnectToken(token string) (GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/connect_tokens/%s", token),
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateConnectToken creates and obtains a new connect token.
// formValues requires CallbackURL, and optionally may have
// Email, FirstName, LastName, StatusCallbackURL
// 	https://context.io/docs/lite/connect_tokens#post
func (cioLite CioLite) CreateConnectToken(formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/lite/connect_tokens",
		FormValues: formValues,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteConnectToken removes a given connect token
// 	https://context.io/docs/lite/connect_tokens#id-delete
func (cioLite CioLite) DeleteConnectToken(token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/connect_tokens/%s", token),
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// EmailAccountMatching searches its EmailAccounts array for the provided email address,
// and returns the GetUsersEmailAccountsResponse Email Account that matches it.
func (user GetConnectTokenUserResponse) EmailAccountMatching(email string) (GetUsersEmailAccountsResponse, error) {
	return FindEmailAccountMatching(user.EmailAccounts, email)
}

// CheckConnectToken checks and returns nil if the connect token was used, the email
// authorized matches the expected email, and that CIO has access to the account.
func (cioLite CioLite) CheckConnectToken(connectToken GetConnectTokenResponse, email string) error {

	// Confirm email matches
	if strings.ToLower(connectToken.Email) != strings.ToLower(email) {
		return errors.New("Email does not match Context.io token")
	}

	// Confirm token was used (accepted/authorized)
	if connectToken.Used == 0 || connectToken.Expires.Unused() {
		return errors.New("Context.io token not used yet")
	}

	// Confirm user exists
	if len(connectToken.User.ID) == 0 {
		return errors.New("Context.io user not created yet")
	}

	// Confirm we have access
	account, err := connectToken.User.EmailAccountMatching(email)
	if err != nil || account.Status != "OK" {
		return errors.New("Unable to access account using Context.io")
	}

	return nil
}

// ExpiresMixed is a special type to handle the fact that 'expires' can be an int or false.
// 	Unix timestamp of when this token will expire and be purged.
// 	Once the token is used, this property will be set to false
// 	https://context.io/docs/lite/users/connect_tokens
type ExpiresMixed struct {
	Expires *int
}

// Unused returns true if Expires is a Unix timestamp,
// and returns false if this token has been used.
func (expires *ExpiresMixed) Unused() bool {
	return expires.Expires != nil
}

// Timestamp returns the expires timestamp if the token is unused,
// and returns -1 if the token has been used.
func (expires *ExpiresMixed) Timestamp() int {
	if expires.Expires == nil {
		return -1
	}
	return *expires.Expires
}

// MarshalJSON allows ExpiresMixed to implement json.Marshaler
func (expires ExpiresMixed) MarshalJSON() ([]byte, error) {
	if expires.Expires == nil {
		return []byte(`false`), nil
	}
	return []byte(strconv.Itoa(*expires.Expires)), nil
}

// UnmarshalJSON allows ExpiresMixed to implement json.Unmarshaler
func (expires *ExpiresMixed) UnmarshalJSON(data []byte) error {
	stringData := string(data)
	if strings.ToLower(stringData) == "false" {
		return nil
	}
	expiresInt, err := strconv.Atoi(stringData)
	expires.Expires = &expiresInt
	return err
}
