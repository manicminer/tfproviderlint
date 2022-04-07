package providertypeassertexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.TypeAssertExprAnalyzer(
	"providertypeassertexpr",
	schema.IsFunc,
	schema.PackagePath,
	schema.TypeNameProvider,
)
