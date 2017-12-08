package ciolite

import (
	"bytes"
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

func htmlError() error {
	return RequestError{fmt.Errorf("status >= 400 & < 9000"), ErrorMetaData{StatusCode: 401, Payload: "<html>&amp;</html>", Method: "GET", URL: "https://cio"}}
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

// TestRequestErrorHTML tests the behavior of a RequestError that contains characters that might get escaped in json
func TestRequestErrorHTML(t *testing.T) {
	t.Parallel()

	err := htmlError()
	if err == nil {
		t.Fatal("Expected non-nil error; Got: ", err)
	}

	expectedSuffix := "{StatusCode:401 Payload:<html>&amp;</html> Method:GET URL:https://cio}"
	if err.Error() != ("status >= 400 & < 9000; " + expectedSuffix) {
		t.Error("Expected error string of: ", "status >= 400 & < 9000", "; Got: ", err.Error())
	}

	if errors.Cause(err).Error() != "status >= 400 & < 9000" {
		t.Error("Expected error cause string of: ", "status >= 400 & < 9000", "; Got: ", err.Error())
	}

	expectedPrefix := "status >= 400 & < 9000\n"
	if plusV := fmt.Sprintf("%+v", err); !strings.HasPrefix(plusV, expectedPrefix) || !strings.HasSuffix(plusV, expectedSuffix) {
		t.Error("Expected +v formatting of: ", expectedPrefix, expectedSuffix, "; Got: ", plusV)
	}

	if code := ErrorStatusCode(err); code != 401 {
		t.Error("Expected error status code of: ", 401, "; Got: ", code)
	}

	if val := ErrorPayload(err); val != "<html>&amp;</html>" {
		t.Error("Expected error payload of: ", "<html>&amp;</html>", "; Got: ", val)
	}

	if val := ErrorMethod(err); val != "GET" {
		t.Error("Expected error method of: ", "GET", "; Got: ", val)
	}

	if val := ErrorURL(err); val != "https://cio" {
		t.Error("Expected error url of: ", "https://cio", "; Got: ", val)
	}

	_, ok := err.(RequestError)
	if !ok {
		t.Error("Expected error to be of type: ", "RequestError", "; Got: ", err)
	}

	{
		jsonBytes, jsonErr := json.Marshal(err)
		if jsonErr != nil {
			t.Error(jsonErr)
		}

		expectedEscapedJSON := `{"Err":"status \u003e= 400 \u0026 \u003c 9000","StatusCode":401,"Payload":"\u003chtml\u003e\u0026amp;\u003c/html\u003e","Method":"GET","URL":"https://cio"}`
		if string(jsonBytes) != expectedEscapedJSON {
			t.Error("Expected json to be: ", expectedEscapedJSON, "; Got: ", string(jsonBytes))
		}

		if unmarshalled, libErr := UnmarshalJSON(jsonBytes); libErr != nil {
			t.Error(libErr)

		} else if !reflect.DeepEqual(unmarshalled.(RequestError).ErrorMetaData, err.(RequestError).ErrorMetaData) {
			t.Error("Expected unmarshalled json error meta data to be: ", err.(RequestError).ErrorMetaData, "; Got: ", unmarshalled.(RequestError).ErrorMetaData)

		} else if unmarshalled.(RequestError).Err.Error() != err.(RequestError).Err.Error() {
			t.Error("Expected unmarshalled json error string to be: ", err.(RequestError).Err, "; Got: ", unmarshalled.(RequestError).Err)
		}
	}

	{
		// Marshel without escaping
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		jsonErr := encoder.Encode(err)
		if jsonErr != nil {
			t.Error(jsonErr)
		}
		jsonBytes := bytes.Trim(buffer.Bytes(), " \r\n")

		expectedUnescapedJSON := `{"Err":"status >= 400 & < 9000","StatusCode":401,"Payload":"<html>&amp;</html>","Method":"GET","URL":"https://cio"}`
		if string(jsonBytes) != expectedUnescapedJSON {
			t.Error("Expected unescaped json to be: ", expectedUnescapedJSON, "; Got: ", string(jsonBytes))
		}

		if unmarshalled, libErr := UnmarshalJSON(jsonBytes); libErr != nil {
			t.Error(libErr)

		} else if !reflect.DeepEqual(unmarshalled.(RequestError).ErrorMetaData, err.(RequestError).ErrorMetaData) {
			t.Error("Expected unmarshalled unescaped json error meta data to be: ", err.(RequestError).ErrorMetaData, "; Got: ", unmarshalled.(RequestError).ErrorMetaData)

		} else if unmarshalled.(RequestError).Err.Error() != err.(RequestError).Err.Error() {
			t.Error("Expected unmarshalled unescaped json error string to be: ", err.(RequestError).Err, "; Got: ", unmarshalled.(RequestError).Err)
		}
	}
}
