package util

import "testing"

func TestIsValidURI(t *testing.T) {
	t.Run("IsValidURI returns false", func(t *testing.T) {
		got := IsValidURI("123")
		want := false
		if got != want {
			t.Errorf(`got %v; want %v`, got, want)
		}
	})

	t.Run("IsValidURI returns true", func(t *testing.T) {
		got := IsValidURI("https://edmondchuc.com")
		want := true
		if got != want {
			t.Errorf(`got %v; want %v`, got, want)
		}
	})
}
