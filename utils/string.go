package utils

import (
	"strings"
)

// A wrapper for doing string operations with chaining
// Because strings.Replace(strings.Replace(strings.Replace(...))) is awful compared to
// utils.NewString().
//   Replace(...).
//   Replace(...).
//   Replace(...)

type String struct {
	str string
}

func NewString(towrap string) *String {
	return &String{str:towrap}
}

func (s *String) Add(str string) *String {
	s.str += str
	return s
}

func (s *String) Replace(one, two string) *String {
	s.str = strings.Replace(s.str, one, two, -1)
	return s
}

func (s *String) Build() string {
	return s.str
}
