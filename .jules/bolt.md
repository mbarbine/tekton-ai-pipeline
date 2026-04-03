## 2024-05-18 - Slice nil strictness in Tekton pipeline tests
**Learning:** Tekton pipeline test suite contains strict strict assertions around `nil` slice equality for `ChildStatusReference` and `TaskRun` slice slices. Using `make([]T, 0, len)` unconditionally instead of keeping `var x []T` will cause test failures because it initializes an empty slice rather than `nil`.
**Action:** Always conditionally allocate slice capacities using `if len(items) > 0 { x = make([]T, 0, len(items)) }` when replacing `var x []T` declarations in this codebase to preserve `nil` behaviors.
