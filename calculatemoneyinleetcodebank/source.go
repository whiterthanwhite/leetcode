package calculatemoneyinleetcodebank

func totalMoney(n int) int {
	total := 0
	startValue := 1
	for n > 7 {
		total += calculateWeek(startValue, 7)
		startValue++
		n -= 7
	}
	total += calculateWeek(startValue, n)
	return total
}

func calculateWeek(x, n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += x + (i - 1)
	}
	return sum
}
