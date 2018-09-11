package ciolite

// Api functions that support: users/email_accounts/connect_tokens

import (
	"fmt"
)

// GetUserEmailAccountConnectTokens gets a list of connect tokens created for a user email account.
func (cioLite CioLite) GetUserEmailAccountConnectTokens(userID string, label string) ([]GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/email_accounts/%s/connect_tokens", userID, label),
		UserID: userID,
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountConnectToken gets information about a given connect token for a specific user email account.
func (cioLite CioLite) GetUserEmailAccountConnectToken(userID string, label string, token string) (GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/email_accounts/%s/connect_tokens/%s", userID, label, token),
		UserID: userID,
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccountConnectToken creates and obtains a new connect_token for a specific user email account.
// formValues requires CallbackURL
func (cioLite CioLite) CreateUserEmailAccountConnectToken(userID string, label string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/users/%s/email_accounts/%s/connect_tokens", userID, label),
		FormValues: formValues,
		UserID:     userID,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserEmailAccountConnectToken removes a given connect token for a specific user email account.
func (cioLite CioLite) DeleteUserEmailAccountConnectToken(userID string, label string, token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/users/%s/email_accounts/%s/connect_tokens/%s", userID, label, token),
		UserID: userID,
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
