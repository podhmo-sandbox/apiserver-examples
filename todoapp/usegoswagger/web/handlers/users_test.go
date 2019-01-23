package handlers_test

import (
	"testing"

	"net/http"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/components"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/testutil"
	webtest "github.com/podhmo/go-webtest"
)

func TestCreateUser(t *testing.T) {
	server, teardown := testutil.CreateServer(
		testutil.CreateAPI(),
		components.WithNow(mytime.MustParse("2000-01-01T00:00:00Z")),
	)
	defer teardown()

	webtest.TryJSONRequest(
		t,
		server.GetHandler(),
		"POST",
		"/users",
		http.StatusCreated,
		webtest.WithJSONBody(`{"name": "foo"}`),
		webtest.WithAssertJSONResponse(`{"id": "1", "name": "foo", "createdAt": "2000-01-01T00:00:00Z"}`),
	)
}
