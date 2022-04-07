package resourceproviderselectorexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/terraform"
)

var Analyzer = analysisutils.SelectorExprAnalyzer(
	"resourceproviderselectorexpr",
	terraform.IsFunc,
	terraform.PackagePath,
	terraform.TypeNameResourceProvider,
)
