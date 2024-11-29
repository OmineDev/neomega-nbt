package gen_code

import (
	"fa/nfa"
	"fa/nfa/dfa"
	"fmt"
	"sort"
	"strings"
)

func GenCode(name string, table []dfa.TransitEdges) string {
	shift, jmps, maxMarkVal := computeFlattenTransTable(table)
	ord := []int{}
	{
		for k := range jmps {
			ord = append(ord, k)
		}
		sort.Ints(ord)
	}
	name = strings.ToUpper(string(name[0])) + string(name[1:])
	s := "package main\n\n"
	s += fmt.Sprintf("type %vMarkFns [%v]func()\n\n", name, maxMarkVal+1)
	s += fmt.Sprintf("func %vFeed(state uint8, b uint8) (nextMarkTermAccept uint16){\n", name)
	if shift <= 5 {
		s += fmt.Sprintf("	jmpT:=state|(b>>6)<<%v\n", shift)
	} else {
		s += fmt.Sprintf("	jmpT:=uint16(state)|(uint16(b)>>6)<<%v\n", shift)
	}
	s += "	cmp:=uint64(1)<<((b<<2)>>2)\n"
	s += "    switch jmpT{\n"

	for _, jmpT := range ord {
		s += fmt.Sprintf("	case %v:\n", jmpT)
		edges := jmps[jmpT]
		for _, edge := range edges {
			s += fmt.Sprintf("        if cmp & %v > 0 {\n", edge.cond)
			s += fmt.Sprintf("        return %v\n        }\n", edge.genRetCode())
		}

	}
	s += "    }\n    return 2\n}\n"
	return s
}

func GenCodeFromStr(name string, re string) string {
	dfa := nfa.GenDFAFromStr(re)
	table := dfa.TransitTable
	s := GenCode(name, table)
	return s
}
