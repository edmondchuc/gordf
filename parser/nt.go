package parser

import (
	"fmt"
	"github.com/iand/ntriples"
	"gordf/rdf"
	"os"
)

func getGoRDFNode(node ntriples.RdfTerm) rdf.Node {
	if node.IsIRI() {
		return rdf.NewURI(node.Value)
	} else if node.IsBlank() {
		panic("Blank nodes for memory graph not supported yet.")
	} else if node.IsLiteral() {
		return rdf.NewLiteral(node.Value)
	} else if node.IsLanguageLiteral() {
		return rdf.NewLiteralWithLanguage(node.Value, node.Language)
	} else if node.IsTypedLiteral() {
		return rdf.NewLiteralWithDatatype(node.Value, rdf.NewURI(node.DataType))
	} else {
		panic("Error.")
	}
}

func ParseNT(graph *rdf.Graph, file string) {
	ntfile, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}
	defer ntfile.Close()

	r := ntriples.NewReader(ntfile)

	for triple, err := r.Read(); err == nil;  triple, err = r.Read() {
		subj := getGoRDFNode(triple.S)
		pred := getGoRDFNode(triple.P)
		obj := getGoRDFNode(triple.O)
		graph.Add(subj, pred, obj)
	}
}