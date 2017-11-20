package ciolite

import (
	"reflect"
	"testing"
)

// TestActualDiscoveryRequestToCioForGoogle tests sending an actual
// GetDiscovery request to CIO, for gmail and googleapps accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForGoogle(t *testing.T) {
	t.Parallel()

	cioLite, logger := NewTestCioLiteWithLogger(t)

	expected := GetDiscoveryResponse{
		Email: "test@gmail.com",
		Type:  "gmail",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.gmail.com",
			Username: "test@gmail.com",
			UseSSL:   true,
			OAuth:    true,
			Port:     993,
		},
	}

	// gmail
	response, err := cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@gmail.com"})
	expected.Documentation = response.Documentation

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err, "; With Log: ", logger.String())
	}

	// googleapps
	expected.Email = "test@google.com"
	expected.IMAP.Username = "test@google.com"
	expected.Type = "googleapps"
	expected.IMAP.Server = "imap.googlemail.com"

	response, err = cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@google.com"})
	expected.Documentation = response.Documentation

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}

// TestActualDiscoveryRequestToCioForMicrosoft tests sending an actual
// GetDiscovery request to CIO, for outlook and hotmail accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForMicrosoft(t *testing.T) {
	t.Parallel()

	cioLite := NewTestCioLite(t)

	expected := GetDiscoveryResponse{
		Email: "test@hotmail.com",
		Type:  "msliveconnect",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap-mail.outlook.com",
			Username: "test@hotmail.com",
			UseSSL:   true,
			OAuth:    true,
			Port:     993,
		},
	}

	// hotmail
	response, err := cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@hotmail.com"})
	expected.Documentation = response.Documentation

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}

	// outlook
	expected.Email = "test@outlook.com"
	expected.IMAP.Username = "test@outlook.com"

	response, err = cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@outlook.com"})
	expected.Documentation = response.Documentation

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForYahoo tests sending an actual
// GetDiscovery request to CIO, for yahoo accounts
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForYahoo(t *testing.T) {
	t.Parallel()

	cioLite := NewTestCioLite(t)

	expected := GetDiscoveryResponse{
		Email: "test@yahoo.com",
		Type:  "yahoo",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.mail.yahoo.com",
			Username: "test@yahoo.com",
			UseSSL:   true,
			OAuth:    false,
			Port:     993,
		},
	}

	// yahoo
	response, err := cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@yahoo.com"})
	expected.Documentation = response.Documentation

	// yahoo can somtimes be oauth, depending on the cio account settings
	if response.Type == "yahoo" && response.IMAP.OAuth {
		expected.Type = response.Type
		expected.IMAP.OAuth = response.IMAP.OAuth
	} else if response.Type == "generic" && !response.IMAP.OAuth {
		expected.Type = response.Type
		expected.IMAP.OAuth = response.IMAP.OAuth
	}

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}

	// yahoo.com.cn
	expected.Email = "test@yahoo.com.cn"
	expected.IMAP.Username = "test@yahoo.com.cn"

	response, err = cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@yahoo.com.cn"})
	expected.Documentation = response.Documentation

	// yahoo can somtimes be oauth, depending on the cio account settings
	if response.Type == "yahoo" && response.IMAP.OAuth {
		expected.Type = response.Type
		expected.IMAP.OAuth = response.IMAP.OAuth
	} else if response.Type == "generic" && !response.IMAP.OAuth {
		expected.Type = response.Type
		expected.IMAP.OAuth = response.IMAP.OAuth
	}

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForAol tests sending an actual
// GetDiscovery request to CIO, for an AOL account
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForAol(t *testing.T) {
	t.Parallel()

	cioLite := NewTestCioLite(t)

	expected := GetDiscoveryResponse{
		Email: "test@aol.com",
		Type:  "aol",
		Found: true,
		IMAP: GetDiscoveryIMAPResponse{
			Server:   "imap.aol.com",
			Username: "test@aol.com",
			UseSSL:   true,
			OAuth:    false,
			Port:     993,
		},
	}

	// aol
	response, err := cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@aol.com"})
	expected.Documentation = response.Documentation

	if err != nil || !reflect.DeepEqual(expected, response) {
		t.Error("Expected GetDiscovery Response: ", expected, "; Got: ", response, "; With Error: ", err)
	}
}

// TestActualDiscoveryRequestToCioForNonExistent tests sending an actual
// GetDiscovery request to CIO, for a non-existent email service provider
// (internet connection required, real CIO key/secret required).
func TestActualDiscoveryRequestToCioForNonExistent(t *testing.T) {
	t.Parallel()

	cioLite := NewTestCioLite(t)

	// not found
	response, err := cioLite.GetDiscovery(GetDiscoveryParams{Email: "test@bogusblahblahfoobar.com"})

	if err != ErrNotFound {
		t.Error("Expected GetDiscovery Error ErrNotFound; Got: ", response, "; With Error: ", err)
	}
}
