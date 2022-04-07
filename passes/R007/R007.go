package R007

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/resourcedatapartialcallexpr"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/resourcedatapartialselectorexpr"
)

var Analyzer = analysisutils.DeprecatedReceiverMethodSelectorExprAnalyzer(
	"R007",
	resourcedatapartialcallexpr.Analyzer,
	resourcedatapartialselectorexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"Partial",
)
