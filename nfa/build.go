package nfa

type NFANewState = func() State

func NFAAnd(fa1, fa2 *NFA) *NFA {
	n := &NFA{
		Init:   fa1.Init,
		Accept: fa2.Accept,
		Marks:  fa1.Marks,
		Trans:  fa1.Trans,
	}
	for s, c := range fa2.Trans {
		n.Trans[s] = c
	}
	for s, c := range fa2.Marks {
		n.Marks[s] = c
	}
	n.AddConn(fa1.Accept, fa2.Init, Epsilon)
	return n
}

func NFAOr(fa1, fa2 *NFA, newState NFANewState) *NFA {
	i := newState()
	ac := newState()
	n := &NFA{
		Init:   i,
		Accept: ac,
		Marks:  fa1.Marks,
		Trans:  fa1.Trans,
	}
	for s, c := range fa2.Trans {
		n.Trans[s] = c
	}
	for s, c := range fa2.Marks {
		n.Marks[s] = c
	}
	n.AddConn(i, fa1.Init, Epsilon)
	n.AddConn(i, fa2.Init, Epsilon)
	n.AddConn(fa1.Accept, ac, Epsilon)
	n.AddConn(fa2.Accept, ac, Epsilon)
	return n
}

func NFARepeat(fa0 *NFA) *NFA {
	fa0.AddConn(fa0.Accept, fa0.Init, Epsilon)
	fa0.AddConn(fa0.Init, fa0.Accept, Epsilon)
	return fa0
}

func NFAMark(fa0 *NFA, mark int) *NFA {
	if _, ok := fa0.Marks[fa0.Accept]; !ok {
		fa0.Marks[fa0.Accept] = NewSet[int]()
	}
	fa0.Marks[fa0.Accept].Add(mark)
	return fa0
}

func NFAChar(c Char, newState NFANewState) *NFA {
	i := newState()
	ac := newState()
	n := &NFA{
		Init:   i,
		Accept: ac,
		Marks:  make(map[State]Set[int]),
		Trans:  map[State]*ClosureAndtrans{},
	}
	n.AddConn(i, ac, c)
	return n
}

func PostfixToNFA(seq TokenSeq) *NFA {
	s0 := 0
	newState := func() State {
		s0 += 1
		return State(s0)
	}
	stack := &Stack[*NFA]{s: make([]*NFA, 0)}
	for _, t := range seq {
		if t.IsChar() {
			stack.Push(NFAChar(t.Char(), newState))
		} else if t.Type() == REPEAT {
			stack.Push(NFARepeat(stack.Pop()))
		} else if t.Type() == MARK {
			stack.Push(NFAMark(stack.Pop(), t.Mark()))
		} else if t.Type() == AND {
			fa2 := stack.Pop()
			fa1 := stack.Pop()
			stack.Push(NFAAnd(fa1, fa2))
		} else if t.Type() == OR {
			stack.Push(NFAOr(stack.Pop(), stack.Pop(), newState))
		}
	}
	out := stack.Pop()
	if !stack.Empty() {
		panic("operate val num error")
	}
	return out
}
