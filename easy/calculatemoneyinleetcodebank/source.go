package calculatemoneyinleetcodebank

func totalMoney(n int) int {
	total := 0
	N := n / 7
	if N > 0 {
		for i := 1; i <= N; i++ {
			total += 28 + (i-1)*7
		}
		n -= N * 7
	}
	startValue := N + 1
	for i := 1; i <= n; i++ {
		total += startValue + (i - 1)
	}
	return total
}
