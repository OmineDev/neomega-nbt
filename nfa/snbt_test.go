package nfa

import (
	"fmt"
	"testing"
)

func TestSNBT(t *testing.T) {
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
	dfa := GenDFAFromStr(`((/ |-)(((/d/d*)#)|(((/d/d*./d*)|(/d*./d/d*))#))(/ |(E(/ |-)/d/d*)#)(/ |(b#)|(s#)|(L#)|(f#)|(d#)))|(true#)|(false#)`)
	runtime := dfa.Runtime(nil)

	FInterger := false
	FDotFloat := false
	FScientisc := false
	FT := ""
	FTrue := false
	FFalse := false
	reset := func() {
		FInterger = false
		FDotFloat = false
		FScientisc = false
		FT = ""
		FTrue = false
		FFalse = false
	}
	runtime.Marks = append(runtime.Marks, func() {
		FInterger = true
	})
	runtime.Marks = append(runtime.Marks, func() {
		FDotFloat = true
	})
	runtime.Marks = append(runtime.Marks, func() {
		FScientisc = true
	})
	runtime.Marks = append(runtime.Marks, func() {
		if FT != "" {
			t.FailNow()
		}
		FT = "B"
	})
	runtime.Marks = append(runtime.Marks, func() {
		if FT != "" {
			t.FailNow()
		}
		FT = "S"
	})
	runtime.Marks = append(runtime.Marks, func() {
		if FT != "" {
			t.FailNow()
		}
		FT = "L"
	})
	runtime.Marks = append(runtime.Marks, func() {
		if FT != "" {
			t.FailNow()
		}
		FT = "F"
	})
	runtime.Marks = append(runtime.Marks, func() {
		if FT != "" {
			t.FailNow()
		}
		FT = "D"
	})
	runtime.Marks = append(runtime.Marks, func() {
		FTrue = true
	})
	runtime.Marks = append(runtime.Marks, func() {
		FFalse = true
	})
	t.Log(dfa)
	assertOk := func(seq string) {
		runtime.Reset()
		reset()
		for _, c := range []byte(seq) {

			if _, canContinue := runtime.Feed(c); !canContinue {
				t.Log(seq)
				t.FailNow()
			}
			fmt.Printf("%s -> %v\n", string([]byte{c}), runtime.Ptr)
		}
		if !runtime.IsAccept() {
			t.Log(seq)
			t.FailNow()
		}
	}
	assertInterger := func() {
		if FDotFloat == true || FInterger == false || FScientisc == true || FTrue == true || FFalse == true {
			t.FailNow()
		}
	}
	assertFloatByDot := func() {
		if FDotFloat == false || FScientisc == true || FTrue == true || FFalse == true {
			t.FailNow()
		}
	}
	assertFloatBySci := func() {
		if FScientisc == false || FTrue == true || FFalse == true {
			t.FailNow()
		}
	}
	assertByteTrue := func() {
		if FInterger == true || FScientisc == true || FDotFloat == true || FTrue == false || FFalse == true {
			t.FailNow()
		}
	}
	assertByteFalse := func() {
		if FInterger == true || FScientisc == true || FDotFloat == true || FTrue == true || FFalse == false {
			t.FailNow()
		}
	}
	assertPrefix := func(p string) {
		if FT != p {
			t.FailNow()
		}
	}
	assertOk("-1b")
	assertInterger()
	assertPrefix("B")
	assertOk("2s")
	assertInterger()
	assertPrefix("S")
	assertOk("-3L")
	assertInterger()
	assertPrefix("L")
	assertOk("4.0f")
	assertFloatByDot()
	assertPrefix("F")
	assertOk("5.0d")
	assertFloatByDot()
	assertPrefix("D")
	assertOk("123")
	assertInterger()
	assertPrefix("")
	assertOk("-1.0")
	assertPrefix("")
	assertFloatByDot()
	assertPrefix("")
	assertOk("true")
	assertPrefix("")
	assertByteTrue()
	assertOk("false")
	assertByteFalse()
	assertPrefix("")
	assertOk(".2")
	assertFloatByDot()
	assertPrefix("")
	assertOk("2.")
	assertFloatByDot()
	assertPrefix("")
	assertOk("12.00")
	assertFloatByDot()
	assertPrefix("")
	assertOk("-.20")
	assertFloatByDot()
	assertPrefix("")
	assertOk("20.")
	assertFloatByDot()
	assertPrefix("")
	assertOk("-123E10")
	assertFloatBySci()
	assertPrefix("")
	assertOk("-0.2E10f")
	assertFloatBySci()
	assertPrefix("F")
	// assertOk("abc")
	// t.Log("ok")
	// t.FailNow()
}
