package iprangecallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"iprangecallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
