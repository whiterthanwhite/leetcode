package minimumbitflipstoconvertnumber

import "fmt"

func minBitFlips(start, goal int) int {
	if start == goal {
		return 0
	}

	startStr := toBinaryString(start)
	goalStr := toBinaryString(goal)

	startLen := len(startStr)
	goalLen := len(goalStr)
	if startLen > goalLen {
		goalStr = getSymbolString(startLen-goalLen, "0") + goalStr
	} else if goalLen > startLen {
		startStr = getSymbolString(goalLen-startLen, "0") + startStr
		startLen = goalLen
	}

	flips := 0
	for i := 0; i < startLen; i++ {
		if startStr[i] != goalStr[i] {
			flips++
		}
	}

	return flips
}

func toBinaryString(n int) string {
	binaryStr := ""
	for i := n; i > 0; i /= 2 {
		binaryStr = fmt.Sprint(i%2) + binaryStr
	}
	return binaryStr
}

func getSymbolString(size int, sybmol string) string {
	t := ""
	for i := 0; i < size; i++ {
		t += "0"
	}
	return t
}
