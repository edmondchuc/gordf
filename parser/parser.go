package parser

import (
	"fmt"
	"github.com/edmondchuc/gordf/rdf"
)

func ParseFile(graph *rdf.Graph, file, format string) {
	switch format {
	case "nt":
		ParseNT(graph, file)
	default:
		panic(fmt.Sprintf("Invalid format: %v", format))
	}
}