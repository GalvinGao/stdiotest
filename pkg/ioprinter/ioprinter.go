package ioprinter

import (
	"fmt"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func Diff(expected, actual string) string {
	edits := myers.ComputeEdits(span.URIFromPath("expected"), expected, actual)
	return fmt.Sprint(gotextdiff.ToUnified("expected", "actual", expected, edits))
}
