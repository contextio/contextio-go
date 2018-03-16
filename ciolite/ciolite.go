// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"
)

const (
	// DefaultHost is the default host of CIO Lite API
	DefaultHost = "https://api.context.io"

	// DefaultRequestTimeout is the default timeout duration used on HTTP requests
	DefaultRequestTimeout = 120 * time.Second
)

// CioLite struct contains the api key and secret, along with an optional logger,
// and provides convenience functions for accessing all CIO Lite endpoints.
type CioLite struct {
	apiKey    string
	apiSecret string
	Host      string

	// Allow setting your own *http.Client, otherwise default is client with DefaultRequestTimeout
	HTTPClient *http.Client

	// PreRequestHook is a function (mostly for logging) that will be executed
	// before the request is made.
	// 	Its arguments are:
	// 	User ID (if present),
	// 	Account Label (if present),
	// 	Method (GET/POST/etc),
	// 	URL,
	// 	redacted body values.
	PreRequestHook func(string, string, string, string, url.Values)

	// PostRequestShouldRetryHook is a function (mostly for logging) that will be
	// executed after each request is made, and will be called at least once.
	// 	Its arguments are:
	// 	request Attempt # (starts at 1),
	// 	User ID (if present),
	// 	Account Label (if present),
	// 	Method (GET/POST/etc),
	// 	URL,
	// 	response Status Code,
	// 	response Payload,
	// 	time at start of most recent attempt,
	// 	time at start of all attempts,
	// 	any error received while attempting this request.
	// The returned boolean is whether this request should be retried or not, which
	// if False then this is the last call of this function, but if True means this
	// function will be called again.
	PostRequestShouldRetryHook func(int, string, string, string, string, int, string, time.Time, time.Time, error) bool

	// ResponseBodyCloseErrorHook is a function (purely for logging) that will
	// execute if there is an error closing the response body.
	ResponseBodyCloseErrorHook func(error)
}

// NewCioLite returns a CIO Lite struct (without a logger) for accessing the CIO Lite API.
func NewCioLite(key string, secret string) CioLite {

	return CioLite{
		apiKey:     key,
		apiSecret:  secret,
		Host:       DefaultHost,
		HTTPClient: &http.Client{Timeout: DefaultRequestTimeout},
	}
}

