package parser

import (
	"fmt"
	"gordf/gordf"
)

func ParseFile(graph *gordf.Graph, file, format string) {
	switch format {
	case "nt":
		ParseNT(graph, file)
	default:
		panic(fmt.Sprintf("Invalid format: %v", format))
	}
}