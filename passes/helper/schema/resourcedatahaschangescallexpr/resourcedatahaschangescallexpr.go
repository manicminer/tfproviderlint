package resourcedatahaschangescallexpr

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.ReceiverMethodCallExprAnalyzer(
	"resourcedatahaschangescallexpr",
	schema.IsReceiverMethod,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"HasChanges",
)
