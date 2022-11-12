// 205. Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/?envType=study-plan&id=level-1
package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	isomorphicMap := make(map[string]string)
	existsMapKeys := make(map[string]bool)

	for sIndex, sCharNum := range s {
		sCharStr := string(sCharNum)
		tCharStr := string(t[sIndex])

		existsValue, ok := isomorphicMap[tCharStr]

		if ok {
			if existsValue != sCharStr {
				return false
			}
		} else {
			_, alreadyUsedKey := existsMapKeys[sCharStr]
			if alreadyUsedKey {
				return false
			}
			isomorphicMap[tCharStr] = sCharStr
			existsMapKeys[sCharStr] = true
		}
	}
	return true
}

func main() {
	tc := [][]string{
		{"egg", "add"},
		{"foo", "bar"},
		{"paper", "title"},
		{"badc", "baba"},
	}

	rs := []bool{
		true,
		false,
		true,
		false,
	}

	for i, v := range tc {
		result := isIsomorphic(v[0], v[1])
		if result != rs[i] {
			fmt.Println(i, "wrong", rs[i], result)
		} else {
			fmt.Println(i, "pass")
		}
	}
}
