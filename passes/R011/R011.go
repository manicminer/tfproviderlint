package R011

import (
	"go/ast"

	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/manicminer/tfproviderlint/passes/commentignore"
	"github.com/manicminer/tfproviderlint/passes/helper/schema/resourceinfo"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Resource with MigrateState configured

The R011 analyzer reports cases of resources which configure MigrateState.
After Terraform 0.12, resources must configure new state migrations via
StateUpgraders. Existing implementations of MigrateState prior to Terraform
0.12 can be ignored currently.`

const analyzerName = "R011"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		resourceinfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resourceInfos := pass.ResultOf[resourceinfo.Analyzer].([]*schema.ResourceInfo)
	for _, resourceInfo := range resourceInfos {
		if ignorer.ShouldIgnore(analyzerName, resourceInfo.AstCompositeLit) {
			continue
		}

		if resourceInfo.Resource.MigrateState == nil {
			continue
		}

		switch t := resourceInfo.AstCompositeLit.Type.(type) {
		default:
			pass.Reportf(resourceInfo.AstCompositeLit.Lbrace, "%s: resource should configure StateUpgraders instead of MigrateState (implementations prior to Terraform 0.12 can be ignored)", analyzerName)
		case *ast.SelectorExpr:
			pass.Reportf(t.Sel.Pos(), "%s: resource should configure StateUpgraders instead of MigrateState (implementations prior to Terraform 0.12 can be ignored)", analyzerName)
		}
	}

	return nil, nil
}
