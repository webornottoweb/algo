package helpers

func Make2d[T any](r, c int) [][]T {
	resp := make([][]T, r)
	respBackground := make([]T, r*c)
	for i := 0; i < r; i++ {
		rowStart := i * c
		rowEnd := rowStart + c
		currentRow := respBackground[rowStart:rowEnd]
		resp[i] = currentRow
	}

	return resp
}
