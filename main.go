package main

import (
	"fa/nfa/gen_code"
	"os"
)

func main() {
	// `((/ |-#)(((/d/d*)#)|(((/d/d*./d*)|(/d*./d/d*))#))(/ |(E(/ |-)/d/d*)#)(/ |(b#)|(s#)|(L#)|(f#)|(d#)))|(true#)|(false#)`
	code := gen_code.GenCodeFromStr("number", `((/ |-#)(((/d/d*)#(/ |(b#)|(s#)|(l#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |(E(/ |-#)(/d/d*)#))(/ |(f#)|(d#)))|(true#)|(false#)`)
	os.WriteFile("gen.go?", []byte(code), 0755)
}
