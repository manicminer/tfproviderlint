package S025_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/passes/S025"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestS025(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S025.Analyzer, "a")
}
