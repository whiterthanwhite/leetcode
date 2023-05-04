package slowestkey

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxKey := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		t := releaseTimes[i] - releaseTimes[i-1]

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
