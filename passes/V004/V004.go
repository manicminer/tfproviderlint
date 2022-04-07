package V004

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/singleipcallexpr"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/singleipselectorexpr"
)

var Analyzer = analysisutils.DeprecatedEmptyCallExprWithReplacementSelectorExprAnalyzer(
	"V004",
	singleipcallexpr.Analyzer,
	singleipselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameSingleIP,
	validation.PackagePath,
	validation.FuncNameIsIPAddress,
)
