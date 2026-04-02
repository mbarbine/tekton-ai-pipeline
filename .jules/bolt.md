## 2025-01-28 - Regex Compilation Cache Speedup
**Learning:** `regexp.Compile(pattern)` inside hot paths in string substitution logic (e.g. `ExtractVariablesFromString`) caused massive performance overhead. Adding a simple `sync.Map` cache for regular expressions yields a ~4-5x speedup for parsing.
**Action:** When identifying `regexp.Compile` or similar resource-heavy object creations in high-throughput data parsing/substitution functions, introduce a concurrency-safe cache (`sync.Map`) instead of compiling on every request.