// Interface is just to help generate a mocked client, for testing elsewhere.
// mockgen -source=ciolite.go -destination=ciolite_mock.go -package ciolite
type Interface interface {
	ValidateCallback(token string, signature string, timestamp int) bool

	GetStatusCallbackURL() (GetStatusCallbackURLResponse, error)
	CreateStatusCallbackURL(formValues CreateStatusCallbackURLParams) (CreateDeleteStatusCallbackURLResponse, error)
	DeleteStatusCallbackURL() (CreateDeleteStatusCallbackURLResponse, error)

	GetConnectTokens() ([]GetConnectTokenResponse, error)
	GetConnectToken(token string) (GetConnectTokenResponse, error)
	CreateConnectToken(formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error)
	DeleteConnectToken(token string) (DeleteConnectTokenResponse, error)
	CheckConnectToken(connectToken GetConnectTokenResponse, email string) error

	GetDiscovery(queryValues GetDiscoveryParams) (GetDiscoveryResponse, error)

	GetOAuthProviders() ([]GetOAuthProvidersResponse, error)
	GetOAuthProvider(key string) (GetOAuthProvidersResponse, error)
	CreateOAuthProvider(formValues CreateOAuthProviderParams) (CreateOAuthProviderResponse, error)
	DeleteOAuthProvider(key string) (DeleteOAuthProviderResponse, error)

	GetUserConnectTokens(userID string) ([]GetConnectTokenResponse, error)
	GetUserConnectToken(userID string, token string) (GetConnectTokenResponse, error)
	CreateUserConnectToken(userID string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error)
	DeleteUserConnectToken(userID string, token string) (DeleteConnectTokenResponse, error)

	GetUserEmailAccountConnectTokens(userID string, label string) ([]GetConnectTokenResponse, error)
	GetUserEmailAccountConnectToken(userID string, label string, token string) (GetConnectTokenResponse, error)
	CreateUserEmailAccountConnectToken(userID string, label string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error)
	DeleteUserEmailAccountConnectToken(userID string, label string, token string) (DeleteConnectTokenResponse, error)

	GetUserEmailAccountsFolderMessageAttachments(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) ([]GetUserEmailAccountsFolderMessageAttachmentsResponse, error)
	GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error)
	GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageBodyParams) ([]GetUserEmailAccountsFolderMessageBodyResponse, error)
	GetUserEmailAccountsFolderMessageFlags(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageFlagsResponse, error)

	GetUserEmailAccountsFolderMessageHeaders(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageHeadersParams) (GetUserEmailAccountsFolderMessageHeadersResponse, error)
	GetUserEmailAccountsFolderMessageRaw(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageRawResponse, error)
	MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error)
	MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error)

	GetUserEmailAccountsFolderMessages(userID string, label string, folder string, queryValues GetUserEmailAccountsFolderMessageParams) ([]GetUsersEmailAccountFolderMessagesResponse, error)
	GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageParams) (GetUsersEmailAccountFolderMessagesResponse, error)
	MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues MoveUserEmailAccountFolderMessageParams) (MoveUserEmailAccountFolderMessageResponse, error)

	GetUserEmailAccountsFolders(userID string, label string, queryValues GetUserEmailAccountsFoldersParams) ([]GetUsersEmailAccountFoldersResponse, error)
	GetUserEmailAccountFolder(userID string, label string, folder string, queryValues EmailAccountFolderDelimiterParam) (GetUsersEmailAccountFoldersResponse, error)
	CreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (CreateEmailAccountFolderResponse, error)
	SafeCreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (bool, error)

	GetUserEmailAccounts(userID string, queryValues GetUserEmailAccountsParams) ([]GetUsersEmailAccountsResponse, error)
	GetUserEmailAccount(userID string, label string) (GetUsersEmailAccountsResponse, error)
	CreateUserEmailAccount(userID string, formValues CreateUserParams) (CreateEmailAccountResponse, error)
	ModifyUserEmailAccount(userID string, label string, formValues ModifyUserEmailAccountParams) (ModifyEmailAccountResponse, error)
	DeleteUserEmailAccount(userID string, label string) (DeleteEmailAccountResponse, error)

	GetUserWebhooks(userID string) ([]GetUsersWebhooksResponse, error)
	GetUserWebhook(userID string, webhookID string) (GetUsersWebhooksResponse, error)
	CreateUserWebhook(userID string, formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error)
	ModifyUserWebhook(userID string, webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error)
	DeleteUserWebhookAccount(userID string, webhookID string) (DeleteWebhookResponse, error)

	GetUsers(queryValues GetUsersParams) ([]GetUsersResponse, error)
	GetUser(userID string) (GetUsersResponse, error)
	CreateUser(formValues CreateUserParams) (CreateUserResponse, error)
	ModifyUser(userID string, formValues ModifyUserParams) (ModifyUserResponse, error)
	DeleteUser(userID string) (DeleteUserResponse, error)

	GetWebhooks() ([]GetUsersWebhooksResponse, error)
	GetWebhook(webhookID string) (GetUsersWebhooksResponse, error)
	CreateWebhook(formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error)
	ModifyWebhook(webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error)
	DeleteWebhookAccount(webhookID string) (DeleteWebhookResponse, error)
}

// NewTestCioLiteServer is a convenience function that returns a CioLite object
// and a *httptest.Server (which must be closed when done being used).
// The CioLite instance will hit the test server for all requests.
func NewTestCioLiteServer(handler http.Handler) (CioLite, *httptest.Server) {
	testServer := httptest.NewServer(handler)
	testCioLite := CioLite{
		Host:       testServer.URL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
	return testCioLite, testServer
}

// ValidateCallback returns true if this Webhook Callback or User Account Status Callback authenticates
func (cio CioLite) ValidateCallback(token string, signature string, timestamp int) bool {
	// Hash timestamp and token with secret, compare to signature
	message := strconv.Itoa(timestamp) + token
	hash := hashHmac(sha256.New, message, cio.apiSecret)
	return len(hash) > 0 && signature == hash
}

// hashHmac returns the hash of a message hashed with the provided hash function, using the provided secret
func hashHmac(hashAlgorithm func() hash.Hash, message string, secret string) string {
	h := hmac.New(hashAlgorithm, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		panic("hash.Hash unable to write message bytes, with error: " + err.Error())
	}
	return hex.EncodeToString(h.Sum(nil))
}
