package identifier

import (
	"unicode"

	. "github.com/marcobarao/parser/internal/automata"
)

func NewIdentifierAutomata() *Automata {
	aut := NewAutomata(0, false)

	aut.AddState(1, true)

	aut.AddTransition(0, func(input rune) bool { return unicode.IsLetter(input) }, 1)
	aut.AddTransition(1, func(input rune) bool { return unicode.IsLetter(input) || unicode.IsNumber(input) }, 1)

	return aut
}
