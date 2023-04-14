package rearrangespacesbetweenwords

func reorderSpaces(text string) string {
	spacesCounter := 0
	words := make([]string, 0)
	tempWord := ""
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			spacesCounter++
			if tempWord != "" {
				words = append(words, tempWord)
				tempWord = ""
			}
		} else {
			tempWord += string(text[i])
		}
	}
	if tempWord != "" {
		words = append(words, tempWord)
		tempWord = ""
	}
	qtySpacesBtwWords := 0
	if len(words) > 1 {
		qtySpacesBtwWords = spacesCounter / (len(words) - 1)
	} else {
		qtySpacesBtwWords = spacesCounter
	}
	for i := 0; i < len(words); i++ {
		tempWord += words[i]
		for j := 0; j < qtySpacesBtwWords && spacesCounter != 0; j++ {
			spacesCounter--
			tempWord += " "
		}
	}
	for spacesCounter > 0 {
		spacesCounter--
		tempWord += " "
	}
	return tempWord
}
