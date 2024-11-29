package main

type NumberMarkFns [12]func()

// gen_code.GenCodeFromStr("number", `((/ |-#)(((/d/d*)#(/ |(b#)|(s#)|(l#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |(E(/ |-#)(/d/d*)#))(/ |(f#)|(d#)))|(true#)|(false#)`)
func NumberFeed(state uint8, b uint8) (nextMarkTermAccept uint16) {
	jmpT := state | (b>>6)<<5
	cmp := 1 << ((b << 2) >> 2)
	switch jmpT {
	case 0:
		if cmp&70368744177664 > 0 {
			return 476
		}
		if cmp&287948901175001088 > 0 {
			return 517
		}
		if cmp&35184372088832 > 0 {
			return 1280
		}
	case 1:
		if cmp&287948901175001088 > 0 {
			return 1813
		}
	case 2:
		if cmp&70368744177664 > 0 {
			return 4629
		}
		if cmp&287948901175001088 > 0 {
			return 4869
		}
	case 5:
		if cmp&70368744177664 > 0 {
			return 476
		}
		if cmp&287948901175001088 > 0 {
			return 517
		}
	case 7:
		if cmp&287948901175001088 > 0 {
			return 2581
		}
	case 10:
		if cmp&287948901175001088 > 0 {
			return 2581
		}
	case 11:
		if cmp&35184372088832 > 0 {
			return 3608
		}
		if cmp&287948901175001088 > 0 {
			return 3869
		}
	case 14:
		if cmp&287948901175001088 > 0 {
			return 3869
		}
	case 15:
		if cmp&287948901175001088 > 0 {
			return 6173
		}
	case 18:
		if cmp&287948901175001088 > 0 {
			return 5397
		}
	case 19:
		if cmp&70368744177664 > 0 {
			return 4629
		}
		if cmp&287948901175001088 > 0 {
			return 4869
		}
	case 21:
		if cmp&287948901175001088 > 0 {
			return 5653
		}
	case 22:
		if cmp&287948901175001088 > 0 {
			return 5653
		}
	case 24:
		if cmp&287948901175001088 > 0 {
			return 6173
		}
	case 32:
		if cmp&274877906944 > 0 {
			return 988
		}
		if cmp&4503599627370496 > 0 {
			return 1244
		}
	case 34:
		if cmp&17592186044416 > 0 {
			return 4113
		}
		if cmp&2251799813685248 > 0 {
			return 4365
		}
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&17179869184 > 0 {
			return 5129
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 35:
		if cmp&8589934592 > 0 {
			return 6108
		}
	case 36:
		if cmp&1125899906842624 > 0 {
			return 1756
		}
	case 38:
		if cmp&9007199254740992 > 0 {
			return 2268
		}
	case 39:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 40:
		if cmp&137438953472 > 0 {
			return 2347
		}
	case 42:
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
		if cmp&32 > 0 {
			return 3036
		}
	case 47:
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 48:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 49:
		if cmp&274877906944 > 0 {
			return 3363
		}
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
	case 50:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 51:
		if cmp&274877906944 > 0 {
			return 3363
		}
		if cmp&17592186044416 > 0 {
			return 4113
		}
		if cmp&2251799813685248 > 0 {
			return 4365
		}
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&17179869184 > 0 {
			return 5129
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
	case 52:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 53:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 54:
		if cmp&32 > 0 {
			return 3036
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
		if cmp&274877906944 > 0 {
			return 3363
		}
	case 55:
		if cmp&17592186044416 > 0 {
			return 6620
		}
	case 56:
		if cmp&274877906944 > 0 {
			return 3363
		}
		if cmp&68719476736 > 0 {
			return 3111
		}
	case 57:
		if cmp&2251799813685248 > 0 {
			return 6876
		}
	case 58:
		if cmp&137438953472 > 0 {
			return 6959
		}
	}
	return 2
}
