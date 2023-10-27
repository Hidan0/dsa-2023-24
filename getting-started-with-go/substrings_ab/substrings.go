package substringsab

func NumberOfSubstrings(str *string) int {

	out := 0

	for s, ch := range *str {
		if ch == 'a' {
			for _, ch2 := range (*str)[s:] {
				if ch2 == 'b' {
					out++
				}
			}
		}
	}

	return out
}
