package lflb

func ReadFinityWithCounter[ST Source, FT Finity](source ST, finity FT) (ok bool, count int) {
	status := Status(0)
	count = 0
	defer func() {
		ok = StatusFlag(status)&OK == OK
		if !ok {
			source.Back(count)
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
			count += 1
			status = finity.Feed(char)
			if status&FLAG != 0 {
				return
			}
		}
	}
}

func ReadFinity[ST Source, FT Finity](source ST, finity FT) (ok bool) {
	ok, _ = ReadFinityWithCounter(source, finity)
	return ok
}
