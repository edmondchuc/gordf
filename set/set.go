package set

import "fmt"

type void struct{}
var member void

// Generic set implementation inspired by https://yourbasic.org/golang/implement-set/
// TODO: Implement Union(), Intersection(), Difference()
type Set struct {
	m map[interface{}]void
}

func (s *Set) Add(value interface{}) {
	if s.m == nil {
		s.m = make(map[interface{}]void)
	}
	s.m[value] = member
}

func (s *Set) Remove(value interface{}) {
	if _, ok :=s.m[value]; ok {
		delete(s.m, value)
	}
}

func (s Set) String() string {
	var value = ""
	var i = 0
	length := len(s.m)
	for key, _ := range s.m {
		_, ok := key.(string)
		var str string
		if ok {
			str = fmt.Sprintf(`"%v"`, key)
		} else {
			str = fmt.Sprintf(`%v`, key)
		}

		value += str
		i++

		if i != length {
			value += ", "
		}
	}
	return fmt.Sprintf(`(%v)`, value)
}

func (s Set) Len() int {
	return len(s.m)
}

func (s Set) Exists(v interface{}) bool {
	if _, ok := s.m[v]; ok {
		return true
	}
	return false
}

func (s Set) Range() chan interface{} {
	values := make(chan interface{})

	go func() {
		for value, _ := range s.m {
			values <- value
		}

		close(values)
	}()

	return values
}