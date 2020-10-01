package main

import (
	"github.com/MakotoNaruse/todocomment"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(todocomment.Analyzer) }
