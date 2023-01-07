/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

package main

// @lc code=start
func isIsomorphic(s string, t string) bool {
	mapST, mapTS := map[uint8]uint8{}, map[uint8]uint8{}
	for i := range s {
		cS, cT := s[i], t[i]
		_, okS := mapST[cS]
		_, okT := mapTS[cT]
		if okS && mapST[cS] != cT || okT && mapTS[cT] != cS {
			return false
		}
		mapST[cS] = cT
		mapTS[cT] = cS
	}
	return true
}

// @lc code=end
