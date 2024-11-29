package gen_code

import (
	"fa/nfa/dfa"
)

func compueteMaxMarks(table []dfa.TransitEdges) (maxMarks int, maxMarkValue int) {
	maxMarks = 0
	for _, edges := range table {
		for _, edge := range edges {
			if len(edge.Marks) > maxMarks {
				maxMarks = len(edge.Marks)
			}
			for _, m := range edge.Marks {
				if m > maxMarkValue {
					maxMarkValue = m
				}
			}
		}
	}
	return maxMarks, maxMarkValue
}

func computeShift(states int) int {
	shift := 0
	for states != 0 {
		states >>= 1
		shift += 1
	}
	return shift
}

type FlattenEdge struct {
	cond              uint64
	nextState         uint8
	mark              uint8
	nextStateIsAccept bool
	nextStateIsTerm   bool
}

func (fe *FlattenEdge) genRetCode() uint16 {
	tCode := uint16(0)
	if fe.nextStateIsTerm {
		tCode = 2
	}
	if fe.nextStateIsAccept {
		tCode += 1
	}
	return uint16(fe.nextState)<<8 | uint16(((fe.mark<<2)>>2))<<2 | tCode
}

func computeFlattenTransTable(table []dfa.TransitEdges) (
	shift int,
	jmps map[int][]*FlattenEdge,
	maxMarkValue int,
) {
	shift = computeShift(len(table))
	var maxMarks int
	maxMarks, maxMarkValue = compueteMaxMarks(table)
	if maxMarks > 1 || len(table) > 255 || maxMarkValue > 63 {
		panic("allow only max marks<=1 && states <=255 && maxMarkValue <=63")
	}
	termState := map[uint8]struct{}{}
	for stateVal, egdes := range table {
		if len(egdes) == 0 {
			termState[uint8(stateVal)] = struct{}{}
		}
	}
	jmps = map[int][]*FlattenEdge{}
	for stateVal, edges := range table {
		if len(edges) == 0 {
			continue
		}
		for _, edge := range edges {
			state := uint8(edge.State)
			_, isTerm := termState[state]
			mark := ^uint8(8)
			if len(edge.Marks) > 0 {
				mark = uint8(edge.Marks[0])
			}
			for h2, cond := range edge.TransitCond {
				if cond == 0 {
					continue
				}
				jmpT := h2<<shift | stateVal
				if _, ok := jmps[jmpT]; !ok {
					jmps[jmpT] = make([]*FlattenEdge, 0)
				}
				flattenEdge := &FlattenEdge{
					cond:              cond,
					nextState:         state,
					mark:              mark,
					nextStateIsAccept: edge.Accept,
					nextStateIsTerm:   isTerm,
				}
				jmps[jmpT] = append(jmps[jmpT], flattenEdge)
			}
		}
	}
	return
}

// func genRetStr(name string, State dfa.State, Accept bool, marks []int) string {
// 	if maxMarks > 1 {
// 		markStr := "nil"
// 		if len(marks) != 0 {
// 			markStr = "[]int{"
// 			for i, m := range marks {
// 				if i != 0 {
// 					markStr += ", "
// 				}
// 				markStr += fmt.Sprintf("%v", m)
// 			}
// 			markStr += "}"
// 		}
// 		return fmt.Sprintf("return %vDFAState(%v), %v, true, %v\n		}\n", name, State, Accept, markStr)
// 	} else {
// 		v := 255
// 		if len(marks) != 0 {
// 			v = marks[0]
// 		}
// 		ret := uint32(State)
// 		ret <<= 8
// 		if Accept {
// 			ret |= 1
// 		}
// 		ret <<= 8
// 		ret |= 1
// 		ret <<= 8
// 		ret |= uint32(v)
// 		return fmt.Sprintf("			return %v\n		}\n", ret)
// 	}

// }
