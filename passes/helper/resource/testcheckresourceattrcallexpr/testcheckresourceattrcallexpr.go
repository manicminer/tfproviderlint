package testcheckresourceattrcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattrcallexpr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttr,
)
