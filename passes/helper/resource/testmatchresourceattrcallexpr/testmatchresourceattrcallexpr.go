package testmatchresourceattrcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testmatchresourceattrcallexpr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestMatchResourceAttr,
)
