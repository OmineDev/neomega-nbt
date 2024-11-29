package number

type NumberMarkFns [12]func()

func NumberFeed(state uint8, b uint8) (nextMarkTermAccept uint16) {
	jmpT := state | (b>>6)<<5
	cmp := 1 << ((b << 2) >> 2)
	switch jmpT {
	case 0:
		if cmp&35184372088832 > 0 {
			return 256
		}
		if cmp&70368744177664 > 0 {
			return 732
		}
		if cmp&287948901175001088 > 0 {
			return 773
		}
	case 1:
		if cmp&70368744177664 > 0 {
			return 732
		}
		if cmp&287948901175001088 > 0 {
			return 773
		}
	case 2:
		if cmp&287948901175001088 > 0 {
			return 4373
		}
	case 3:
		if cmp&70368744177664 > 0 {
			return 3093
		}
		if cmp&287948901175001088 > 0 {
			return 3333
		}
	case 12:
		if cmp&287948901175001088 > 0 {
			return 4117
		}
	case 13:
		if cmp&70368744177664 > 0 {
			return 3093
		}
		if cmp&287948901175001088 > 0 {
			return 3333
		}
	case 14:
		if cmp&287948901175001088 > 0 {
			return 5917
		}
		if cmp&35184372088832 > 0 {
			return 6168
		}
	case 16:
		if cmp&287948901175001088 > 0 {
			return 4629
		}
	case 17:
		if cmp&287948901175001088 > 0 {
			return 4885
		}
	case 18:
		if cmp&287948901175001088 > 0 {
			return 4629
		}
	case 19:
		if cmp&287948901175001088 > 0 {
			return 4885
		}
	case 23:
		if cmp&287948901175001088 > 0 {
			return 6429
		}
	case 24:
		if cmp&287948901175001088 > 0 {
			return 5917
		}
	case 25:
		if cmp&287948901175001088 > 0 {
			return 6429
		}
	case 32:
		if cmp&274877906944 > 0 {
			return 1244
		}
		if cmp&4503599627370496 > 0 {
			return 1500
		}
	case 35:
		if cmp&17179869184 > 0 {
			return 1801
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
		if cmp&17592186044416 > 0 {
			return 2577
		}
		if cmp&2251799813685248 > 0 {
			return 2829
		}
		if cmp&32 > 0 {
			return 3804
		}
	case 36:
		if cmp&8589934592 > 0 {
			return 4060
		}
	case 37:
		if cmp&1125899906842624 > 0 {
			return 1756
		}
	case 38:
		if cmp&9007199254740992 > 0 {
			return 5852
		}
	case 39:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 42:
		if cmp&274877906944 > 0 {
			return 2339
		}
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
	case 43:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 44:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 45:
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
		if cmp&17592186044416 > 0 {
			return 2577
		}
		if cmp&2251799813685248 > 0 {
			return 2829
		}
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&17179869184 > 0 {
			return 1801
		}
	case 47:
		if cmp&17592186044416 > 0 {
			return 5340
		}
	case 48:
		if cmp&274877906944 > 0 {
			return 2339
		}
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
	case 49:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 50:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 51:
		if cmp&32 > 0 {
			return 3804
		}
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 52:
		if cmp&2251799813685248 > 0 {
			return 5596
		}
	case 53:
		if cmp&137438953472 > 0 {
			return 6703
		}
	case 54:
		if cmp&137438953472 > 0 {
			return 6955
		}
	case 55:
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	case 57:
		if cmp&68719476736 > 0 {
			return 2087
		}
		if cmp&274877906944 > 0 {
			return 2339
		}
	}
	return 2
}
