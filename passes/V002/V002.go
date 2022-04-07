package V002

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/cidrnetworkselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V002",
	cidrnetworkselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
	validation.PackagePath,
	validation.FuncNameIsCIDRNetwork,
)
