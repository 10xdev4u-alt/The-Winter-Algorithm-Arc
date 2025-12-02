# Day 31: Sudoku Solver (Backtracking)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(9^M)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Fill a 9x9 Sudoku grid such that every row, column, and 3x3 sub-grid contains digits 1-9 exactly once.

**Input:** 9x9 Grid (digits or `.`).
**Output:** Solved 9x9 Grid.

| Input (Partial) | Output (Solved) |
| :--- | :--- |
| `5 3 . . 7 ...` | `5 3 4 6 7 ...` |

---

## 🧠 The Engineering Approach

### ✅ Backtracking (Depth-First Search)
We systematically try numbers in empty cells. If we hit a dead end (no valid number can be placed), we **backtrack** to the previous cell and try a different number.



[Image of Depth First Search Algorithm]


**Algorithm:**
1.  **Find Empty:** Scan grid for the first `.` character.
2.  **Try Candidates:** Loop through digits '1' to '9'.
3.  **Validate:** Check if placing the digit violates Sudoku rules:
    * Row Check
    * Column Check
    * 3x3 Box Check
4.  **Recurse:** If valid, place the digit and recursively call `solve()`.
5.  **Backtrack:** If the recursive call returns `false`, reset the cell to `.` and try the next digit.

### 📊 Complexity Analysis
* **Time Complexity:** $O(9^M)$, where $M$ is the number of empty cells. In the worst case, we try 9 possibilities for every empty slot.
* **Space Complexity:** $O(M)$ for the recursion stack depth.