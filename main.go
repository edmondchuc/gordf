package main

import (
	"fmt"
	"gordf/parser"
	"gordf/rdf"
)

func main() {
	g := rdf.Graph{}

	parser.ParseFile(&g, "data.nt", "nt")

	rdfType := rdf.NewURI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	schemaName := rdf.NewURI("http://schema.org/name")

	for triple := range g.Triples(rdf.None{}, rdfType, rdf.NewURI("http://schema.org/Organization")) {
		for triple := range g.Triples(triple.S, schemaName, rdf.None{}) {
			fmt.Println(triple)
		}
	}
}