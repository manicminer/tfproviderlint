package fmterrorfcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmterrorfcallexpr",
	"fmt",
	"Errorf",
)
