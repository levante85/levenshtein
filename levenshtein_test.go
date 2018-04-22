package levenshtein

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	mat, dist := Distance("GUMBO", "GAMBOL")
	if dist != 2 {
		t.Fatal("Distance should be 2 instead is: ", dist)
	}
	fmt.Println(mat)
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Distance("GUMBO", "GAMBOL")
	}
}

func distanceSlow(s, t string, len_s, len_t int) int {
	var cost int

	if len_s == 0 {
		return len_t
	}
	if len_t == 0 {
		return len_s
	}

	if s[len_s-1] == t[len_t-1] {
		cost = 0
	} else {
		cost = 1
	}

	return minumum(distanceSlow(s, t, len_s-1, len_t)+1,
		distanceSlow(s, t, len_s, len_t-1)+1,
		distanceSlow(s, t, len_s-1, len_t-1)+cost)
}

func BenchmarkDistanceSlow(b *testing.B) {
	s := "gumbo"
	t := "gambol"
	for i := 0; i < b.N; i++ {
		distanceSlow(s, t, len(s), len(t))
	}
}
