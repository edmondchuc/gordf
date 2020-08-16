package set

import "testing"

func TestSet(t *testing.T) {
	s := Set{}
	s.Add("yellow")

	assertCorrectMessage := func(t *testing.T, got, want interface{}) {
		t.Helper()
		if got != want {
			t.Errorf(`got %v; want %v`, got, want)
		}
	}

	t.Run("Set.String() and Set.Add()", func(t *testing.T){
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
}