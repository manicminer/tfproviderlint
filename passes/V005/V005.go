package V005

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/validatejsonstringselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V005",
	validatejsonstringselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
	validation.PackagePath,
	validation.FuncNameStringIsJSON,
)
