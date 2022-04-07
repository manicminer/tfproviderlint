package V012_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/passes/V012"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestV012(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, V012.Analyzer, "a")
}
