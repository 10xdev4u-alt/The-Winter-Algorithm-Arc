# Day 01: Perfect Squares 

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(sqrt(N))-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Given an integer range `1` to `N`, identify all perfect squares and calculate the total count.

**Input:** `N` (Integer, $1 \le N \le 10^5$)
**Output:** List of squares + Total Count.

| Input | Output Sequence | Count |
| :--- | :--- | :--- |
| `20` | `1 4 9 16` | `4` |
| `5` | `1 4` | `2` |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Brute Force)
Iterating from `1` to `N` and checking `Math.Sqrt(x)` for every number.
*   **Operations:** $N$ checks.
*   **Cost:** Expensive floating-point math on every iteration.
*   **Verdict:** ❌ Too Slow ($O(N)$).

### ✅ The Optimized Approach (Generation)
Instead of *checking* numbers, we **generate** them.
We rely on the mathematical property that perfect squares grow quadratically.
$$Sequence = \{ i^2 \mid i \in \mathbb{Z}^+, i^2 \le N \}$$

**Algorithm:**
1.  Initialize `i = 1`.
2.  Compute `sq = i * i`.
3.  If `sq > N`, **Break**.
4.  Print `sq`, increment count, increment `i`.

---

### 📊 Complexity Analysis
*   **Time Complexity:** $O(\sqrt{N})$
    *   *Why?* We only run the loop $\sqrt{N}$ times. For $N=100$, we loop 10 times.
*   **Space Complexity:** $O(1)$
    *   *Why?* We only store a single counter variable.
