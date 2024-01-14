package main

// Input: poured = 2, query_row = 1, query_glass = 1
// Output: 0.50000
func champagneTower(poured int, query_row int, query_glass int) float64 {
	prevRow := make([]float64, 1)
	prevRow[0] = float64(poured)

	for row := 1; row < query_row+1; row++ {
		currentRow := make([]float64, row+1)
		for i := 0; i < row; i++ {
			extra := prevRow[i] - 1
			if extra > 0 {
				half := extra * 0.5
				currentRow[i] += half
				currentRow[i+1] += half
			}
		}
		prevRow = currentRow
	}
	if prevRow[query_glass] > 1 {
		return 1
	}
	return prevRow[query_glass]
}

// func main() {
// 	println(champagneTower(100000009, 33, 17)) // 1
// }
