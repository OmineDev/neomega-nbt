package dfa

import (
	"fmt"
)

type State int

type TransitEdge struct {
	TransitCond
	State
	Marks  []int
	Accept bool
}

func (te TransitEdge) String() string {
	return fmt.Sprintf("[%v]->%v", te.TransitCond.String(), te.State)
}

type TransitEdges []TransitEdge

func (es TransitEdges) AppendEdge(newEdge TransitEdge) TransitEdges {
	var i int
	var e TransitEdge
	for i, e = range es {
		if e.State == newEdge.State {
			es[i].TransitCond = es[i].TransitCond.Union(newEdge.TransitCond)
			return es
		}
	}
	return append(es, newEdge)
}

func (es TransitEdges) Feed(b byte) (s State, marks []int, accept bool, ok bool) {
	if len(es) == 0 {
		return s, nil, false, false
	}
	off := b / 64
	flg := (uint64(1) << (b % 64))
	for _, e := range es {
		if e.TransitCond[off]&flg == flg {
			return e.State, e.Marks, e.Accept, true
		}
	}
	return s, nil, false, false
}

type DFA struct {
	Init         State
	InitAccepted bool
	TransitTable []TransitEdges
}

func (dfa *DFA) String() string {
	s := fmt.Sprintf("Init: %v\n", dfa.Init)
	for src, edges := range dfa.TransitTable {
		s += fmt.Sprintf("  %v:\n", src)
		for _, edge := range edges {
			s += "    " + edge.String()
			if len(edge.Marks) > 0 {
				s += " " + fmt.Sprintf("%v  ", edge.Marks)
			}
			if edge.Accept {
				s += "*\n"
			} else {
				s += "\n"
			}
		}
	}
	return s
}

type DFARuntime struct {
	*DFA
	Ptr       State
	currEdges TransitEdges
	currMarks []int
	Marks     []func()
	isAccept  bool
}

func (rt *DFARuntime) Feed(c byte) (isAccept, canContinue bool) {
	rt.Ptr, rt.currMarks, rt.isAccept, canContinue = rt.currEdges.Feed(c)
	if len(rt.currMarks) != 0 && rt.Marks != nil {
		for _, m := range rt.currMarks {
			rt.Marks[m]()
		}
	}
	if canContinue {
		rt.currEdges = rt.DFA.TransitTable[rt.Ptr]
		// if len(rt.currEdges) == 0 {
		// 	canContinue = false
		// }
	}
	return rt.isAccept, canContinue
}

func (rt *DFARuntime) IsAccept() (accept bool) {
	return rt.isAccept
}

func (rt *DFARuntime) CanAccept(seq []byte) bool {
	rt.Reset()
	for _, c := range seq {
		if _, canContinue := rt.Feed(c); !canContinue {
			return false
		}
	}
	return rt.IsAccept()
}

func (rt *DFARuntime) Reset() {
	rt.Ptr = rt.DFA.Init
	rt.isAccept = rt.DFA.InitAccepted
	rt.currEdges = rt.DFA.TransitTable[rt.Ptr]
}

func (dfa *DFA) Runtime(marks []func()) *DFARuntime {
	rt := &DFARuntime{
		DFA:   dfa,
		Marks: marks,
	}
	rt.Reset()
	return rt
}
