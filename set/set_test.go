package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := Set{}
	s.Add("yellow")

	assertCorrectMessage := func(t *testing.T, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf(`got %v; want %v`, got, want)
		}
	}

	t.Run("Set.String() for string type and Set.Add()", func(t *testing.T){
		got := s.String()
		want := `("yellow")`
		assertCorrectMessage(t, got, want)
	})

	t.Run("Set.Exists() true", func(t *testing.T){
		got := s.Exists("yellow")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("Set.Len() and Set.Remove()", func(t *testing.T){
		got := s.Len()
		want := 1
		assertCorrectMessage(t, got, want)

		s.Remove("yellow")
		got = s.Len()
		want = 0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Set.Exists() false", func(t *testing.T) {
		got := s.Exists("yellow")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("Set.Range() returns values", func(t *testing.T) {
		s.Add("blue")
		s.Add(123)
		values := s.Range()
		for value := range values {
			fmt.Println(value)
			if value != "blue" && value != 123 {
				t.Errorf("Expected value to be either 'blue' or 123")
			}
		}
	})

	t.Run("Set.String() for a non-string type", func(t *testing.T) {
		s.Remove("blue")
		got := s.String()
		want := "(123)"
		if got != want {
			t.Errorf(`got %v; want %v`, got, want)
		}
	})

	t.Run("Set.String() for non-string type with comma", func(t *testing.T) {
		s.Add(true)
		got := s.String()
		want1 := "(123, true)"
		want2 := "(true, 123)"

		if got != want1 && got != want2 {
			t.Errorf(`got %v; want1 %v or want2 %v`, got, want1, want2)
		}
	})
}