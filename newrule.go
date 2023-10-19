package main

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis/singlechecker"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "newrule",
	Doc:  "check for 'nohup' string",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if str, ok := n.(*ast.BasicLit); ok && str.Kind == token.STRING {
				if strings.Contains(str.Value, "nohup") {
					pass.Reportf(str.Pos(), "Found 'nohup' string")
				}
			}
			return true
		})
	}
	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
