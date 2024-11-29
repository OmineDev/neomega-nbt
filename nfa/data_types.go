package nfa

import (
	"fmt"
	"sort"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(elem T) {
	map[T]struct{}(s)[elem] = struct{}{}
}

func (s Set[T]) Remove(elem T) {
	delete(map[T]struct{}(s), elem)
}

func (s Set[T]) Has(elem T) bool {
	_, ok := map[T]struct{}(s)[elem]
	return ok
}

func (s Set[T]) ToList() []T {
	d := map[T]struct{}(s)
	out := make([]T, 0, len(d))
	for k := range d {
		out = append(out, k)
	}
	return out
}

func (s Set[T]) Iter() func(func(T) bool) {
	return func(f func(T) bool) {
		if s == nil {
			return
		}
		for k := range map[T]struct{}(s) {
			if !f(k) {
				return
			}
		}
	}
}

func (s Set[T]) Size() int {
	if s == nil {
		return 0
	}
	return len(map[T]struct{}(s))
}

func (s Set[T]) Pop() (k T, ok bool) {
	for k = range map[T]struct{}(s) {
		ok = true
	}
	if ok {
		delete(s, k)
		return
	}
	return k, false
}

func (s Set[T]) Hash() string {
	if s == nil {
		return ""
	}
	l := s.ToList()
	sl := []string{}
	for _, e := range l {
		sl = append(sl, fmt.Sprintf("%v", e))
	}
	sort.Strings(sl)
	return fmt.Sprintf("%v", sl)
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	if s == nil && s2 == nil {
		return nil
	}
	out := NewSet[T]()
	if s != nil {
		for e := range s.Iter() {
			out.Add(e)
		}
	}
	if s2 != nil {
		for e := range s2.Iter() {
			out.Add(e)
		}
	}
	return out
}

func NewSet[T comparable]() Set[T] {
	return Set[T](map[T]struct{}{})
}

type Stack[T any] struct {
	s []T
	p int
}

func (s *Stack[T]) Push(v T) {
	if s.p == len(s.s) {
		s.p += 1
		s.s = append(s.s, v)
	} else {
		s.s[s.p] = v
		s.p += 1
	}
}

func (s *Stack[T]) Pop() T {
	s.p -= 1
	return s.s[s.p]
}

func (s *Stack[T]) Curr() T {
	return s.s[s.p-1]
}

func (s *Stack[T]) Empty() bool {
	return s.p == 0
}
