package R008

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/resourcedatasetpartialcallexpr"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/resourcedatasetpartialselectorexpr"
)

var Analyzer = analysisutils.DeprecatedReceiverMethodSelectorExprAnalyzer(
	"R008",
	resourcedatasetpartialcallexpr.Analyzer,
	resourcedatasetpartialselectorexpr.Analyzer,
	schema.PackagePath,
	schema.TypeNameResourceData,
	"SetPartial",
)
