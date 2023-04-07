package decodethemessage

func decodeMessage(key string, message string) string {
	keyMap := make(map[byte]byte)
	var c byte = 'a'
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' {
			_, ok := keyMap[key[i]]
			if !ok {
				keyMap[key[i]] = c
				c++
			}
		}
	}

	r := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		v, ok := keyMap[message[i]]
		if ok {
			r[i] = v
		} else {
			r[i] = message[i]
		}
	}

	return string(r)
}
