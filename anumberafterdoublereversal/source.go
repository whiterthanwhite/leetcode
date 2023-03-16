package anumberafterdoublereversal

func isSameAfterReversals(num int) bool {
	if num > 0 {
		return num%10 != 0
	}
	return true
}
