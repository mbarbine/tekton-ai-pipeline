## 2026-04-01 - Avoid repeated `regexp.Compile` in hot loops
**Learning:** `regexp.Compile` or `regexp.MustCompile` can be a significant bottleneck if executed continuously within functions like string substitution logic (e.g. `ExtractVariablesFromString`). String interpolation to form patterns and compiling on the fly can be slow.
**Action:** Use a caching layer like `sync.Map` to map pattern prefixes/variables to compiled regular expressions. Cache compiled regexes per unique pattern to reduce repeated compilation overhead, drastically improving execution speed for parameter parsing.
