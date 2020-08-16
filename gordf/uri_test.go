package gordf

import "testing"

func TestNewURI(t *testing.T) {
	got := NewURI("http://schema.org/Person")

	want := "http://schema.org/Person"
	if got.value != want {
		t.Errorf(`got.value = %v; want %v`, got.value, want)
	}
}

func TestURI_String(t *testing.T) {
	got := NewURI("http://schema.org/Person")

	want := `<http://schema.org/Person>`
	if got.String() != want {
		t.Errorf(`got.String() = %v; want %v`, got, want)
	}
}

func TestURI_Equals(t *testing.T) {
	uri1 := NewURI("http://schema.org/Person")
	uri2 := NewURI("http://schema.org/Person")

	if !uri1.Equals(uri2) {
		t.Errorf(`Expected uri1 and uri2 to be the same value.`)
	}
}

func TestURI_Equals2(t *testing.T) {
	uri1 := NewURI("http://schema.org/Person")
	uri2 := NewURI("http://schema.org/Organization")

	if uri1.Equals(uri2) {
		t.Errorf(`Expected uri1 and uri2 to not be the same value.`)
	}
}