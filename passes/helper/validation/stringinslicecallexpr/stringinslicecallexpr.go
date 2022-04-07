package stringinslicecallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringinslicecallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringInSlice,
)
