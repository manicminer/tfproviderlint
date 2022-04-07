package XR007

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/passes/stdlib/osexeccommandcallexpr"
	"github.com/manicminer/tfproviderlint/passes/stdlib/osexeccommandselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR007",
	osexeccommandcallexpr.Analyzer,
	osexeccommandselectorexpr.Analyzer,
	"os/exec",
	"Command",
)
