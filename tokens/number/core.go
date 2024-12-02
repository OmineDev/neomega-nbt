package number

type NumberMarkFns [12]func()

// `((/ |-#)(((/d/d*)#(/ |((b|B)#)|((s|S)#)|((l|L)#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |((e|E)(/ |-#)(/d/d*)#))(/ |((f|F)#)|((d|D)#)))|(true#)|(false#)`
func NumberFeed(state uint8, b uint8) (nextMarkTermAccept uint16) {
	jmpT := uint16(state) | (uint16(b)>>6)<<6
	cmp := uint64(1) << ((b << 2) >> 2)
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
			return 1557
		}
	case 2:
		if cmp&287948901175001088 > 0 {
			return 5125
		}
		if cmp&70368744177664 > 0 {
			return 6677
		}
	case 5:
		if cmp&70368744177664 > 0 {
			return 476
		}
		if cmp&287948901175001088 > 0 {
			return 517
		}
	case 6:
		if cmp&287948901175001088 > 0 {
			return 3349
		}
	case 8:
		if cmp&35184372088832 > 0 {
			return 3608
		}
		if cmp&287948901175001088 > 0 {
			return 3869
		}
	case 11:
		if cmp&35184372088832 > 0 {
			return 3608
		}
		if cmp&287948901175001088 > 0 {
			return 3869
		}
	case 13:
		if cmp&287948901175001088 > 0 {
			return 3349
		}
	case 14:
		if cmp&287948901175001088 > 0 {
			return 3869
		}
	case 15:
		if cmp&287948901175001088 > 0 {
			return 4125
		}
	case 16:
		if cmp&287948901175001088 > 0 {
			return 4125
		}
	case 20:
		if cmp&287948901175001088 > 0 {
			return 5125
		}
		if cmp&70368744177664 > 0 {
			return 6677
		}
	case 26:
		if cmp&287948901175001088 > 0 {
			return 6933
		}
	case 27:
		if cmp&287948901175001088 > 0 {
			return 8213
		}
	case 32:
		if cmp&287948901175001088 > 0 {
			return 8213
		}
	case 64:
		if cmp&274877906944 > 0 {
			return 988
		}
		if cmp&4503599627370496 > 0 {
			return 1244
		}
	case 66:
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&17179869184 > 0 {
			return 4873
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&4096 > 0 {
			return 5393
		}
		if cmp&2251799813685248 > 0 {
			return 5645
		}
		if cmp&17592186044416 > 0 {
			return 5905
		}
		if cmp&4 > 0 {
			return 6153
		}
		if cmp&524288 > 0 {
			return 6413
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
	case 67:
		if cmp&8589934592 > 0 {
			return 4572
		}
	case 68:
		if cmp&1125899906842624 > 0 {
			return 7388
		}
	case 70:
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
	case 77:
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
	case 79:
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
	case 80:
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
	case 81:
		if cmp&17592186044416 > 0 {
			return 4828
		}
	case 82:
		if cmp&2251799813685248 > 0 {
			return 7644
		}
	case 83:
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
	case 84:
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&4 > 0 {
			return 6153
		}
		if cmp&524288 > 0 {
			return 6413
		}
		if cmp&17592186044416 > 0 {
			return 5905
		}
		if cmp&2251799813685248 > 0 {
			return 5645
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&4096 > 0 {
			return 5393
		}
		if cmp&17179869184 > 0 {
			return 4873
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&64 > 0 {
			return 2339
		}
	case 85:
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
	case 86:
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
	case 87:
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
	case 88:
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
	case 89:
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
	case 90:
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
	case 91:
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
	case 92:
		if cmp&9007199254740992 > 0 {
			return 7900
		}
	case 93:
		if cmp&137438953472 > 0 {
			return 8495
		}
	case 94:
		if cmp&137438953472 > 0 {
			return 7979
		}
	case 96:
		if cmp&68719476736 > 0 {
			return 2599
		}
		if cmp&137438953472 > 0 {
			return 3036
		}
		if cmp&274877906944 > 0 {
			return 3107
		}
		if cmp&16 > 0 {
			return 1831
		}
		if cmp&32 > 0 {
			return 2268
		}
		if cmp&64 > 0 {
			return 2339
		}
	}
	return 2
}
