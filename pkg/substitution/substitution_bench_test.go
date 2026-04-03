package substitution

import (
	"testing"
)

func BenchmarkExtractVariablesFromString(b *testing.B) {
	input := "echo $(params.foo) $(params.bar) $(params.baz) $(params.qux) $(params.quux)"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExtractVariablesFromString(input, "params")
	}
}

func BenchmarkExtractEntireVariablesFromString(b *testing.B) {
	input := "echo $(params.foo) $(params.bar) $(params.baz) $(params.qux) $(params.quux)"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractEntireVariablesFromString(input, "params")
	}
}
