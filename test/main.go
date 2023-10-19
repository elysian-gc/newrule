package main

import (
	"awesomeProject/newrule"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(newrule.Analyzer)
}
