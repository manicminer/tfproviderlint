package resourceproviderfactoryselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderfactoryselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProviderFactory,
)
