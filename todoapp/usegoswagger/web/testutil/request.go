package testutil

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
)

// RequestOption :
type RequestOption struct {
	Method     string
	Path       string
	Body       io.Reader
	Assertions []func(t *testing.T, res *TryRequestResponse)
	Response   TryRequestResponse
}

// TryRequestResponse :
type TryRequestResponse struct {
	Body bytes.Buffer
	*http.Response
}

// TryRequest :
func TryRequest(t *testing.T, mux http.Handler, method, path string, status int, options ...func(*RequestOption) error) *TryRequestResponse {
	t.Helper()
	rop := &RequestOption{
		Method: method,
		Path:   path,
	}

	for _, op := range options {
		if err := op(rop); err != nil {
			t.Fatalf("apply option %+v", err)
			return nil
		}
	}
	req := httptest.NewRequest(rop.Method, rop.Path, rop.Body)

	// todo: to option
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	res := rec.Result()
	tresponse := TryRequestResponse{
		Response: res,
	}

	{
		if _, err := io.Copy(&tresponse.Body, res.Body); err != nil {
			t.Fatalf("parse responose, something wrong: %+v", err)
			return nil
		}
		res.Body = ioutil.NopCloser(&tresponse.Body)
	}

	if expected, got := status, res.StatusCode; got != expected {
		t.Fatalf("status code: expected %d, but %d\n%s", expected, got, tresponse.Body.String())
		return nil
	}

	for _, assert := range rop.Assertions {
		assert(t, &tresponse)
	}
	return &tresponse
}

// WithRequestJSONBody :
func WithRequestJSONBody(body string) func(rop *RequestOption) error {
	return func(rop *RequestOption) error {
		rop.Body = bytes.NewBufferString(body)
		return nil
	}
}

// WithRequestAssert :
func WithRequestAssert(assert func(t *testing.T, res *TryRequestResponse)) func(rop *RequestOption) error {
	return func(rop *RequestOption) error {
		rop.Assertions = append(rop.Assertions, assert)
		return nil
	}
}

// WithRequestAssertJSONResponse :
func WithRequestAssertJSONResponse(body string) func(rop *RequestOption) error {
	return func(rop *RequestOption) error {
		var expected string
		{
			var ob interface{}
			if err := json.Unmarshal([]byte(body), &ob); err != nil {
				return errors.Wrap(err, "prepare unmarsal")
			}
			b, err := json.MarshalIndent(&ob, "", "  ")
			if err != nil {
				return errors.Wrap(err, "prepare marsal")
			}
			expected = string(b)
		}

		rop.Assertions = append(rop.Assertions, func(t *testing.T, res *TryRequestResponse) {
			var actual string
			var ob interface{}

			{
				decoder := json.NewDecoder(&res.Body)
				if err := decoder.Decode(&ob); err != nil {
					t.Fatalf("unexpected response: %q", res.Body.String())
				}
			}

			{
				b, err := json.MarshalIndent(&ob, "", "  ")
				if err != nil {
					panic(err) // something wrong
				}
				actual = string(b)
			}

			if expected != actual {
				t.Fatalf("mismatch response:\n***diff (- missing, + excess)***\n%s\n\n***expected***\n%s\n\n***actual***\n%s", StringDiff(expected, actual), expected, actual)
			}
		})
		return nil
	}
}
