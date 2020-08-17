package rdf

import "fmt"

// The Literal struct implements the Node interface.
type Literal struct {
	value    interface{}
	datatype URI
	language string
}

func (literal Literal) Equals(n Node) bool {
	literal2, ok := n.(Literal)
	if ok {
		return literal.value == literal2.value
	}
	return false
}

func (literal Literal) Value() string {
	return fmt.Sprintf("%v", literal.value)
}

func (literal Literal) IsURI() bool {
	return false
}

func (literal Literal) String() string {

	// Literal is a number.
	// TODO: Check if checking these types is enough for numbers. I think just checking int is enough
	// 	as it should be the set of all ints, but not 100% sure.
	_, okInt := literal.value.(int)
	_, okFloat32 := literal.value.(float32)
	_, okFloat64 := literal.value.(float64)
	if okInt || okFloat32 || okFloat64 {
		return fmt.Sprintf(`%v`, literal.value)
	}

	// Literal is a bool.
	_, ok := literal.value.(bool)
	if ok {
		return fmt.Sprintf(`"%v"^^<http://www.w3.org/2001/XMLSchema#boolean>`, literal.value)
	}

	// Literal is a string with a language tag and a datatype.
	if literal.language != "" && literal.datatype.value != "" {
		return fmt.Sprintf(`"%v"@%v^^%v`, literal.value, literal.language, literal.datatype)
	}

	// Literal is a string and has either a language tag or datatype.
	if literal.datatype.value != "" {
		return fmt.Sprintf(`"%v"^^%v`, literal.value, literal.datatype)
	} else if literal.language != "" {
		return fmt.Sprintf(`"%v"@%v`, literal.value, literal.language)
	}

	// Literal is a string and has no language tag or datatype.
	return fmt.Sprintf(`"%v"`, literal.value)
}

func NewLiteralWithLanguageAndDatatype(value string, datatype URI, language string) Literal {
	return Literal{value: value, datatype: datatype, language: language}
}

func NewLiteralWithDatatype(value string, datatype URI) Literal {
	return Literal{value: value, datatype: datatype}
}

func NewLiteralWithLanguage(value string, language string) Literal {
	return Literal{value: value, language: language}
}

func NewLiteral(value interface{}) Literal {
	return Literal{value: value}
}