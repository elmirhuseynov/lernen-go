package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	prefix := longestCommonPrefix(strs)

	fmt.Println(prefix)
}

func longestCommonPrefix(strs []string) string {
	result := strs[0]

	for i := 1; i < len(strs); i++ {
		str := strs[i]

		for strings.Index(str, result) != 0 {
			result = result[0 : len(result)-1]
		}
	}

	return result
}
