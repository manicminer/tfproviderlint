package stringmatchcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringmatchcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringMatch,
)
