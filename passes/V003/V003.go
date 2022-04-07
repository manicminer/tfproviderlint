package V003

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/iprangecallexpr"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/iprangeselectorexpr"
)

var Analyzer = analysisutils.DeprecatedEmptyCallExprWithReplacementSelectorExprAnalyzer(
	"V003",
	iprangecallexpr.Analyzer,
	iprangeselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameIPRange,
	validation.PackagePath,
	validation.FuncNameIsIPv4Range,
)
