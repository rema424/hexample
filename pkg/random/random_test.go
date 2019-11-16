package random

import "testing"

// go test -benchmem -bench .

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(100, Alphanumeric, Symbols)
	}
}

func BenchmarkStringPreOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPreOptimized(100, Alphanumeric, Symbols)
	}
}
