package main

import "github.com/wontw/algo/helpers"

func isMatch(s string, p string) bool {
	resp := helpers.Make2d[bool](len(s)+1, len(p)+1)
	resp[0][0] = true

	// fill first row
	for i := 1; i <= len(p); i++ {
		if string(p[i-1]) == "*" {
			resp[0][i] = resp[0][i-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			// two cases:
			// 1) current pattern symbol equals to current string symbol or equals to '.'
			if string(p[j-1]) == string(s[i-1]) || string(p[j-1]) == "." {
				resp[i][j] = resp[i-1][j-1]
				continue
			}

			// 2) current pattern symbol equals to '*'
			// handle it as empty substring or repeated symbol from src string
			if string(p[j-1]) == "*" {
				emptyRepeatMatched := resp[i][j-2]
				var repeatedStringMatched bool
				if string(p[j-2]) == string(s[i-1]) || string(p[j-2]) == "." {
					repeatedStringMatched = true
				}

				resp[i][j] = emptyRepeatMatched || (repeatedStringMatched && resp[i-1][j])
				continue
			}
		}
	}

	return resp[len(s)][len(p)]
}
