package S037

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.SchemaAttributeReferencesAnalyzer("S037", schema.SchemaFieldExactlyOneOf)
