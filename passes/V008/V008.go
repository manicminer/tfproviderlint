package V008

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/validaterfc3339timestringselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V008",
	validaterfc3339timestringselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateRFC3339TimeString,
	validation.PackagePath,
	validation.FuncNameIsRFC3339Time,
)
