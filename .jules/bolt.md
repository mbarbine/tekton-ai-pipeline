## 2026-04-03 - Avoid map allocation for regex subexp indices
**Learning:** In Go, looking up regex subexpressions by name using a map created from `pattern.SubexpNames()` is much slower and uses more memory than querying the indices once with `SubexpIndex(name)` and directly indexing the match array. The benchmark shows it is nearly an order of magnitude faster and avoids multiple small map allocations.
**Action:** When extracting multiple named subexpressions repeatedly from matches, use `SubexpIndex` on the compiled regexp and index into the match slice directly instead of creating an intermediate dictionary mapping of group names to matched strings.

## 2024-05-18 - Pre-allocating small array versus dynamically sized slices in tight loops
**Learning:** In tight inner loops, such as iterating over a few regex sub-expression indices, defining `[]int{...}` statically evaluates to a slice and creates a heap allocation for the underlying array on every iteration. Moving to a statically sized array `[3]int{...}` defined outside the loop avoids these small allocations entirely, giving measurable and free performance improvements.
**Action:** When you only need to iterate over a few constant or statically known variables inside a loop, use a fixed size array (e.g. `[N]int{...}`) defined outside the loop rather than dynamically sizing a slice.
