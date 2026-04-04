package substitution

import (
	"testing"
)

func BenchmarkExtractVariablesFromString(b *testing.B) {
	s := "$(params.foo) $(params.bar) $(params.baz)"
	prefix := "params"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExtractVariablesFromString(s, prefix)
	}
}

func BenchmarkExtractEntireVariablesFromString(b *testing.B) {
	s := "$(params.foo) $(params.bar) $(params.baz)"
	prefix := "params"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractEntireVariablesFromString(s, prefix)
	}
}
