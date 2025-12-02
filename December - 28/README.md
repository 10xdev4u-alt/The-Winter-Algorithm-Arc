# Day 28: Sweet Rewards for Students (Greedy Two-Pass)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Distribute the minimum total number of sweets to `N` students standing in a line based on their performance scores.
**Rules:**
1.  Every student gets at least 1 sweet.
2.  If a student has a higher score than a neighbor, they must get more sweets than that neighbor.

**Input:**
1.  Integer `N`.
2.  Array of scores.

**Output:** Minimum total sweets required.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `1 0 2` | `5` | Scores: `1 0 2` $\to$ Sweets: `2 1 2`. Sum: 5. |
| `1 2 2` | `4` | Scores: `1 2 2` $\to$ Sweets: `1 2 1`. Sum: 4. |

---

## 🧠 The Engineering Approach

### 🚫 The Recursive Trap
Trying to satisfy both left and right neighbors simultaneously often leads to complex dependency chains or $O(N^2)$ backtracking.

### ✅ The Two-Pass Greedy Strategy
We can satisfy the conditions by breaking the problem into two linear scans ($O(N)$).

1.  **Initialization:** Everyone gets **1 sweet**.
2.  **Left-to-Right Pass:** compare only with the **Left Neighbor**.
    * If `Score[i] > Score[i-1]`, then `Sweets[i] = Sweets[i-1] + 1`.
    * This ensures the "Left Condition" is met.
3.  **Right-to-Left Pass:** Compare only with the **Right Neighbor**.
    * If `Score[i] > Score[i+1]`, we *might* need more sweets.
    * **Crucial Logic:** `Sweets[i] = Max(Sweets[i], Sweets[i+1] + 1)`.
    * We take the `Max` to ensure we don't break the "Left Condition" established in Pass 1.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We iterate the array three times (Init, Forward, Backward). $3N$ simplifies to $O(N)$.
* **Space Complexity:** $O(N)$
    * We store the `sweets` distribution array.