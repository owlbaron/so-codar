package operator

import (
	. "github.com/marcobarao/parser/internal/automata"
)

func NewOperatorAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, false)
	aut.AddState(2, true)
	aut.AddState(3, true)
	aut.AddState(4, true)

	aut.AddTransition(0, func(input rune) bool { return input == '+' || input == '-' || input == '/' || input == '%' }, 4)
	aut.AddTransition(0, func(input rune) bool { return input == '!' }, 1)
	aut.AddTransition(1, func(input rune) bool { return input == '=' }, 4)
	aut.AddTransition(0, func(input rune) bool { return input == '*' }, 2)
	aut.AddTransition(2, func(input rune) bool { return input == '*' }, 4)
	aut.AddTransition(2, func(input rune) bool { return input == '*' }, 4)
	aut.AddTransition(0, func(input rune) bool { return input == '<' || input == '>' || input == '=' }, 3)
	aut.AddTransition(3, func(input rune) bool { return input == '=' }, 4)

	return aut
}
