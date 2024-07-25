package p733

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	fill(image, sc, sr, image[sr][sc], color)
	return image
}

func fill(image [][]int, x, y, oldColor, newColor int) {

	if y < 0 || y >= len(image) {
		return
	}
	if x < 0 || x >= len(image[0]) {
		return
	}

	if image[y][x] == newColor || image[y][x] != oldColor {
		return
	}

	image[y][x] = newColor

	fill(image, x-1, y, oldColor, newColor)
	fill(image, x+1, y, oldColor, newColor)
	fill(image, x, y-1, oldColor, newColor)
	fill(image, x, y+1, oldColor, newColor)

}
