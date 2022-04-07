package XR008

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/passes/stdlib/osexeccommandcontextcallexpr"
	"github.com/manicminer/tfproviderlint/passes/stdlib/osexeccommandcontextselectorexpr"
)

var Analyzer = analysisutils.AvoidSelectorExprAnalyzer(
	"XR008",
	osexeccommandcontextcallexpr.Analyzer,
	osexeccommandcontextselectorexpr.Analyzer,
	"os/exec",
	"CommandContext",
)
