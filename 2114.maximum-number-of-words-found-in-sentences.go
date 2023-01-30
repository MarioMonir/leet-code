/*
 * @lc app=leetcode id=2114 lang=golang
 *
 * [2114] Maximum Number of Words Found in Sentences
 */

package main

import (
	"fmt"
	"strings"
)

// @lc code=start
func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		spaces := strings.Count(sentence, " ")
		if max < spaces {
			max = spaces
		}
	}
	return max + 1
}

// @lc code=end

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
	fmt.Println(mostWordsFound([]string{"please wait", "continue to fight", "continue to win"}))
}
