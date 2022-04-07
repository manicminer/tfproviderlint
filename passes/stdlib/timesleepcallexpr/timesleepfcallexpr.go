package timesleepcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.StdlibFunctionCallExprAnalyzer(
	"timesleepcallexpr",
	"time",
	"Sleep",
)
