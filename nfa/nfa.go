package nfa

import (
	"fmt"
)

type TransitEdge struct {
	TransitCond
	State
}

func NewTransitEdgeFromCommon(b byte, dst State) TransitEdge {
	return TransitEdge{TransitCond: TransitCond{}.Allow(b), State: dst}
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

func (es TransitEdges) Feed(b byte) (s Set[State]) {
	for _, e := range es {
		if e.TransitCond.Accept(b) {
			if s == nil {
				s = Set[State]{}
			}
			s.Add(e.State)
		}
	}
	return s
}

type ClosureAndtrans struct {
	Closure Set[State]
	Edges   TransitEdges
}

func (ct *ClosureAndtrans) AddConn(d State, cond Char) {
	if cond.IsEpsilon() {
		ct.Closure.Add(d)
	} else if cond.IsCommon() {
		ct.Edges = ct.Edges.AppendEdge(NewTransitEdgeFromCommon(cond.Common(), d))
	} else {
		ct.Edges = ct.Edges.AppendEdge(TransitEdge{cond.Extend().TransitCond(), d})
	}
}

func NewClosureAndTrans() *ClosureAndtrans {
	return &ClosureAndtrans{
		Closure: NewSet[State](),
		Edges:   make(TransitEdges, 0),
	}
}

type NFA struct {
	Init   State
	Accept State
	Marks  map[State]Set[int]
	Trans  map[State]*ClosureAndtrans
}

type Rule struct {
	src, dst State
	cond     Char
}

func (nfa *NFA) AddConn(s, d State, cond Char) {
	if ct, ok := nfa.Trans[s]; !ok {
		ct = NewClosureAndTrans()
		ct.AddConn(d, cond)
		nfa.Trans[s] = ct
	} else {
		ct.AddConn(d, cond)
	}
}

func (nfa *NFA) AddRule(r Rule) {
	nfa.AddConn(r.src, r.dst, r.cond)
}

// func (nfa *NFA) IterRules() func(func(Rule) bool) {
// 	return func(fn func(Rule) bool) {
// 		for s, ed := range nfa.Trans {
// 			for cond, ds := range ed {
// 				for dn := range ds {
// 					if !fn(Rule{s, dn, cond}) {
// 						return
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func (nfa *NFA) getClosureAndTrans(curr Set[State]) (closure Set[State], trans TransitEdges) {
	checked := NewSet[State]()
	toCheck := NewSet[State]()
	for c := range curr {
		toCheck.Add(c)
	}
	transTable := TransitEdges{}
	for toCheck.Size() != 0 {
		n, _ := toCheck.Pop()
		checked.Add(n)
		trans := nfa.Trans[n]
		if trans != nil && trans.Closure != nil {
			for next := range trans.Closure {
				if checked.Has(next) || toCheck.Has(next) {
					continue
				}
				toCheck.Add(next)
			}
		}
		if trans != nil && trans.Edges != nil {
			for _, edge := range trans.Edges {
				transTable = transTable.AppendEdge(edge)
			}
		}
	}
	return checked, transTable
}

type NfaBroadFirstRuntime struct {
	curr      Set[State]
	currTrans TransitEdges
	*NFA
}

func (rt *NfaBroadFirstRuntime) IsAccept() bool {
	_, ok := rt.curr[rt.NFA.Accept]
	return ok
}

func (rt *NfaBroadFirstRuntime) CheckMark() Set[int] {
	var marks Set[int]
	for s := range rt.curr {
		m := rt.NFA.Marks[s]
		marks = marks.Union(m)
	}
	return marks
}

func (rt *NfaBroadFirstRuntime) Feed(c byte) (canContinue bool) {
	if next := rt.currTrans.Feed(c); next.Size() == 0 {
		return false
	} else {
		rt.move(next)
		return true
	}
}

func (rt *NfaBroadFirstRuntime) CanAccept(seq []byte) bool {
	rt.Reset()
	defer fmt.Println()
	for _, c := range seq {
		if canContinue := rt.Feed(c); !canContinue {
			return false
		}
		marks := rt.CheckMark()
		fmt.Print(string([]byte{c}))
		if marks.Size() != 0 {
			fmt.Printf("[trap mark: %v]", marks.Hash())
		}
	}
	return rt.IsAccept()
}

func (rt *NfaBroadFirstRuntime) move(curr Set[State]) {
	rt.curr, rt.currTrans = rt.getClosureAndTrans(curr)
}

func (rt *NfaBroadFirstRuntime) Reset() {
	rt.curr = NewSet[State]()
	rt.curr.Add(rt.Init)
	rt.move(rt.curr)
}

func NewNFABroadFirstRuntime(nfa *NFA) *NfaBroadFirstRuntime {
	rt := &NfaBroadFirstRuntime{
		curr: NewSet[State](),
		NFA:  nfa,
	}
	rt.Reset()
	return rt
}
