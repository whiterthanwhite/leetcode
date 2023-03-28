package slowestkey

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var maxKey byte
	var maxTime int
	for i := range keysPressed {
		t := 0
		if i == 0 {
			t = releaseTimes[i]
		} else {
			t = releaseTimes[i] - releaseTimes[i-1]
		}

		if maxTime < t {
			maxTime = t
			maxKey = keysPressed[i]
		} else if maxTime == t {
			if maxKey < keysPressed[i] {
				maxKey = keysPressed[i]
			}
		}
	}

	return maxKey
}
