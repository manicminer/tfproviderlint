package osexeccommandselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionSelectorExprAnalyzer(
	"osexeccommandselectorexpr",
	"os/exec",
	"Command",
)
