package todocomment_test

import (
	"golang.org/x/tools/go/analysis/passes/findcall"
	"testing"

	"github.com/MakotoNaruse/todocomment"
	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	findcall.Analyzer.Flags.Set("issue", "https://github.com/test/test/issues/")
}

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, todocomment.Analyzer, "a")
}
