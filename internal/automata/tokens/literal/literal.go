package literal

import (
	"unicode"

	. "github.com/marcobarao/parser/internal/automata"
)

func NewLiteralAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, true)
	aut.AddState(2, false)
	aut.AddState(3, true)
	aut.AddState(4, false)
	aut.AddState(5, false)
	aut.AddState(6, true)
  aut.AddState(7, false)
  aut.AddState(8, false)
  aut.AddState(9, true)
  aut.AddState(10, false)
  aut.AddState(11, false)
  aut.AddState(12, false)
  aut.AddState(13, true)
  aut.AddState(14, false)
  aut.AddState(15, false)
  aut.AddState(16, false)
  aut.AddState(17, false)
  aut.AddState(18, true)

	aut.AddTransition(0, func(input rune) bool { return unicode.IsNumber(input) }, 1)
	aut.AddTransition(0, func(input rune) bool { return input == '.' }, 2)
	aut.AddTransition(1, func(input rune) bool { return unicode.IsNumber(input) }, 1)
	aut.AddTransition(1, func(input rune) bool { return input == '.' }, 2)
	aut.AddTransition(1, func(input rune) bool { return input == 'e' }, 4)
	aut.AddTransition(2, func(input rune) bool { return unicode.IsNumber(input) }, 3)
	aut.AddTransition(3, func(input rune) bool { return unicode.IsNumber(input) }, 3)
	aut.AddTransition(3, func(input rune) bool { return input == 'e' }, 4)
	aut.AddTransition(4, func(input rune) bool { return input == '-' || input == '+' }, 5)
	aut.AddTransition(4, func(input rune) bool { return unicode.IsNumber(input) }, 6)
	aut.AddTransition(5, func(input rune) bool { return unicode.IsNumber(input) }, 6)
	aut.AddTransition(6, func(input rune) bool { return unicode.IsNumber(input) }, 6)
  aut.AddTransition(0, func(input rune) bool { return input == '"' }, 7)
  aut.AddTransition(7, func(input rune) bool { return input != '"' }, 8)
  aut.AddTransition(8, func(input rune) bool { return input != '"' }, 8)
  aut.AddTransition(8, func(input rune) bool { return input == '"' }, 9)
  aut.AddTransition(0, func(input rune) bool { return input == 't' }, 10)
  aut.AddTransition(10, func(input rune) bool { return input == 'r' }, 11)
  aut.AddTransition(11, func(input rune) bool { return input == 'u' }, 12)
  aut.AddTransition(12, func(input rune) bool { return input == 'e' }, 13)
  aut.AddTransition(0, func(input rune) bool { return input == 'f' }, 14)
  aut.AddTransition(14, func(input rune) bool { return input == 'a' }, 15)
  aut.AddTransition(15, func(input rune) bool { return input == 'l' }, 16)
  aut.AddTransition(16, func(input rune) bool { return input == 's' }, 17)
  aut.AddTransition(17, func(input rune) bool { return input == 'e' }, 18)

	return aut
}
