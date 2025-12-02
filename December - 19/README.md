# Day 19: Balanced Team Assignment (Partition Problem)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N*Sum)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Divide a group of employees into two teams such that the absolute difference between the total skill levels of the two teams is minimized.

**Input:**
1.  Integer `N`.
2.  Array `skills` (positive integers).

**Output:** The minimum difference (integer).

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `1 2 3 5` | `1` | Team A: {1, 5} (Sum 6), Team B: {2, 3} (Sum 5). Diff: 1. |

---

## 🧠 The Engineering Approach

### 🚫 The Greedy Trap
Sorting and taking turns picking the largest available number fails.
* *Counter-Example:* `[5, 5, 4, 4, 2]`. Greedy gives diff 2. Optimal gives diff 0.

### ✅ The Optimized Approach (Dynamic Programming)
This is a variation of the **Knapsack Problem** or **Subset Sum Problem**.
We want to find a subset of skills that sums up to $X$, where $X$ is as close to $\text{TotalSum} / 2$ as possible.

**Algorithm:**
1.  Calculate `TotalSum`. Target is `TotalSum / 2`.
2.  Create a boolean DP array `dp` of size `Target + 1`. `dp[i]` is true if sum `i` is achievable.
3.  Initialize `dp[0] = true`.
4.  For every skill `s`:
    * Iterate backwards from `Target` down to `s`.
    * `dp[j] = dp[j] || dp[j - s]`.
5.  Find the largest index `bestSum` in `dp` that is `true`.
6.  Result is `TotalSum - 2 * bestSum`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N \times S)$
    * $N$ is number of employees, $S$ is the total sum of skills. Since $S$ is small (max $30 \times 100 = 3000$), this is extremely fast.
* **Space Complexity:** $O(S)$
    * We need a boolean array of size `TotalSum / 2`.