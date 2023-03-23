package reversevowelsofastring

func reserseVowels(s string) string {
	r := []byte(s)

	i := uint32(0)
	j := uint32(len(s) - 1)
	for j > i {
		var ch1, ch2 bool
		switch r[i] {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
			ch1 = true
		default:
			i++
		}

		switch r[j] {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
			ch2 = true
		default:
			j--
		}

		if ch1 && ch2 {
			t := r[i]
			r[i] = r[j]
			r[j] = t

			i++
			j--
		}
	}

	return string(r)
}
