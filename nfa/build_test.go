package nfa

import (
	"fmt"
	"testing"
)

func TestBuildNFA(t *testing.T) {
	var runtime *NfaBroadFirstRuntime
	build := func(inp string) {
		tokens := StringToTokenSeq(inp)
		fmt.Println(tokens)
		tokens = InfixToPostfix(tokens)
		fmt.Println(tokens)
		nfaData := PostfixToNFA(tokens)
		runtime = NewNFABroadFirstRuntime(nfaData)
	}
	assertOk := func(seq string) {
		if !runtime.CanAccept([]byte(seq)) {
			t.FailNow()
		}
	}
	assertFail := func(seq string) {
		if runtime.CanAccept([]byte(seq)) {
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
