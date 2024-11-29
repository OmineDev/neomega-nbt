package nfa

func StringToTokenSeq(inp string) TokenSeq {
	bs := []byte(inp)
	out := []Token{}
	lastIsVal := false
	mi := 0
	for i := 0; i < len(bs); i++ {
		switch bs[i] {
		case '|':
			out = append(out, Or)
			lastIsVal = false
		case '(':
			if lastIsVal {
				out = append(out, And)
			}
			lastIsVal = false
			out = append(out, LBr)
		case ')':
			out = append(out, RBr)
			lastIsVal = true
		case '*':
			out = append(out, Repeat)
		case '#':
			out = append(out, NewMark(mi))
			mi += 1
		case '\\':
			if lastIsVal {
				out = append(out, And)
			}
			lastIsVal = true
			out = append(out, CommonCharToken(bs[i+1]))
			i++
		case '/':
			if lastIsVal {
				out = append(out, And)
			}
			lastIsVal = true
			out = append(out, Token(SpecialCharMapping[bs[i+1]]))
			i++
		default:
			if lastIsVal {
				out = append(out, And)
			}
			lastIsVal = true
			out = append(out, CommonCharToken(bs[i]))
		}
	}
	return TokenSeq(out)
}

func InfixToPostfix(ts TokenSeq) TokenSeq {
	// out := []Token{}
	tmpStack := Stack[Token]{s: make([]Token, 0)}
	opStack := Stack[Token]{s: make([]Token, 0)}
	for _, t := range ts {
		// fmt.Printf("%v %v %v\n", t.String(), TokenSeq(opStack.s), TokenSeq(tmpStack.s))
		if t.IsChar() {
			tmpStack.Push(t)
		} else if t.Type() == LBR {
			opStack.Push(t)
		} else if t.Type() == RBR {
			for {
				_t := opStack.Pop()
				if _t.Type() == LBR {
					break
				}
				tmpStack.Push(_t)
			}
		} else {
			for !opStack.Empty() {
				if uint8(t.Type()) <= uint8(opStack.Curr().Type()) {
					_t := opStack.Pop()
					tmpStack.Push(_t)
				} else {
					break
				}
			}
			opStack.Push(t)
		}
		// fmt.Printf("> %v %v\n", TokenSeq(opStack.s), TokenSeq(tmpStack.s))
	}
	valCount := 0
	for !opStack.Empty() {
		t := opStack.Pop()
		if t.IsChar() {
			valCount += 1
		} else if t.Type() == OR {

		}
		tmpStack.Push(t)
	}

	return TokenSeq(tmpStack.s)
}
