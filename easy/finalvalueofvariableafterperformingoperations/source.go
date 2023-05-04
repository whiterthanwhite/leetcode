package finalvalueofvariableafterperformingoperations

func finalValueAfterOperations(operations []string) int {
	var sum int
	for _, v := range operations {
		if v[0] == '+' || v[1] == '+' {
			sum++
		} else {
			sum--
		}
	}
	return sum
}
