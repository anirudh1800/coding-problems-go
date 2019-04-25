package main

var low int
var maxlen int

// Find longest palindrome in a string

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var i int
	for i = 0; i < len(s)-1; i++ {
		expandAroundCenter(s, i, i)
		expandAroundCenter(s, i, i+1)
	}
	return s[low : low+maxlen]
}

func expandAroundCenter(s string, j int, k int) {

	for j >= 0 && k < len(s) && s[j] == s[k] {
		j = j - 1
		k = k + 1
	}

	if maxlen < k-j-1 {
		low = j + 1
		maxlen = k - j - 1
	}

}
