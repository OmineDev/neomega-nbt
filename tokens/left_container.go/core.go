package left_container

type LeftContainerMarkFns [7]func()

func LeftContainerFeed(state uint8, b uint8) (nextMarkTermAccept uint16) {
	jmpT := state | (b>>6)<<4
	cmp := 1 << ((b << 2) >> 2)
	switch jmpT {
	case 0:
		if cmp&17179869184 > 0 {
			return 259
		}
		if cmp&549755813888 > 0 {
			return 519
		}
	case 5:
		if cmp&576460752303423488 > 0 {
			return 2583
		}
	case 6:
		if cmp&576460752303423488 > 0 {
			return 2331
		}
	case 7:
		if cmp&576460752303423488 > 0 {
			return 2067
		}
	case 16:
		if cmp&134217728 > 0 {
			return 781
		}
		if cmp&576460752303423488 > 0 {
			return 1035
		}
	case 19:
		if cmp&512 > 0 {
			return 1500
		}
		if cmp&4096 > 0 {
			return 1756
		}
		if cmp&4 > 0 {
			return 2012
		}
	}
	return 2
}
