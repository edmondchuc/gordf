package rdf

// An RDF resource interface.
type Node interface {
	String() string
	Equals(n Node) bool
	Value() string
	IsURI() bool
}

type None struct {}

func (none None) String() string {
	return "None"
}

func (none None) Equals(n Node) bool {
	return false
}

func (none None) Value() string {
	return ""
}

func (none None) IsURI() bool {
	return false
}