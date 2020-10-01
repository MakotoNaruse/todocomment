package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"practice/todocomment"
)

func main() { unitchecker.Main(todocomment.Analyzer) }
