// Is Subsequence
// https://leetcode.com/problems/is-subsequence/?envType=study-plan&id=level-1
package main

import "fmt"

func findSubsequence(c string, seed string) (int, bool) {
	for i, v := range seed {
		if c == string(v) {
			return i, true
		}
	}
	return -1, false
}

func isSubsequence(s string, t string) bool {
	sLength := len(s)
	tLength := len(t)

	if sLength == 0 {
		return true
	}

	if tLength == 0 {
		return false
	}

	var startIndex int
	for _, sChar := range s {
		lastIndex, ok := findSubsequence(string(sChar), t[startIndex:])
		if ok == false {
			return false
		}
		startIndex = startIndex + lastIndex + 1
	}
	return true
}

func main() {
	tc := [][]string{
		{"abc", "ahbgdc"},
		{"axc", "ahbgdc"},
		{"", "abc"},
		{"b", "abc"},
		{"aaaaaa", "bbaaaa"},
	}

	rs := []bool{
		true,
		false,
		true,
		true,
		false,
	}

	for i, v := range tc {
		result := isSubsequence(v[0], v[1])
		if result != rs[i] {
			fmt.Println(i, "wrong", rs[i], result)
		} else {
			fmt.Println(i, "pass")
		}
	}
}
