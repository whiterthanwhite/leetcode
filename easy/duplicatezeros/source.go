package duplicatezeros

func duplicateZeros(arr []int) {
	N := len(arr)
	newArr := make([]int, 0)
	for i := 0; i < N; i++ {
		if arr[i] == 0 {
			newArr = append(newArr, 0, 0)
		} else {
			newArr = append(newArr, arr[i])
		}
	}
	copy(arr, newArr)
}
