package teststepinfo

import (
	"go/ast"
	"reflect"

	"github.com/manicminer/tfproviderlint/helper/terraformtype/helper/resource"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "teststepinfo",
	Doc:  "find github.com/hashicorp/terraform-plugin-sdk/helper/resource.TestStep literals for later passes",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:        run,
	ResultType: reflect.TypeOf([]*resource.TestStepInfo{}),
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	var result []*resource.TestStepInfo

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		x := n.(*ast.CompositeLit)

		if !isResourceTestStep(pass, x) {
			return
		}

		result = append(result, resource.NewTestStepInfo(x, pass.TypesInfo))
	})

	return result, nil
}

func isResourceTestStep(pass *analysis.Pass, cl *ast.CompositeLit) bool {
	switch v := cl.Type.(type) {
	default:
		return false
	case *ast.SelectorExpr:
		return resource.IsTypeTestStep(pass.TypesInfo.TypeOf(v))
	}
}
