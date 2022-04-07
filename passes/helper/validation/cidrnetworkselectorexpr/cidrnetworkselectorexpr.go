package cidrnetworkselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"cidrnetworkselectorexpr",
	validation.IsFunc,
	validation.PackagePath,
	validation.FuncNameCIDRNetwork,
)
