package rdf

import "fmt"

// The URI struct implements the Node interface.
type URI struct {
	value string
}

func (uri URI) Equals(n Node) bool {
	uri2, ok := n.(URI)
	if ok {
		return uri.value == uri2.value
	}
	return false
}

// String returns the URI string.
func (uri URI) String() string {
	return fmt.Sprintf(`<%v>`, uri.value)
}

func (uri URI) Value() string {
	return uri.value
}

// NewURI returns a new URI.
func NewURI(uri string) URI {
	return URI{value: uri}
}