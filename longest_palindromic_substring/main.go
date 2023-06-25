package main

// longestPalindrome
// check palindrome on every element
// get max of found
func longestPalindrome(s string) string {
	maxPalindrome := ""

	for mainIdx, symbol := range s {
		var currentPalindrome, currentDoubled string
		currentPalindrome = getPalindrome(s, int64(mainIdx), int64(mainIdx))
		if mainIdx < len(s)-1 && symbol == rune(s[mainIdx+1]) {
			currentDoubled = getPalindrome(s, int64(mainIdx), int64(mainIdx+1))
			if len(currentDoubled) > len(currentPalindrome) {
				currentPalindrome = currentDoubled
			}
		}

		if len(currentPalindrome) > len(maxPalindrome) {
			maxPalindrome = currentPalindrome
		}
	}

	return maxPalindrome
}

// getPalindrome
// returns palindrome, generated from selected point of interest
func getPalindrome(s string, leftIdx, rightIdx int64) string {
	if leftIdx < 0 || int(rightIdx) >= len(s) {
		return ""
	}

	resp := s[leftIdx : rightIdx+1]
	spread := 1
	for {
		if leftIdx-int64(spread) < 0 || int(rightIdx+int64(spread)) >= len(s) {
			break
		}
		if s[leftIdx-int64(spread)] != s[int(rightIdx+int64(spread))] {
			break
		}
		resp = string(s[leftIdx-int64(spread)]) + resp + string(s[int(rightIdx+int64(spread))])
		spread++
	}

	return resp
}
