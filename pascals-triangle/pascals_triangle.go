package pascal

// Triangle return the pascal trinagule array for r rows.
func Triangle(r int) [][]int {
	res := [][]int{{1}}

	for i := 1; i < r; i++ {
		res = append(res, nextRow(res[i-1]))
	}
	return res
}

// nextRow tell us the next pascal triangule row
func nextRow(prevRow []int) []int {
	newRow := []int{}
	newRow = append(newRow, 1)
	if len(prevRow) > 1 {
		for i := 0; i < len(prevRow)-1; i++ {
			newRow = append(newRow, prevRow[i]+prevRow[i+1])
		}
	}
	newRow = append(newRow, 1)
	return newRow
}
