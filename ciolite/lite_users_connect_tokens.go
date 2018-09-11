package ciolite

// Api functions that support: users/connect_tokens

import (
	"fmt"
)

// GetUserConnectTokens gets a list of connect tokens created for a user.
func (cioLite CioLite) GetUserConnectTokens(userID string) ([]GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/connect_tokens", userID),
		UserID: userID,
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserConnectToken gets information about a given connect token for a specific user.
func (cioLite CioLite) GetUserConnectToken(userID string, token string) (GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/users/%s/connect_tokens/%s", userID, token),
		UserID: userID,
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserConnectToken creates and obtains a new connect_token for a specific user.
// formValues requires CallbackURL, and may optionally have
// Email, FirstName, LastName, StatusCallbackURL
func (cioLite CioLite) CreateUserConnectToken(userID string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/users/%s/connect_tokens", userID),
		FormValues: formValues,
		UserID:     userID,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserConnectToken removes a given connect token for a specific user.
func (cioLite CioLite) DeleteUserConnectToken(userID string, token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/users/%s/connect_tokens/%s", userID, token),
		UserID: userID,
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
