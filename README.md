# gordf

The gordf module is an RDF library for Go.

## Current features
- Parse N-Triples from file
- Naive memory graph data structure
- Idiomatic querying of the memory graph
- Structs to represent URI and literals (both typed and lang tags)

## Planned features
- Blank Nodes
- Parsing other RDF formats such as Turtle
- Serialise to disk in RDF formats (N-Triple, Turtle, etc)
- Idiomatic representation of CURIES E.g. `rdf:type` -> `RDF{l: "type"}`

## Example usage
```go
package main

import (
	"fmt"
	"github.com/edmondchuc/gordf/parser"
	"github.com/edmondchuc/gordf/rdf"
)

func main() {
	// Create new graph.
	g := rdf.Graph{}

	// Parse triples from file.
	parser.ParseFile(&g, "data.nt", "nt")

	rdfType := rdf.NewURI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	schemaName := rdf.NewURI("http://schema.org/name")

	// Query the graph by looping.
	for triple := range g.Triples(rdf.None{}, rdfType, rdf.NewURI("http://schema.org/Organization")) {
		for triple := range g.Triples(triple.S, schemaName, rdf.None{}) {
			fmt.Println(triple)
		}
	}

	// Add new triples.
	g.Add(rdf.NewURI("https://edmondchuc.com/me"), schemaName, rdf.NewLiteral("Edmond Chuc"))
}
```