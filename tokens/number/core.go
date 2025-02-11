package number

type NumberMarkFns [12]func()

// `((/ |-#)(((/d/d*)#(/ |((b|B)#)|((s|S)#)|((l|L)#)))|((/d/d*./d*)|(/d*./d/d*))#)(/ |((e|E)(/ |-#)(/d/d*)#))(/ |((f|F)#)|((d|D)#)))|(true#)|(false#)`
// explain:
// (/ |-#) optional sign
// (/d/d*)# base number (interger or number before . in float/double)
// (/ |((b|B)#)|((s|S)#)|((l|L)#)) optional type mark, b,s,l, for int32, / |
// ((/d/d*./d*)|(/d*./d/d*))#) float which in .123 or 123. formate
// (e|E) exp mark
// (/ |((f|F)#)|((d|D)#))) type mark for flot32, float64
// cast (true#)|(false#) to 1/0

func NumberFeed(state uint8, b uint8) (nextMarkTermAccept uint16) {
	jmpT := uint16(state) | (uint16(b)>>6)<<6
	cmp := uint64(1) << ((b << 2) >> 2)
	switch jmpT {
	case 0:
		if cmp&287948901175001088 > 0 {
			return 261
		}
		if cmp&35184372088832 > 0 {
			return 1536
		}
		if cmp&70368744177664 > 0 {
			return 2012
		}
	case 1:
		if cmp&287948901175001088 > 0 {
			return 5125
		}
		if cmp&70368744177664 > 0 {
			return 7189
		}
	case 6:
		if cmp&70368744177664 > 0 {
			return 2012
		}
		if cmp&287948901175001088 > 0 {
			return 261
		}
	case 7:
		if cmp&287948901175001088 > 0 {
			return 9237
		}
	case 16:
		if cmp&35184372088832 > 0 {
			return 8216
		}
		if cmp&287948901175001088 > 0 {
			return 8477
		}
	case 17:
		if cmp&35184372088832 > 0 {
			return 8216
		}
		if cmp&287948901175001088 > 0 {
			return 8477
		}
	case 20:
		if cmp&70368744177664 > 0 {
			return 7189
		}
		if cmp&287948901175001088 > 0 {
			return 5125
		}
	case 28:
		if cmp&287948901175001088 > 0 {
			return 7701
		}
	case 30:
		if cmp&287948901175001088 > 0 {
			return 7957
		}
	case 31:
		if cmp&287948901175001088 > 0 {
			return 7957
		}
	case 32:
		if cmp&287948901175001088 > 0 {
			return 8477
		}
	case 33:
		if cmp&287948901175001088 > 0 {
			return 10781
		}
	case 36:
		if cmp&287948901175001088 > 0 {
			return 10517
		}
	case 41:
		if cmp&287948901175001088 > 0 {
			return 10517
		}
	case 42:
		if cmp&287948901175001088 > 0 {
			return 10781
		}
	case 64:
		if cmp&64 > 0 {
			return 732
		}
		if cmp&1048576 > 0 {
			return 988
		}
		if cmp&274877906944 > 0 {
			return 1244
		}
		if cmp&4503599627370496 > 0 {
			return 1500
		}
	case 65:
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&524288 > 0 {
			return 4621
		}
		if cmp&2251799813685248 > 0 {
			return 4877
		}
		if cmp&4096 > 0 {
			return 5393
		}
		if cmp&17179869184 > 0 {
			return 5641
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&17592186044416 > 0 {
			return 6929
		}
		if cmp&4 > 0 {
			return 7433
		}
	case 66:
		if cmp&2 > 0 {
			return 2780
		}
		if cmp&8589934592 > 0 {
			return 3036
		}
	case 67:
		if cmp&262144 > 0 {
			return 2268
		}
		if cmp&1125899906842624 > 0 {
			return 2524
		}
	case 68:
		if cmp&2 > 0 {
			return 2780
		}
		if cmp&8589934592 > 0 {
			return 3036
		}
	case 69:
		if cmp&262144 > 0 {
			return 2268
		}
		if cmp&1125899906842624 > 0 {
			return 2524
		}
	case 72:
		if cmp&2097152 > 0 {
			return 3292
		}
		if cmp&9007199254740992 > 0 {
			return 3548
		}
	case 73:
		if cmp&2097152 > 0 {
			return 3292
		}
		if cmp&9007199254740992 > 0 {
			return 3548
		}
	case 74:
		if cmp&4096 > 0 {
			return 9180
		}
		if cmp&17592186044416 > 0 {
			return 8924
		}
	case 75:
		if cmp&17592186044416 > 0 {
			return 8924
		}
		if cmp&4096 > 0 {
			return 9180
		}
	case 76:
		if cmp&32 > 0 {
			return 3627
		}
		if cmp&137438953472 > 0 {
			return 3883
		}
	case 77:
		if cmp&32 > 0 {
			return 3627
		}
		if cmp&137438953472 > 0 {
			return 3883
		}
	case 82:
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
	case 83:
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
	case 84:
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&2251799813685248 > 0 {
			return 4877
		}
		if cmp&4 > 0 {
			return 7433
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&524288 > 0 {
			return 4621
		}
		if cmp&17179869184 > 0 {
			return 5641
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&4096 > 0 {
			return 5393
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&17592186044416 > 0 {
			return 6929
		}
	case 85:
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
	case 86:
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
	case 91:
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
	case 92:
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
	case 93:
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
	case 94:
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
	case 95:
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
	case 97:
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
	case 98:
		if cmp&524288 > 0 {
			return 9692
		}
		if cmp&2251799813685248 > 0 {
			return 9948
		}
	case 99:
		if cmp&524288 > 0 {
			return 9692
		}
		if cmp&2251799813685248 > 0 {
			return 9948
		}
	case 100:
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
	case 101:
		if cmp&32 > 0 {
			return 10031
		}
		if cmp&137438953472 > 0 {
			return 10287
		}
	case 102:
		if cmp&32 > 0 {
			return 10031
		}
		if cmp&137438953472 > 0 {
			return 10287
		}
	case 105:
		if cmp&16 > 0 {
			return 5927
		}
		if cmp&32 > 0 {
			return 4316
		}
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&137438953472 > 0 {
			return 4572
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
	case 106:
		if cmp&64 > 0 {
			return 6179
		}
		if cmp&68719476736 > 0 {
			return 6439
		}
		if cmp&274877906944 > 0 {
			return 6691
		}
		if cmp&16 > 0 {
			return 5927
		}
	}
	return 2
}
