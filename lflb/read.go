package lflb

func ReadFinity[ST Source, FT Finity](source ST, finity FT) (ok bool) {
	status := Status(0)
	c := 0
	defer func() {
		ok = StatusFlag(status)&OK == OK
		if !ok {
			source.Back(c)
		} else {
			back := int(status >> 2)
			if back > 0 {
				source.Back(back)
			}
		}
	}()
	char, isEof := byte(0), false
	for {
		char, isEof = source.ThisThenNext()
		if isEof {
			status = finity.FeedEof()
			return
		} else {
			c += 1
			status = finity.Feed(char)
			if status&FLAG != 0 {
				return
			}
		}
	}
}
