package S036

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
)

var Analyzer = analysisutils.SchemaAttributeReferencesAnalyzer("S036", schema.SchemaFieldConflictsWith)
