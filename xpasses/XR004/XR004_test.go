package XR004_test

import (
	"testing"

	"github.com/manicminer/tfproviderlint/xpasses/XR004"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestXR004(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, XR004.Analyzer, "a")
}
