package renderer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func newJSONEncoder(w io.Writer) *json.Encoder {
	d := json.NewEncoder(w)
	d.SetIndent("", "  ")
	return d
}

// JSON :
func JSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "/application/json")
	if err := newJSONEncoder(w).Encode(res); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Error :
func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintf(w, `{"error": %q}`, error)
}
