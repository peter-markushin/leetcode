package subsequence

func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	ls := len(s)
	lt := len(t)

	if ls == 0 {
		return true
	}

	if lt == 0 {
		return false
	}

	for i < ls {
		for s[i] != t[j] {
			j += 1

			if j == lt {
				return false
			}
		}

		i += 1
		j += 1

		if j == lt && i < ls {
			return false
		}
	}

	return true
}
