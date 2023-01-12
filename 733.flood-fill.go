/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */
// @lc code=start
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if sr < len(image) && sc < len(image[0]) && image[sr][sc] != color {
		// Get Current Pixel Color
		cur := image[sr][sc]

		// Color The Current pixel
		image[sr][sc] = color

		// Top direction
		if sr-1 >= 0 && cur == image[sr-1][sc] {
			floodFill(image, sr-1, sc, color)
		}
		// Down Direction
		if sr+1 < len(image) && cur == image[sr+1][sc] {
			floodFill(image, sr+1, sc, color)
		}
		// Left direction
		if sc-1 >= 0 && cur == image[sr][sc-1] {
			floodFill(image, sr, sc-1, color)
		}
		// right
		if sc+1 < len(image[0]) && cur == image[sr][sc+1] {
			floodFill(image, sr, sc+1, color)
		}
	}
	return image
}

// @lc code=end
