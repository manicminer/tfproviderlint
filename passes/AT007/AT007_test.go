package AT007_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/passes/AT007"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT007(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT007.Analyzer, "a")
}
