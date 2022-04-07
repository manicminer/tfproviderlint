package osexeccommandcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"osexeccommandcallexpr",
	"os/exec",
	"Command",
)
