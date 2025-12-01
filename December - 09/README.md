# Day 09: Sum of Unique Elements (Frequency Map)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Calculate the sum of all elements that appear **exactly once** in an array. If an element appears two or more times, it is excluded from the sum. If no unique elements exist, return 0.

**Input:**
1.  Integer `N` (Array size).
2.  Array of `N` integers.

**Output:** Integer sum of unique elements.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `1 2 2 3 4 4` | `4` | Unique elements are `1` and `3`. Sum: `1+3=4`. |
| `5 5 5 5` | `0` | No element is unique. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Nested Loops)
For every element in the array, loop through the entire array again to check if it appears elsewhere.
* **Time Complexity:** $O(N^2)$.
* **Verdict:** Inefficient for large inputs ($N > 10,000$).

### ✅ The Optimized Approach (Hash Map)
We utilize a **Frequency Map** (Hash Table) to store the count of every number. This allows us to solve the problem in two linear passes.

**Algorithm:**
1.  **Build Frequency Map:** Iterate through the array `arr`. Update the map: `counts[number]++`.
2.  **Filter & Sum:** Iterate through the map. If `counts[key] == 1`, add the `key` to the `FinalSum`.

**Note on Go Maps:** Iterating over a map in Go is random. However, since addition is commutative ($A + B = B + A$), the random order does not affect the final sum.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * *Why?* Pass 1 builds the map ($N$ operations). Pass 2 iterates the map (at most $N$ operations).
* **Space Complexity:** $O(N)$
    * *Why?* In the worst case (all numbers unique), the map stores $N$ entries.