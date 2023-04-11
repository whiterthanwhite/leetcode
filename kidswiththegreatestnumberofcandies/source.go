package kidswiththegreatestnumberofcandies

func kidsWithCandies(candies []int, extraCandies int) []bool {
	N := len(candies)

	maxValue := 0
	for i := 0; i < N; i++ {
		if maxValue < candies[i] {
			maxValue = candies[i]
		}
	}

	result := make([]bool, N)
	for i := 0; i < N; i++ {
		result[i] = candies[i]+extraCandies >= maxValue
	}

	return result
}
