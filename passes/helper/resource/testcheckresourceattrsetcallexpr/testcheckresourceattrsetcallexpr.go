package testcheckresourceattrsetcallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/resource"
)

var Analyzer = analysisutils.FunctionCallExprAnalyzer(
	"testcheckresourceattrsetcallexpr",
	resource.IsFunc,
	resource.PackagePath,
	resource.FuncNameTestCheckResourceAttrSet,
)
