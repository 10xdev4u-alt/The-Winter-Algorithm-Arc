# Day 12: The Missing Positive Element (Array Marking)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Given an unsorted array of size `N` containing numbers `1` to `N`, find the one missing number and the one duplicate number.

**Constraint:** You must do this in $O(N)$ time and $O(1)$ extra space (no Hash Maps, no new arrays).

**Input:** `N` integers.
**Output:** Missing and Duplicate values.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `3 1 2 2 5` | `Missing: 4`<br>`Duplicate: 2` | `2` appears twice, `4` is missing. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approaches
1.  **Sorting:** Sort the array ($O(N \log N)$) and scan for the gap. Fails Time constraint.
2.  **Hash Map:** Count frequencies ($O(N)$ Space). Fails Space constraint.

### ✅ The Optimized Approach (Negative Marking)
Since the values are constrained to `1..N`, we can use the **array indices** as our checklist.
We treat the array like a Hash Map where `Key = Index` and `Value = Sign`.

1.  **Iterate:** For every number `x`, go to index `|x| - 1`.
2.  **Mark:** Flip the value at that index to negative.
    * If the value is *already* negative, we have seen `x` before. **This is the Duplicate.**
3.  **Scan:** Run through the array again. If we find a **positive** value at index `i`, it means `i + 1` never appeared in the input. **This is the Missing Number.**

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * Two passes over the array.
* **Space Complexity:** $O(1)$
    * We modify the input array in-place to store state.