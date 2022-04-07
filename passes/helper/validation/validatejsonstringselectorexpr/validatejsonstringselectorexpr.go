package validatejsonstringselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"validatejsonstringselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameValidateJsonString,
)
