package states

import (
	"bytes"
	"fmt"
)

/*
The set of transitions of a state.
*/
type Transitions struct {
	//key: symbol
	transitions map[string]*State

	symbols []string

	list []Transition
}

func NewTransitions(symbols []string) *Transitions {
	return &Transitions{
		transitions: make(map[string]*State),
		symbols:     symbols,
	}
}

func NewTransitionsList(trans []Transition, symbols []string) *Transitions {
	transitions := NewTransitions(symbols)
	transitions.list = trans
	for _, t := range trans {
		transitions.transitions[t.Sym] = t.State
	}
	return transitions
}

func (this *Transitions) List() []Transition {
	if len(this.list) != len(this.transitions) {
		this.list = make([]Transition, 0, len(this.transitions))
		for _, sym := range this.symbols {
			if st, exist := this.transitions[sym]; exist {
				this.list = append(this.list, Transition{State: st, Sym: sym})
			}
		}
	}
	return this.list
}

func (this *Transitions) Replace(sym string, state *State) {
	this.transitions[sym] = state
	this.list = nil
}

func (this *Transitions) String() string {
	w := new(bytes.Buffer)
	for _, t := range this.List() {
		fmt.Fprintf(w, "%s: S%d\n", t.Sym, t.State.Number)
	}
	return w.String()
}

func (this *Transitions) Transition(sym string) *State {
	return this.transitions[sym]
}
