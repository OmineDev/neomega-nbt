package nfa

type ExtendType byte

// const DigitalType = ExtendType(1)
// const AlphabetType = ExtendType(2)
// const AnyType = ExtendType(3)

var ExtendNames map[ExtendType]string = map[ExtendType]string{}

func (et ExtendType) String() string {
	ns, ok := ExtendNames[et]
	if !ok {
		panic(et)
	}
	return ns
	// switch et {
	// case DigitalType:
	// 	return "{0-9}"
	// case AlphabetType:
	// 	return "{a-z,A-Z}"
	// case AnyType:
	// 	return "{*}"
	// default:
	// 	panic(et)
	// }
}

var ExtendTransConds map[ExtendType]TransitCond = map[ExtendType]TransitCond{}

func (et ExtendType) TransitCond() TransitCond {
	cd, ok := ExtendTransConds[et]
	if !ok {
		panic(et)
	}
	return cd
	// switch et {
	// case DigitalType:
	// 	return DigitalTransCondict
	// case AlphabetType:
	// 	return AlphabetTransCondict
	// case AnyType:
	// 	return AnyTransCondict
	// default:
	// 	panic(et)
	// }
}

const (
	Epsilon = Char(EPSILONCHAR) << 16
	// Digital  = Char(EXTENDCHAR)<<16 | Char(DigitalType)
	// Alphabet = Char(EXTENDCHAR)<<16 | Char(AlphabetType)
	// Any      = Char(EXTENDCHAR)<<16 | Char(AnyType)
)

var SpecialCharMapping = map[byte]Char{
	' ': Epsilon,
	// 'd': Digital,
	// 'a': Alphabet,
	// '*': Any,
}

var extendCount = ExtendType(1)

func AddExtend(t byte, name string, cond TransitCond) {
	tp := extendCount
	extendCount += 1
	SpecialCharMapping[t] = Char(EXTENDCHAR)<<16 | Char(tp)
	ExtendTransConds[tp] = cond
	ExtendNames[tp] = name
}
