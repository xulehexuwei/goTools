package go_learning

import (
	"testing"
)

func TestArrayAndSlice(t *testing.T) {
	arraySlice()
}

func BenchmarkArray(b *testing.B) {
	for i := 1; i < b.N; i++ {
		arraySlice()
	}

}
