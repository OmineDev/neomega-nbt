package nfa

import "testing"

func TestLexar(t *testing.T) {
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
	inp := `ab(ab|(a\*c|a*d))e`
	tokens := StringToTokenSeq(inp)
	t.Log(tokens.String())
	if tokens.String() != "a[&]b[&][(]a[&]b[|][(]a[&]*[&]c[|]a[*][&]d[)][)][&]e" {
		t.FailNow()
	}
}

func TestInfixToPostfix(t *testing.T) {
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
	inp := `ab(ab|(a\*c|a*d))e`
	tokens := StringToTokenSeq(inp)
	tokens = InfixToPostfix(tokens)
	if tokens.String() != "ab[&]ab[&]a*[&]c[&]a[*]d[&][|][|][&]e[&]" {
		t.FailNow()
	}

	inp = `a/ b(a/ b|(a\*c|a*d))e`
	tokens = StringToTokenSeq(inp)
	tokens = InfixToPostfix(tokens)
	if tokens.String() != "a{ε}[&]b[&]a{ε}[&]b[&]a*[&]c[&]a[*]d[&][|][|][&]e[&]" {
		t.FailNow()
	}
	inp = `/d/a/ /*`
	tokens = StringToTokenSeq(inp)
	tokens = InfixToPostfix(tokens)
	if tokens.String() != "{0-9}{a-z,A-Z}[&]{ε}[&]{any}[&]" {
		t.Logf(tokens.String())
		t.FailNow()
	}
}
