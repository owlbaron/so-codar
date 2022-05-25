package automata

import (
	. "github.com/marcobarao/parser/internal/automata/errors"
)

type MatchFunc func(input rune) bool

type transition struct {
	match        func(input rune) bool
	destinyState uint
}

type Automata struct {
	states       map[uint]bool
	transitions  map[uint][]transition
	startState   uint
	acceptStates map[uint]bool

	currentState uint
}

func NewAutomata(startState uint, isAcceptState bool) *Automata {

	aut := &Automata{
		states:       make(map[uint]bool),
		transitions:  make(map[uint][]transition),
		startState:   startState,
		acceptStates: make(map[uint]bool),

		currentState: startState,
	}

	aut.AddState(startState, isAcceptState)

	return aut
}

func (a *Automata) AddState(state uint, isAcceptState bool) {
	a.states[state] = true

	if isAcceptState {
		a.acceptStates[state] = true
	}
}

func (a *Automata) AddTransition(sourceState uint, matchFunc MatchFunc, destinyState uint) error {
	if _, exists := a.states[sourceState]; !exists {
		return StateNotFoundError{State: sourceState}
	}

	if _, exists := a.states[destinyState]; !exists {
		return StateNotFoundError{State: destinyState}
	}

	a.transitions[sourceState] = append(a.transitions[sourceState], transition{match: matchFunc, destinyState: destinyState})

	return nil
}

func (a *Automata) Input(input rune) error {
	found := false
	var t transition
	for _, t = range a.transitions[a.currentState] {
		if t.match(input) {
			found = true
			break
		}
	}

	if found {
		a.currentState = t.destinyState
	} else {
		return TransitionNotFoundError{State: a.currentState, Input: input}
	}

	return nil
}

func (a *Automata) Accepted() bool {
	return !!a.acceptStates[a.currentState]
}
