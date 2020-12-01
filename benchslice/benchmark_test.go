package benchslice

import (
	"math/rand"
	"sort"
	"testing"
	"strings"
)

func makeSlices(strLen, size, count int) [][]string{
	res := make([][]string, count)
	for i :=range res {
		res[i] = make([]string, size)
		for j := range res[i] {
			var b strings.Builder
			for i := 0; i < strLen; i++ {
				b.WriteByte('A' + byte(rand.Intn(24)))
			}
			res[i][j] = b.String()
		}
	}
	return res
}

func BenchmarkSortLeghth10Size10(b *testing.B) {
	slices := makeSlices(10,10, b.N)
	for i := 0; i < b.N; i++ {
		s := slices[i]
		sort.Slice(s, func(j, k int) bool {
			return s[j] < s[k]
		})
	}
}