package main

import (
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	output := CalcMd5("abcdef", 609043)
	if !strings.HasPrefix(output, "000001dbbfa") {
		t.Error("Bad Hash", output)
	}
}

func BenchmarkHash(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CalcMd5("abcdef", 609043)
	}
}
