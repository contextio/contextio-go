package ciolite

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func nilError() error {
	return nil
}

func newError() error {
	return RequestError{fmt.Errorf("fmt %s", "error"), ErrorMetaData{StatusCode: 400, Payload: "new", Method: "POST", URL: "http://cio"}}
}

func wrappedError() error {
	return RequestError{errors.Wrap(fmt.Errorf("cause %s", "error"), "outer error"), ErrorMetaData{StatusCode: 500, Payload: "wrapped", Method: "PUT", URL: "https://cio"}}
}

// TestRequestErrorWrapped tests the behavior of a RequestError that wrapped a causing error
func TestRequestErrorWrapped(t *testing.T) {
	t.Parallel()

	err := wrappedError()
	if err == nil {
		t.Fatal("Expected non-nil error; Got: ", err)
	}

	expectedSuffix := "{StatusCode:500 Payload:wrapped Method:PUT URL:https://cio}"
	if err.Error() != ("outer error: cause error; " + expectedSuffix) {
		t.Error("Expected error string of: ", "outer error: cause error", "; Got: ", err.Error())
	}

	if errors.Cause(err).Error() != "cause error" {
		t.Error("Expected error cause string of: ", "cause error", "; Got: ", err.Error())
	}

	expectedPrefix := "cause error\nouter error\ngithub.com/contextio/contextio-go/ciolite.wrappedError\n"
	if plusV := fmt.Sprintf("%+v", err); !strings.HasPrefix(plusV, expectedPrefix) || !strings.HasSuffix(plusV, expectedSuffix) {
		t.Error("Expected +v formatting of: ", expectedPrefix, expectedSuffix, "; Got: ", plusV)
	}

	if code := ErrorStatusCode(err); code != 500 {
		t.Error("Expected error status code of: ", 500, "; Got: ", code)
	}

	if val := ErrorPayload(err); val != "wrapped" {
		t.Error("Expected error payload of: ", "wrapped", "; Got: ", val)
	}

	if val := ErrorMethod(err); val != "PUT" {
		t.Error("Expected error method of: ", "PUT", "; Got: ", val)
	}

	if val := ErrorURL(err); val != "https://cio" {
		t.Error("Expected error url of: ", "https://cio", "; Got: ", val)
	}

	_, ok := err.(RequestError)
	if !ok {
		t.Error("Expected error to be of type: ", "RequestError", "; Got: ", err)
	}

	jsonBytes, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		t.Error(jsonErr)
	}

	if string(jsonBytes) != `{"Err":"outer error: cause error","StatusCode":500,"Payload":"wrapped","Method":"PUT","URL":"https://cio"}` {
		t.Error("Expected json to be: ",
			`{"Error":"outer error: cause error","StatusCode":500,"Payload":"wrapped","Method":"PUT","URL":"https://cio"}`,
			"; Got: ",
			string(jsonBytes),
		)
	}

	if unmarshalled, libErr := UnmarshalJSON(jsonBytes); libErr != nil {
		t.Error(libErr)

	} else if !reflect.DeepEqual(unmarshalled.(RequestError).ErrorMetaData, err.(RequestError).ErrorMetaData) {
		t.Error("Expected unmarshalled json error meta data to be: ", err.(RequestError).ErrorMetaData, "; Got: ", unmarshalled.(RequestError).ErrorMetaData)

	} else if unmarshalled.(RequestError).Err.Error() != err.(RequestError).Err.Error() {
		t.Error("Expected unmarshalled json error string to be: ", err.(RequestError).Err, "; Got: ", unmarshalled.(RequestError).Err)
	}
}

// TestRequestErrorNew tests the behavior of a newly created RequestError type
func TestRequestErrorNew(t *testing.T) {
	t.Parallel()

	err := newError()
	if err == nil {
		t.Fatal("Expected non-nil error; Got: ", err)
	}

	expectedSuffix := "{StatusCode:400 Payload:new Method:POST URL:http://cio}"
	if err.Error() != ("fmt error; " + expectedSuffix) {
		t.Error("Expected error string of: ", "fmt error", "; Got: ", err.Error())
	}

	if errors.Cause(err).Error() != "fmt error" {
		t.Error("Expected error cause string of: ", "fmt error", "; Got: ", err.Error())
	}

	expectedPlusV := "fmt error\n" + expectedSuffix
	if plusV := fmt.Sprintf("%+v", err); plusV != expectedPlusV {
		t.Error("Expected +v formatting of: ", expectedPlusV, "; Got: ", plusV)
	}

	if code := ErrorStatusCode(err); code != 400 {
		t.Error("Expected error status code of: ", 400, "; Got: ", code)
	}

	if val := ErrorPayload(err); val != "new" {
		t.Error("Expected error payload of: ", "new", "; Got: ", val)
	}

	if val := ErrorMethod(err); val != "POST" {
		t.Error("Expected error method of: ", "POST", "; Got: ", val)
	}

	if val := ErrorURL(err); val != "http://cio" {
		t.Error("Expected error url of: ", "http://cio", "; Got: ", val)
	}

	_, ok := err.(RequestError)
	if !ok {
		t.Error("Expected error to be of type: ", "RequestError", "; Got: ", err)
	}

	jsonBytes, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		t.Error(jsonErr)
	}

	if string(jsonBytes) != `{"Err":"fmt error","StatusCode":400,"Payload":"new","Method":"POST","URL":"http://cio"}` {
		t.Error("Expected json to be: ",
			`{"Error":"fmt error","StatusCode":400,"Payload":"new","Method":"POST","URL":"http://cio"}`,
			"; Got: ",
			string(jsonBytes),
		)
	}

	if unmarshalled, libErr := UnmarshalJSON(jsonBytes); libErr != nil {
		t.Error(libErr)

	} else if !reflect.DeepEqual(unmarshalled.(RequestError).ErrorMetaData, err.(RequestError).ErrorMetaData) {
		t.Error("Expected unmarshalled json error meta data to be: ", err.(RequestError).ErrorMetaData, "; Got: ", unmarshalled.(RequestError).ErrorMetaData)

	} else if unmarshalled.(RequestError).Err.Error() != err.(RequestError).Err.Error() {
		t.Error("Expected unmarshalled json error string to be: ", err.(RequestError).Err, "; Got: ", unmarshalled.(RequestError).Err)
	}
}

// TestRequestErrorNil tests the nil behavior of the error helpers
func TestRequestErrorNil(t *testing.T) {
	t.Parallel()

	err := nilError()

	if err != nil {
		t.Error("Expected nil error; Got: ", err)
	}

	if code := ErrorStatusCode(err); code != 0 {
		t.Error("Expected error status code of: ", 0, "; Got: ", code)
	}

	if val := ErrorPayload(err); val != "" {
		t.Error("Expected error payload of: ", "", "; Got: ", val)
	}

	if val := ErrorMethod(err); val != "" {
		t.Error("Expected error method of: ", "", "; Got: ", val)
	}

	if val := ErrorURL(err); val != "" {
		t.Error("Expected error url of: ", "", "; Got: ", val)
	}

	jsonBytes, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		t.Error(jsonErr)
	}

	if string(jsonBytes) != `null` {
		t.Error("Expected json to be: ", `null`, "; Got: ", string(jsonBytes))
	}

	if unmarshalled, libErr := UnmarshalJSON(jsonBytes); libErr != nil {
		t.Error(libErr)

	} else if unmarshalled != nil {
		t.Error("Expected unmarshalled json error to be: nil; Got: ", unmarshalled)
	}
}
