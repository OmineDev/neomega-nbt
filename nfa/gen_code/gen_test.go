package gen_code

import (
	"os"
	"snbt/nfa"
	"testing"
)

func TestGenCode(t *testing.T) {
	dfa := nfa.GenDFAFromStr(`((/ |-)(((/d/d*)#)|(((/d/d*./d*)|(/d*./d/d*))#))(/ |(E(/ |-)/d/d*)#)(/ |(b#)|(s#)|(L#)|(f#)|(d#)))|(true#)|(false#)`)
	table := dfa.TransitTable
	s := GenCode("number", table)
	os.WriteFile("/Users/dai/projects/fa/compact_dfa_runner/out.go", []byte(s), 0755)
}
