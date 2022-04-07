package V006

import (
	"github.com/manicminer/tfproviderlint/helper/analysisutils"
	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/validation"
	"github.com/manicminer/tfproviderlint/passes/helper/validation/validatelistuniquestringsselectorexpr"
)

var Analyzer = analysisutils.DeprecatedWithReplacementSelectorExprAnalyzer(
	"V006",
	validatelistuniquestringsselectorexpr.Analyzer,
	validation.PackagePath,
	validation.FuncNameValidateListUniqueStrings,
	validation.PackagePath,
	validation.FuncNameListOfUniqueStrings,
)
