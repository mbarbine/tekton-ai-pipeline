## 2024-05-20 - Avoid strings.Split for line extraction
**Learning:** `strings.Split` allocates a new slice containing all elements. When only the first element (or the rest of the string) is needed, this causes unnecessary heap allocations proportional to the number of elements (e.g., number of lines in a script).
**Action:** Use `strings.IndexByte` and string slicing to find the first occurrence and extract prefixes or suffixes without allocating slices.
