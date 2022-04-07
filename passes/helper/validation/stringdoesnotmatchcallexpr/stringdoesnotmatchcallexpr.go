package stringdoesnotmatchcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"stringdoesnotmatchcallexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameStringDoesNotMatch,
)
