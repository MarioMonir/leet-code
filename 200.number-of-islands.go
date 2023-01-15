/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	count := 0
	// g: grid , r: row , c: column
	var visitDFS func(g [][]byte, r int, c int)
	visitDFS = func(g [][]byte, r int, c int) {

		if r < 0 || r >= len(g) || c < 0 || c >= len(g[0]) || g[r][c] == '0' {
			return
		}

		g[r][c] = '0'
		visitDFS(g, r-1, c) // Top
		visitDFS(g, r+1, c) // Bottom
		visitDFS(g, r, c-1) // left
		visitDFS(g, r, c+1) // right
	}

	for r, rowEls := range grid {
		for c, el := range rowEls {
			if el == '1' {
				visitDFS(grid, r, c)
				count++
			}
		}
	}

	return count
}

// @lc code=end
