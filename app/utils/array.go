package utils

func Create2DIntArray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
func Create2DBoolArray(m, n int) [][]bool {
	res := make([][]bool, m)
	for i := range res {
		res[i] = make([]bool, n)
	}
	return res
}
