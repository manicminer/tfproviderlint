package R004_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/manicminer/tfproviderlint/passes/R004"
)

func TestR004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, R004.Analyzer, "a")
}
