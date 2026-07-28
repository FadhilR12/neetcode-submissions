func isPalindrome(s string) bool {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	s = strings.ReplaceAll(s, " ", "")
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	l := 0
	r := len(s) - 1

	for r >= l {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}