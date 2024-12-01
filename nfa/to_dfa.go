package nfa

import (
	"snbt/nfa/dfa"
	"sort"
)

func (nfa *NFA) ToDFA() *dfa.DFA {
	allNfaSates := NewSet[State]()
	allNfaSates.Add(nfa.Init)
	allNfaSates.Add(nfa.Accept)
	// allConds := NewSet[Condition]()

	for _, ts := range nfa.Trans {
		allNfaSates = allNfaSates.Union(ts.Closure)
		for _, e := range ts.Edges {
			allNfaSates.Add(e.State)
		}
	}

	allClosures := map[State]Set[State]{}
	allTrans := map[State]TransitEdges{}
	for s := range allNfaSates.Iter() {
		closure, trans := nfa.getClosureAndTrans(Set[State]{s: {}})
		allClosures[s] = closure
		allTrans[s] = trans
	}

	_dS := 0
	newDfaState := func() dfa.State {
		_dS += 1
		return dfa.State(_dS)
	}

	dI := dfa.State(0)
	dTable := map[dfa.State]dfa.TransitEdges{}
	d := &dfa.DFA{
		Init: dI,
	}

	dAccepts := map[dfa.State]bool{}
	dIEnclosure := allClosures[nfa.Init]
	if dIEnclosure.Has(nfa.Accept) {
		d.InitAccepted = true
		dAccepts[dI] = true
	}
	knownNfaClosures := map[string]Set[State]{dIEnclosure.Hash(): dIEnclosure}
	knownDfaSatets := map[string]dfa.State{dIEnclosure.Hash(): dI}
	enclosuresToCheck := NewSet[string]()
	enclosuresToCheck.Add(dIEnclosure.Hash())

	// move := func(srcs Set[State], cond Condition) (reach Set[State]) {
	// 	reach = NewSet[State]()
	// 	for src := range srcs {
	// 		dsts := allTrans[src][cond]
	// 		for d := range dsts {
	// 			reach = reach.Union(epsilonClosures[d])
	// 		}
	// 	}
	// 	return reach
	// }

	extend := func(srcs Set[State]) (groups map[string]TransitCond, hashLookUp map[string]Set[State]) {
		reachs := TransitEdges{}
		for src := range srcs.Iter() {
			trans := allTrans[src]
			for _, edge := range trans {
				cond := edge.TransitCond
				r := edge.State
				for d := range allClosures[r] {
					reachs = reachs.AppendEdge(TransitEdge{cond, d})
				}
			}
		}
		groups = map[string]TransitCond{}
		hashLookUp = map[string]Set[State]{}
		for i := 0; i < 256; i++ {
			b := byte(i)
			group := reachs.Feed(b)
			if group.Size() == 0 {
				continue
			}
			h := group.Hash()
			if _, ok := hashLookUp[h]; !ok {
				hashLookUp[h] = group
				groups[h] = TransitCond{}.Allow(b)
			} else {
				groups[h] = groups[h].Allow(b)
			}
		}
		return groups, hashLookUp
	}

	for enclosuresToCheck.Size() > 0 {
		eH, _ := enclosuresToCheck.Pop()
		srcClosure := knownNfaClosures[eH]
		sS := knownDfaSatets[eH]
		if _, ok := dTable[sS]; !ok {
			dTable[sS] = dfa.TransitEdges{}
		}
		groups, hashLookUp := extend(srcClosure)
		for dH, dstClosure := range hashLookUp {
			if _, ok := knownDfaSatets[dH]; !ok {
				dS := newDfaState()
				if dstClosure.Has(nfa.Accept) {
					dAccepts[dS] = true
				}
				knownNfaClosures[dH] = dstClosure
				knownDfaSatets[dH] = dS
				enclosuresToCheck.Add(dH)
			}
			dS := knownDfaSatets[dH]
			cond := groups[dH]
			var marks Set[int]
			for s, m := range nfa.Marks {
				if dstClosure.Has(s) {
					marks = marks.Union(m)
				}
			}
			ms := marks.ToList()
			sort.Ints(ms)
			if len(ms) == 0 {
				ms = nil
			}
			dTable[sS] = dTable[sS].AppendEdge(dfa.TransitEdge{TransitCond: dfa.TransitCond(cond), State: dS, Marks: ms, Accept: dAccepts[dS]})
		}
	}

	d.TransitTable = make([]dfa.TransitEdges, len(dTable))
	for i := range d.TransitTable {
		d.TransitTable[i] = dTable[dfa.State(i)]
	}
	return d
}

func GenDFAFromStr(rule string) *dfa.DFA {
	tokens := StringToTokenSeq(rule)
	tokens = InfixToPostfix(tokens)
	nfaData := PostfixToNFA(tokens)
	return nfaData.ToDFA()
}
