package iprangeselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"iprangeselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameIPRange,
)
