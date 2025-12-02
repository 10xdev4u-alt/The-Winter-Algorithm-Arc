# Day 20: Tower Visibility Challenge (Monotonic Stack)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** For each tower in a row, find the height of the **first** tower to its **right** that is strictly taller. If no such tower exists, return -1.

**Input:**
1.  Integer `N` (number of towers).
2.  Array `heights` (tower heights).

**Output:** Array of next greater heights.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `4 5 2 25` | `5 25 25 -1` | 4 sees 5. 5 sees 25. 2 sees 25. 25 sees nothing. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Nested Loops)
For every element, scan all elements to its right to find the first greater one.
* **Complexity:** $O(N^2)$.
* **Verdict:** Fails for large inputs ($N=10^5$).

### ✅ The Optimized Approach (Monotonic Stack)
We scan the array from **Right to Left**. We maintain a Stack that stores potential "Next Greater" candidates. The stack is kept monotonic (sorted), meaning we only keep useful candidates.

**Key Insight:**
If we have a stack `[25, 2]` and we encounter a tower `5` (moving left):
* The `2` is useless because `5` is taller and stands in front of it. Any tower to the left of `5` will be blocked by `5` before it can see `2`.
* Therefore, we **pop** all elements $\le$ current height.
* The element remaining on top of the stack is the answer.

**Algorithm:**
1.  Iterate `i` from `N-1` down to `0`.
2.  **Pop:** Remove stack top while `top <= heights[i]`.
3.  **Peek:** If stack empty, answer is `-1`. Else, answer is `top`.
4.  **Push:** Add `heights[i]` to stack.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We traverse the array exactly once. Each element is pushed and popped at most once.
* **Space Complexity:** $O(N)$
    * In the worst case (strictly decreasing array), the stack stores all elements.