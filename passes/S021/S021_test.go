package S021_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/passes/S021"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS021(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S021.Analyzer, "a")
}
