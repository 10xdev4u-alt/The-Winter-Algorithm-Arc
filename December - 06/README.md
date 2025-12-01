# Day 06: Magic Square Generator (Odd Order)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N^2)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Generate a Magic Square of size $N \times N$ (where $N$ is odd).
A Magic Square is a grid where the sum of every row, column, and diagonal is the same (The Magic Constant).

**Input:** An odd integer `N`.
**Output:** The Magic Constant and the $N \times N$ grid.

| Input | Output |
| :--- | :--- |
| `3` | `Magic number: 15`<br>`8 1 6`<br>`3 5 7`<br>`4 9 2` |

---

## 🧠 The Engineering Approach

### The Siamese Method (De la Loubère)
We don't guess numbers. We follow a specific movement pattern known as the "Siamese Method".

**The Rules:**
1.  **Start:** Place `1` in the middle of the top row.
2.  **Move:** Always move **Diagonally Up-Right** ($\nearrow$).
    * *Wrap Around:* If you go off the top, go to the bottom. If you go off the right, go to the left.
3.  **Collision:** If the spot you want is already filled, move **Down** ($\downarrow$) one step from your current position instead.

**The Math Implementation:**
Instead of using `if/else` for boundaries, we use **Modulo Arithmetic** for cleaner code:
* **Move Up:** `row = (row - 1 + n) % n`
* **Move Right:** `col = (col + 1) % n`

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N^2)$
    * *Why?* We have to fill exactly $N^2$ cells.
* **Space Complexity:** $O(N^2)$
    * *Why?* We need a 2D array to store the grid.