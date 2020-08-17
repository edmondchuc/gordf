package rdf

import (
	"testing"
)

func TestNewLiteralBool(t *testing.T) {
	got := NewLiteral(true)

	want := true
	if got.value != want {
		t.Errorf(`got.value = %v; want %v, <nil>`, got.value, want)
	}
}

func TestNewLiteralNumber(t *testing.T) {
	got := NewLiteral(1992)

	want := 1992
	if got.value != want {
		t.Errorf(`got.value = %v; want %v, <nil>`, got.value, want)
	}
}

// TODO: Test other New functions for Literal (datatype and language).

func TestLiteral_String(t *testing.T) {
	got := NewLiteral("Edmond")

	want := `"Edmond"`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v"`, got.String(), want)
	}
}

func TestLiteral_StringWithDatatype(t *testing.T) {
	datatype := NewURI("http://www.w3.org/2001/XMLSchema#string")
	got := NewLiteralWithDatatype("Edmond", datatype)

	want := `"Edmond"^^<http://www.w3.org/2001/XMLSchema#string>`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got.String(), want)
	}
}

func TestLiteral_StringWithLanguage(t *testing.T) {
	got := NewLiteralWithLanguage("Edmond", "en")

	want := `"Edmond"@en`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got, want)
	}
}

func TestLiteral_StringWithLanguageAndDatatype(t *testing.T) {
	datatype := NewURI("http://www.w3.org/2001/XMLSchema#string")
	got := NewLiteralWithLanguageAndDatatype("Edmond", datatype, "en")

	want := `"Edmond"@en^^<http://www.w3.org/2001/XMLSchema#string>`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got, want)
	}
}

func TestLiteral_Bool(t *testing.T) {
	got := NewLiteral(true)

	want := `"true"^^<http://www.w3.org/2001/XMLSchema#boolean>`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got, want)
	}
}

func TestLiteral_Number(t *testing.T) {
	got := NewLiteral(1992)

	want := `1992`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got, want)
	}
}

func TestLiteral_Equals(t *testing.T) {
	literal1 := NewLiteral("Edmond")
	literal2 := NewLiteral("Edmond")
	literal3 := NewLiteral(123)
	literal4 := NewLiteral(123)
	literal5 := NewLiteral(false)
	literal6 := NewLiteral(false)

	if !literal1.Equals(literal2) {
		t.Errorf(`Expected values to equal.`)
	}

	if !literal3.Equals(literal4) {
		t.Errorf(`Expected values to equal.`)
	}

	if !literal5.Equals(literal6) {
		t.Errorf(`Expected values to equal.`)
	}
}

func TestLiteral_Equals2(t *testing.T) {
	literal1 := NewLiteral("Edmond")
	literal2 := NewLiteral("Edmond2")
	literal3 := NewLiteral(123)
	literal4 := NewLiteral(999)
	literal5 := NewLiteral(true)
	literal6 := NewLiteral(false)

	if literal1.Equals(literal2) {
		t.Errorf(`Expected values to not equal.`)
	}

	if literal3.Equals(literal4) {
		t.Errorf(`Expected values to not equal.`)
	}

	if literal5.Equals(literal6) {
		t.Errorf(`Expected values to not equal.`)
	}
}

func TestLiteral_IsURI(t *testing.T) {
	literal := NewLiteral(12)
	got := literal.IsURI()
	want := false
	if got != want {
		t.Errorf(`got %v; want %v`, got, want)
	}
}