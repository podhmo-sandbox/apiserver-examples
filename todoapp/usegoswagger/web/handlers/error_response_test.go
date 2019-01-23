package handlers_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/podhmo-sandbox/apiserver-examples/todoapp/usegoswagger/web/testutil"
	webtest "github.com/podhmo/go-webtest"
)

type fakeT struct {
	*testing.T
	Output []string
}

func (t *fakeT) Fatalf(format string, args ...interface{}) {
	t.Output = append(t.Output, fmt.Sprintf(format, args...))
}

// TestErrorResponse :
func TestErrorResponse(t *testing.T) {
	server, teardown := testutil.CreateServer(
		testutil.CreateAPI(),
	)
	defer teardown()

	t.Run("404", func(t *testing.T) {
		ft := &fakeT{T: t}
		webtest.TryJSONRequest(
			ft,
			server.GetHandler(),
			"GET",
			"/:missing:",
			http.StatusOK,
		)
		if len(ft.Output) != 1 {
			t.Fatalf("must error, but number of error is %d", len(ft.Output))
		}

		expected := `{"code":404,"message":"path /:missing: was not found"}`
		if !strings.Contains(ft.Output[0], expected) {
			t.Errorf("expecting message is not found.\nexpect:\n%s\nactual:\n%s", expected, ft.Output[0])
		}
	})
}
