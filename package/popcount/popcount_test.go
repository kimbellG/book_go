package popcount

import "testing"

func BenchmarkPCLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(233421)
	}
}

func BenchmarkPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(233421)
	}
}

func BenchmarkPCSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountSwap(233421)
	}
}

func BenchmarkPCReset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountReset(233421)
	}
}
