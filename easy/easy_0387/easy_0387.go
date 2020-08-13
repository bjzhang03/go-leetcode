package easy_0387

func firstUniqChar(s string) int {
	result := -1
	if len(s) > 0 {
		countS := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		for i := 0; i < len(s); i++ {
			countS[s[i]]++
		}
		for i := 0; i < len(s); i++ {
			if countS[s[i]] == 1 {
				result = i
				break
			}
		}
	}
	return result
}