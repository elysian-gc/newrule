package main

import (
	"awesomeProject"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(awesomeProject.Analyzer)
}
