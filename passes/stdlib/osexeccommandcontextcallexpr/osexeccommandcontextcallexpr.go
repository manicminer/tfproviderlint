package osexeccommandcontextcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"osexeccommandcontextcallexpr",
	"os/exec",
	"CommandContext",
)
