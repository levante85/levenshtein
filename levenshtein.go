package levenshtein

// Distance calculates the edit distance between two strings
// using the algorithm names after the russian scientist
// Vladimir Levenshtein, it's implemented using the interative
// version of maximasing speed over space consumption
func Distance(s, t string) ([][]int, int) {
	var (
		cost = 0
		n    = len(s)
		m    = len(t)
	)

	switch {
	case n == 0:
		return nil, m
	case m == 0:
		return nil, n
	}

	d := buildMatrix(n, m)

	for i := 1; i <= n; i++ {
		cs := s[i-1]

		for j := 1; j <= m; j++ {
			ct := t[j-1]

			if cs == ct {
				cost = 0
			} else {
				cost = 1
			}

			d[j][i] = minumum(d[j-1][i]+1, d[j][i-1]+1, d[j-1][i-1]+cost)
		}
	}

	return d, d[m][n]
}

// helper function to build a matrix
func buildMatrix(n, m int) [][]int {
	matrix := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		matrix[i] = append(matrix[i], make([]int, n+1)...)
		matrix[i][0] = i
	}

	for i := 0; i <= n; i++ {
		matrix[0][i] = i
	}

	return matrix
}

// helper function to determine the minumum among 3 integers
func minumum(a, b, c int) int {
	switch {
	case a < b && a < c:
		return a
	case b < a && b < c:
		return b
	}

	return c
}
