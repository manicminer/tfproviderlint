package fmtsprintfcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"fmtsprintfcallexpr",
	"fmt",
	"Sprintf",
)
