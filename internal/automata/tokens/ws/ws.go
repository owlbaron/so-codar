package ws

import (
	. "github.com/marcobarao/parser/internal/automata"
)

func NewWSAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, true)

	aut.AddTransition(0, func(input rune) bool { return input == ' ' || input == '\t' }, 1)
	aut.AddTransition(1, func(input rune) bool { return input == ' ' || input == '\t' }, 1)

	return aut
}
