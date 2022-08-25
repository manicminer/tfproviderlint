package AT005_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/passes/AT005"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAT005(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, AT005.Analyzer, "a")
}
