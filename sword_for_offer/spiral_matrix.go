package main

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	row, column := 0, 0
	row_min, row_max, column_min, column_max := 0, len(matrix)-1, 0, len(matrix[0])-1

	for row_min < row_max && column_min < column_max {
		for column < column_max {
			res = append(res, matrix[row][column])
			column = column + 1
		}
		for row < row_max {
			res = append(res, matrix[row][column])
			row = row + 1
		}
		for column > column_min  {
			res = append(res, matrix[row][column])
			column = column - 1
		}
		for row > row_min  {
			res = append(res, matrix[row][column])
			row = row - 1
		}
		row_min, row_max, column_min, column_max = row_min + 1, row_max - 1, column_min +1, column_max - 1
		row, column = row +1, column + 1
	}
	if row_min == row_max && column_min <= column_max{
		res = append(res, matrix[row_min][column_min:column_max+1]...)
	}else if column_min == column_max && row_min <= row_max{
		for i := row_min; i<=row_max; i++{
			res = append(res, matrix[i][column_min])
		}
	}
	return res
}