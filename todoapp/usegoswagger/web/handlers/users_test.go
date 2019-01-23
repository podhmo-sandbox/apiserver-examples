package handlers_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"net/http/httptest"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/testutil"
)

func TestCreateUser(t *testing.T) {
	manager, teardown := testutil.CreateServer()
	defer teardown()

	mux := manager.GetHandler()
	rec := httptest.NewRecorder()

	body := map[string]interface{}{
		"name": "foo",
	}

	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	if err := encoder.Encode(&body); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("POST", "/users", &b)
	req.Header.Set("Content-Type", "application/json")
	mux.ServeHTTP(rec, req)
	res := rec.Result()

	if expected, got := 200, res.StatusCode; got != expected {
		t.Fatalf("status code: expected %d, but %d", expected, got)
	}
}
