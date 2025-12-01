# Day 07: Baseball Score Record Keeper (Stack)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Compute the final score of a baseball game based on a list of string operations. The operations involve adding integers, doubling scores, summing previous scores, or invalidating scores.

**Input:** An array of strings `ops` (e.g., `["5", "2", "C", "D", "+"]`).
**Output:** Integer total sum.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `["5","2","C","D","+"]` | `30` | 1. Add 5 $\to$ `[5]`<br>2. Add 2 $\to$ `[5, 2]`<br>3. Cancel (remove 2) $\to$ `[5]`<br>4. Double (5*2) $\to$ `[5, 10]`<br>5. Plus (5+10) $\to$ `[5, 10, 15]`<br>**Sum:** 30 |

---

## 🧠 The Engineering Approach

### 🚫 The Trap: Running Sum
A common mistake is trying to maintain a single `TotalSum` variable and modifying it as you iterate.
* **Why it fails:** The `C` (Cancel) operation requires you to "undo" the last valid score. The `+` operation requires access to the last *two* valid scores. A single integer cannot remember history.

### ✅ The Optimized Approach (Stack)
We use a **Stack** data structure (implemented as a dynamic slice in Go) to maintain the history of valid scores. This allows us to push new scores, pop invalid ones, and peek at previous values easily.

**Algorithm:**
1.  Iterate through the `ops` list.
2.  **Integer:** Parse string to int, Push to stack.
3.  **`+`:** Read top 2 elements (`stack[n-1]`, `stack[n-2]`), sum them, Push result.
4.  **`D`:** Read top element (`stack[n-1]`), double it, Push result.
5.  **`C`:** Pop (remove) the top element (`stack[:n-1]`).
6.  **Result:** Sum all elements remaining in the stack.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * *Why?* We iterate through the operations list exactly once.
* **Space Complexity:** $O(N)$
    * *Why?* In the worst case (no "C" operations), we store every score in the stack.