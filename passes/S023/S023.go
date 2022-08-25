// Package S023 defines an Analyzer that checks for
// Schema that should omit Elem with incompatible Type
package S023

import (
	"golang.org/x/tools/go/analysis"

	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/manicminer/tfproviderlint/passes/commentignore"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/schemainfo"
)

const Doc = `check for Schema that should omit Elem with incompatible Type

The S023 analyzer reports cases of schema that declare Elem that should
be removed with incompatible Type.`

const analyzerName = "S023"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		schemainfo.Analyzer,
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)
	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		if schemaInfo.IsOneOfTypes(schema.SchemaValueTypeList, schema.SchemaValueTypeMap, schema.SchemaValueTypeSet) {
			continue
		}

		if !schemaInfo.DeclaresField(schema.SchemaFieldElem) {
			continue
		}

		pass.Reportf(schemaInfo.Fields[schema.SchemaFieldElem].Value.Pos(), "%s: schema should not include Elem with incompatible Type", analyzerName)
	}

	return nil, nil
}
