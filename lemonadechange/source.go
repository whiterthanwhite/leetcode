package lemonadechange

func lemonadeChange(bills []int) bool {
	checkout := make(map[int]int)
	for i := 0; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			checkout[5]++
		case 10:
			qty, ok := checkout[5]
			if qty > 0 && ok {
				checkout[5]--
				checkout[10]++
			} else {
				return false
			}
		case 20:
			qty5, ok5 := checkout[5]
			qty10, ok10 := checkout[10]
			if (qty5 > 0 && ok5) && (qty10 > 0 && ok10) {
				checkout[5]--
				checkout[10]--
				checkout[20]++
			} else if qty5 > 2 && ok5 {
				checkout[5] -= 3
				checkout[20]++
			} else {
				return false
			}
		}
	}
	return true
}
