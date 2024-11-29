package main

import (
	"fa/nfa"
	"fa/nfa/gen_code"
	"os"
)

func main() {
	var DigitalTransCondict = nfa.TransitCond{}
	for i := '0'; i <= '9'; i++ {
		DigitalTransCondict = DigitalTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('d', "{0-9}", DigitalTransCondict)

	var AlphabetTransCondict = nfa.TransitCond{}
	for i := 'a'; i <= 'z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('a', "{a-z,A-Z}", AlphabetTransCondict)

	var AnyTransCondict = nfa.TransitCond{}
	for i := 0; i <= 255; i++ {
		AnyTransCondict = AnyTransCondict.Allow(byte(i))
	}
	nfa.AddExtend('a', "{any}", AnyTransCondict)

	code := gen_code.GenCodeFromStr("number", `((/ |-#)(((/d/d*)#(/ |(b#)|(s#)|(l#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |(E(/ |-#)(/d/d*)#))(/ |(f#)|(d#)))|(true#)|(false#)`)
	os.WriteFile("number_core.go", []byte(code), 0755)

	var WhiteSpaceTransCondict = nfa.TransitCond{}
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\t')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\n')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\v')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\f')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow('\r')
	WhiteSpaceTransCondict = WhiteSpaceTransCondict.Allow(' ')
	nfa.AddExtend('b', "{whiteSpace}", WhiteSpaceTransCondict)

	code = gen_code.GenCodeFromStr("whiteSpace", `/b*`)
	os.WriteFile("white_space_core.go", []byte(code), 0755)
}
