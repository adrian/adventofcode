package utils

func Contains(g byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == g {
			return true
		}
	}
	return false
}
