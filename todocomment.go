package todocomment

import (
	"flag"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "todocomment is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "todocomment",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var issue string

func init() {
	flag.StringVar(&issue, "issue", "https://github.com/test/test/issues/", "report TODO comment doesn't contains issue's string ex. https://github.com/test/test/issues/")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			for _, comment := range n.Comments {
				s := comment.Text()
				if strings.Contains(s, "TODO") {
					if !strings.Contains(s, issue) {
						pass.Reportf(comment.Pos(), "todo comment must contains issue's link")
					}
				}
			}
		}
	})

	return nil, nil
}
