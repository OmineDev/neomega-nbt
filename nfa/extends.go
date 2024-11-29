package nfa

type ExtendType byte

const DigitalType = ExtendType(1)
const AlphabetType = ExtendType(2)
const AnyType = ExtendType(3)

func (et ExtendType) String() string {
	switch et {
	case DigitalType:
		return "{0-9}"
	case AlphabetType:
		return "{a-z,A-Z}"
	case AnyType:
		return "{*}"
	default:
		panic(et)
	}
}

var DigitalTransCondict = TransitCond{}
var AlphabetTransCondict = TransitCond{}
var AnyTransCondict = TransitCond{}

func init() {
	for i := '0'; i <= '9'; i++ {
		DigitalTransCondict = DigitalTransCondict.Allow(byte(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		AlphabetTransCondict = AlphabetTransCondict.Allow(byte(i))
	}
	for i := 0; i <= 255; i++ {
		AnyTransCondict = AnyTransCondict.Allow(byte(i))
	}
}

func (et ExtendType) TransitCond() TransitCond {
	switch et {
	case DigitalType:
		return DigitalTransCondict
	case AlphabetType:
		return AlphabetTransCondict
	case AnyType:
		return AnyTransCondict
	default:
		panic(et)
	}
}

const (
	Epsilon  = Char(EPSILONCHAR) << 16
	Digital  = Char(EXTENDCHAR)<<16 | Char(DigitalType)
	Alphabet = Char(EXTENDCHAR)<<16 | Char(AlphabetType)
	Any      = Char(EXTENDCHAR)<<16 | Char(AnyType)
)

var SpecialCharMapping = map[byte]Char{
	' ': Epsilon,
	'd': Digital,
	'a': Alphabet,
	'*': Any,
}
