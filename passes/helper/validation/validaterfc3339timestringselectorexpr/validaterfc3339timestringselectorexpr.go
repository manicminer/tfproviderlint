package validaterfc3339timestringselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validaterfc3339timestringselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateRFC3339TimeString,
)
