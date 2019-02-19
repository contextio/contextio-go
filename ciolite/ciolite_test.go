package ciolite

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"
)

// TestNewCioLiteWithLogger tests the construction of CioLite
func TestNewCioLite(t *testing.T) {
	t.Parallel()
	NewTestCioLite(t)
}

// TestNewCioLiteWithLogger tests the construction of CioLite and *TestLogger objects
func TestNewTestCioLiteServer(t *testing.T) {
	t.Parallel()
	_, _, testServer, _ := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()
}

// NewTestCioLite returns a new CioLite object
func NewTestCioLite(t *testing.T) CioLite {
	return NewCioLite(getEnv(t, "CIO_API_KEY"), getEnv(t, "CIO_API_SECRET"), false)
}

// NewTestCioLiteWithLogger returns a new CioLite object and *TestLogger object
func NewTestCioLiteWithLogger(t *testing.T) (CioLite, *TestLogger) {
	cioLite := NewTestCioLite(t)
	return cioLite, addLogging(&cioLite)
}

// NewTestCioLiteWithLoggerAndTestServer returns a new CioLite, *TestLogger, and *httptest.Server objects
func NewTestCioLiteWithLoggerAndTestServer(t *testing.T) (CioLite, *TestLogger, *httptest.Server, *http.ServeMux) {
	mux := http.NewServeMux()
	cioLite, server := NewTestCioLiteServer(mux)
	return cioLite, addLogging(&cioLite), server, mux
}

func addLogging(cioLite *CioLite) *TestLogger {

	logger := &TestLogger{Buffer: &bytes.Buffer{}}

	// Implement logging in the hooks
	cioLite.ResponseBodyCloseErrorHook = func(err error) {
		logger.Printf("Unable to close response body from CIO, with error: %s\n", err.Error())
	}

	cioLite.PreRequestHook = func(userID string, label string, method string, url string, redactedBodyValues url.Values) {
		logger.Printf("Creating new %s request to: %s with payload: %s\n", method, url, redactedBodyValues.Encode())
	}

	cioLite.PostRequestShouldRetryHook = func(attemptNum int, userID string, label string, method string, url string, statusCode int, responseBody string, beforeAttempt time.Time, beforeAll time.Time, err error) bool {
		// TODO: redact access_token and access_token_secret inside resBody before logging (only occurs with 3-legged oauth (not presently used))
		// Take only the first 2000 characters from the responseBody, which should be more than enough to debug anything, without killing the logger
		if bodyLen := len(responseBody); bodyLen > 2000 {
			responseBody = responseBody[:2000]
		}
		if err != nil {
			logger.Printf("Received response from %s to: %s with status code: %d and error: %s and payload snippet: %s\n", method, url, statusCode, err, responseBody)
		} else {
			logger.Printf("Received response from %s to: %s with status code: %d and payload snippet: %s\n", method, url, statusCode, responseBody)
		}
		return false // Do not retry the request
	}

	return logger
}

// getEnv returns the named environment variable, or causes t.Fatal
func getEnv(t *testing.T, envVarName string) string {
	val := os.Getenv(envVarName)
	if len(val) == 0 {
		t.Fatal("Empty Environment Variable for: " + envVarName)
	}
	return val
}

// Must panics if error is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// TestLogger is a *bytes.Buffer that implements the logging interface
type TestLogger struct {
	*bytes.Buffer
}

// Printf prints the arguments to the buffer, using fmt.Sprintf
func (l *TestLogger) Printf(format string, v ...interface{}) {
	_, err := l.Write([]byte(fmt.Sprintf(format, v...)))
	if err != nil {
		panic("Error writing to test logger: " + err.Error())
	}
}
