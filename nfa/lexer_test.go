package nfa

import "testing"

func TestLexar(t *testing.T) {
	inp := `ab(ab|(a\*c|a*d))e`
	tokens := StringToTokenSeq(inp)
	t.Log(tokens.String())
	if tokens.String() != "a[&]b[&][(]a[&]b[|][(]a[&]*[&]c[|]a[*][&]d[)][)][&]e" {
		t.FailNow()
	}
}

func TestInfixToPostfix(t *testing.T) {
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
	if tokens.String() != "{0-9}{a-z,A-Z}[&]{ε}[&]{*}[&]" {
		t.FailNow()
	}
}
