package app

import (
	"encoding/json"
	"io"
)

// NewJSONEncoder :
func NewJSONEncoder(w io.Writer) *json.Encoder {
	d := json.NewEncoder(w)
	d.SetIndent("", "  ")
	return d
}
