package singleipselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"singleipselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameSingleIP,
)
