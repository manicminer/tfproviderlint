package V007

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/validateregexpselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V007",
	validateregexpselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateRegexp,
	validation.PackagePath,
	validation.FuncNameStringIsValidRegExp,
)
