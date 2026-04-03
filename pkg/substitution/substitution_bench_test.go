package substitution

import (
	"testing"
)

func BenchmarkExtractVariablesFromString(b *testing.B) {
	input := "echo $(params.foo) $(params.bar) $(params.baz)"
	prefix := "params"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExtractVariablesFromString(input, prefix)
	}
}
