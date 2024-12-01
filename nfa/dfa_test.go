package nfa

import (
	"encoding/json"
	"fa/nfa/dfa"
	"fmt"
	"testing"
)

func TestToDFA(t *testing.T) {
	var DigitalTransCondict = TransitCond{}
	for i := '0'; i <= '9'; i++ {
		DigitalTransCondict = DigitalTransCondict.Allow(byte(i))
	}
	AddExtend('d', "{0-9}", DigitalTransCondict)

	var AlphabetTransCondict = TransitCond{}
	for i := 'a'; i <= 'z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	AddExtend('a', "{a-z,A-Z}", AlphabetTransCondict)

	var AnyTransCondict = TransitCond{}
	for i := 0; i <= 255; i++ {
		AnyTransCondict = AnyTransCondict.Allow(byte(i))
	}
	AddExtend('*', "{any}", AnyTransCondict)
	var runtime *dfa.DFARuntime
	build := func(inp string) {
		tokens := StringToTokenSeq(inp)
		tokens = InfixToPostfix(tokens)
		nfaData := PostfixToNFA(tokens)
		dfaData := nfaData.ToDFA()
		t.Log(inp, dfaData)
		runtime = &dfa.DFARuntime{DFA: dfaData}
		bs, _ := json.Marshal(dfaData)
		fmt.Println(string(bs))
	}
	assertOk := func(seq string) {
		if !runtime.CanAccept([]byte(seq)) {
			t.Log(seq)
			t.FailNow()
		}
	}
	assertFail := func(seq string) {
		if runtime.CanAccept([]byte(seq)) {
			t.Log(seq)
			t.FailNow()
		}
	}
	build(`(ab(ac#)#e)*`)
	assertOk("abace")
	build(`(a/ b(a/ b#(a/ bd|(ab\*c|a*d)))#e)*`)
	assertOk("")
	assertOk("abababde")
	assertOk("ababde")
	assertOk("ababaaaaade")
	assertOk("ababab*ce")
	assertFail("abababce")
	build(`a/ b/d*/a*/**`)
	assertOk("ab123cd")
	build(`(/d|/*)*`)
	assertOk("ab123cd")
	build(`(/d|/a)*`)
	assertOk("ab123cd")
	assertFail("ab1  3cd")
}
