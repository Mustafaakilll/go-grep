package main

import (
	"os"
	"testing"
)

func BenchmarkCompareLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareLines(Config{
			Filename:   "sherlock.txt",
			Word:       "sherlock",
			IgnoreCase: false,
		}, os.Stdin)
	}
}
