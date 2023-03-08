package fizzbuzz

import "fmt"

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 0; i < n; i++ {
		x := i + 1
		if x%3 == 0 && x%5 == 0 {
			answer[i] = "FizzBuzz"
		} else if x%5 == 0 {
			answer[i] = "Buzz"
		} else if x%3 == 0 {
			answer[i] = "Fizz"
		} else {
			answer[i] = fmt.Sprint(x)
		}
	}
	return answer
}
