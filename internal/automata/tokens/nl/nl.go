package nl

import (
	. "github.com/marcobarao/parser/internal/automata"
)

func NewNLAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, true)

	aut.AddTransition(0, func(input rune) bool { return input == '\n' }, 1)

	return aut
}
