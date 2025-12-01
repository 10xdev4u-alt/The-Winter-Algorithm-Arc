# Day 13: Mountain Peaks in a Trail (Array Scanning)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find all "Peak" elements in an array of elevations. A peak is strictly greater than its immediate left and right neighbors.

**Constraint:** The first and last elements cannot be peaks.

**Input:**
1.  Integer `N`.
2.  `N` integers representing heights.

**Output:** Indices of the peaks, or `-1` if none exist.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `1 3 2 4 5 3` | `1 4` | `3` is greater than 1 and 2. `5` is greater than 4 and 3. |
| `1 2 3 4 5` | `-1` | No element is strictly greater than both neighbors. |

---

## 🧠 The Engineering Approach

### ✅ The Linear Scan
Since a peak is defined entirely by its immediate neighbors, we can solve this in a single pass.
We iterate from index `1` to `N-2` (skipping the first and last elements).

**Algorithm:**
1.  Loop `i` from `1` to `len(arr)-2`.
2.  Check Condition: `arr[i] > arr[i-1]` AND `arr[i] > arr[i+1]`.
3.  If true, add `i` to the result list.
4.  If list is empty, print `-1`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We visit every element exactly once.
* **Space Complexity:** $O(P)$
    * We store $P$ indices, where $P$ is the number of peaks found.