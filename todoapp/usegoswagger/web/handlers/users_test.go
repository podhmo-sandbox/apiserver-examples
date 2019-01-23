package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/components"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/testutil"
)

func TestCreateUser(t *testing.T) {
	now := mytime.MustParse("2000-01-01T00:00:00Z")
	manager, teardown := testutil.CreateServer(testutil.CreateAPI(), components.WithNow(now))
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

	if expected, got := http.StatusCreated, res.StatusCode; got != expected {
		t.Fatalf("status code: expected %d, but %d", expected, got)
	}

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		t.Fatal(err)
	}
}
