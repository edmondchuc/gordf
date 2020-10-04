package main

import (
	"fmt"
	"github.com/edmondchuc/gordf/parser"
	"github.com/edmondchuc/gordf/rdf"
	"strings"
	"time"
)

func main() {
	starttime := time.Now()

	// Create new graph.
	g := rdf.Graph{}

	// Parse triples from file.
	parser.ParseFile(&g, "data.nt", "nt")

	rdfType := rdf.NewURI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	schemaName := rdf.NewURI("http://schema.org/name")
	schemaOrg := rdf.NewURI("http://schema.org/Organization")

	// Query the graph by looping.
	for triple := range g.Triples(rdf.None{}, rdfType, schemaOrg) {
		for triple := range g.Triples(triple.S, schemaName, rdf.None{}) {
			fmt.Println(triple)
		}
	}

	// Add new triples.
	g.Add(rdf.NewURI("https://edmondchuc.com/me"), schemaName, rdf.NewLiteral("Edmond Chuc"))

	// Print anything with schema:name value containing "Ed".
	for triple := range g.Triples(rdf.None{}, schemaName, rdf.None{}) {
		if strings.Contains(triple.O.Value(), "Ed") {
			fmt.Println(triple.O)
		}
	}

	fmt.Println("Execution time: ", time.Now().Sub(starttime))
}