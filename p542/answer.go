package p542

func updateMatrix(mat [][]int) [][]int {

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			distanceToNearestZero(mat, i, j, 1)
		}
	}

	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			distanceToNearestZero(mat, i, j, 2)
		}
	}

	return mat
}

const bigValue = 1e8

func distanceToNearestZero(mat [][]int, i, j int, mode int) int {

	if i < 0 || i >= len(mat) {
		return bigValue
	}

	if j < 0 || j >= len(mat[i]) {
		return bigValue
	}

	switch mode {
	case 1: // allow to use 2-direction cached distance

		switch mat[i][j] {
		case 0:
			return 0
		case 1:
			result := min(
				distanceToNearestZero(mat, i, j-1, 3),
				distanceToNearestZero(mat, i-1, j, 3),
			) + 1

			mat[i][j] = result // cache
			return result
		default:
			panic("unexpected value")
		}

	case 2: // allow to use 4-direction cached distance

		switch mat[i][j] {
		case 0:
			return 0 // do nothing for zero-cell
		default:
			// get 4-direction cached distance and plus 1
			result := min(
				distanceToNearestZero(mat, i, j+1, 3),
				distanceToNearestZero(mat, i+1, j, 3),
				distanceToNearestZero(mat, i, j-1, 3),
				distanceToNearestZero(mat, i-1, j, 3),
			) + 1

			mat[i][j] = result
			return result
		}

	case 3:
		// give me cached distance directly
		return mat[i][j]

	default:
		panic("unexpected mode")
	}
}
