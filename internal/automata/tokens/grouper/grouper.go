package grouper

import (
	. "github.com/marcobarao/parser/internal/automata"
)

func NewGrouperAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, true)

	aut.AddTransition(0, func(input rune) bool { return input == '(' || input == ')' || input == '{' || input == '}' || input == ';' }, 1)

	return aut
}
