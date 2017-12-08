package ciolite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// RequestError is the error type returned by DoFormRequest and all cio api calls
type RequestError struct {
	Err error
	ErrorMetaData
}

// ErrorMetaData holds some meta-data about the error: StatusCode, Response Payload, Method used, and URL
type ErrorMetaData struct {
	StatusCode int
	Payload    string
	Method     string
	URL        string
}

// ErrorStatusCode returns the RequestError's StatusCode (ex: 200 for OK, 0 if no status code)
func (e RequestError) ErrorStatusCode() int {
	return e.StatusCode
}

// ErrorPayload returns the RequestError's payload, if present
func (e RequestError) ErrorPayload() string {
	return e.Payload
}

// ErrorMethod returns the RequestError's Method (ex: GET, POST, etc)
func (e RequestError) ErrorMethod() string {
	return e.Method
}

// ErrorURL returns the RequestError's URL
func (e RequestError) ErrorURL() string {
	return e.URL
}

// Cause returns the cause of any wrapped errors, or just the base error if no wrapped error.
// Can use with github.com/pkg/errors
func (e RequestError) Cause() error {
	return errors.Cause(e.Err)
}

// Format prints out the error, any causes, a stacktrace, and the other fields in the struct
func (e RequestError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n%+v", e.Err, e.ErrorMetaData)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

// Error returns the Error string, any Wrapped Causes, and any StatusCode, Payload, Method, and URL that were set
func (e RequestError) Error() string {
	return fmt.Sprintf("%s; %+v", e.Err, e.ErrorMetaData)
}

// String returns the same as Error()
func (e RequestError) String() string {
	return e.Error()
}

// MarshalJSON allows RequestError to implement json.Marshaler (for use with logging in json)
func (e RequestError) MarshalJSON() ([]byte, error) {
	type Temp struct {
		Err string
		ErrorMetaData
	}
	// Marshall
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(Temp{e.Err.Error(), e.ErrorMetaData})
	return bytes.Trim(buffer.Bytes(), " \r\n"), err
}

// UnmarshalJSON allows RequestError to implement json.Unmarshaler (for completeness since RequestError implements Marshaler).
// 	Note that if the json property is possibly null, you must unmarshal to *RequestError, as the cioutil.UnmarshalJSON helper does.
// 	Use of this loses the ability to unwrap errors with errors.Cause() or get the original stacktrace
func (e *RequestError) UnmarshalJSON(data []byte) error {
	type Temp struct {
		Err string
		ErrorMetaData
	}
	var re Temp
	err := json.Unmarshal(data, &re)
	e.Err = errors.New(re.Err)
	e.ErrorMetaData = re.ErrorMetaData
	return err
}

// UnmarshalJSON is a helper method to unmarshal the json representation of a RequestError,
// getting back a valid nil error if it should be nil, and a valid RequestError error otherwise.
// 	Pay attention to return value order: First argument is the unmarshalled RequestError as an error interface,
// 	Second argument is any actual error encountered while unmarshalling (from json.Unmarshal).
func UnmarshalJSON(data []byte) (error, error) {
	var unmarshalled *RequestError
	jsonErr := json.Unmarshal(data, &unmarshalled)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if unmarshalled == nil {
		return nil, nil
	}
	return *unmarshalled, nil
}
