package testutil

import (
	"bytes"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// StringDiff :
func StringDiff(before, after string) string {
	dmp := diffmatchpatch.New()
	res1, res2, res3 := dmp.DiffLinesToChars(before, after)
	diffs := dmp.DiffCharsToLines(dmp.DiffMain(res1, res2, false), res3)

	var result bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			_, _ = result.WriteString("\n+ ")
			s := strings.Split(diff.Text, "\n")
			_, _ = result.WriteString(strings.Join(s[:len(s)-1], "\n+ "))
			_, _ = result.WriteString("\n")
		case diffmatchpatch.DiffDelete:
			_, _ = result.WriteString("\n- ")
			s := strings.Split(diff.Text, "\n")
			_, _ = result.WriteString(strings.Join(s[:len(s)-1], "\n- "))
			_, _ = result.WriteString("\n")
		}
	}
	return result.String()
}
