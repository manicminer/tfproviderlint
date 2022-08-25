package S016_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/manicminer/tfproviderlint/passes/S016"
)

func TestS016(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, S016.Analyzer, "a")
}
