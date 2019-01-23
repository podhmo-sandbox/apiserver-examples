package handlers_test

import (
	"testing"

	"net/http"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/mylib/mytime"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/components"
	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/testutil"
)

func TestCreateUser(t *testing.T) {
	server, teardown := testutil.CreateServer(
		testutil.CreateAPI(),
		components.WithNow(mytime.MustParse("2000-01-01T00:00:00Z")),
	)
	defer teardown()

	testutil.TryRequest(
		t,
		server.GetHandler(),
		"POST",
		"/users",
		http.StatusCreated,
		testutil.WithRequestJSONBody(`{"name": "foo"}`),
		testutil.WithRequestAssertJSONResponse(`{"name": "foo", "nilckname": "F"}`),
	)
}
