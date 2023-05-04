package bitand2bitcharacters

func isOneBitCharacter(bits []int) bool {
	r := false
	buffer := 0
	for i := 0; i < len(bits); i++ {
		switch bits[i] {
		case 0:
			if buffer != 0 {
				r = false
				buffer = 0
			} else {
				r = true
			}
		case 1:
			if buffer != 0 {
				r = false
				buffer = 0
			} else {
				buffer++
			}
		}
	}
	return r
}
