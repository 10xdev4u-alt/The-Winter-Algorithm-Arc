# Day 25: Treasure Hunt in the Locked Maze (BFS + Bitmask)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(M*N*2^K)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find the minimum steps to reach the Treasure `T` from Start `S` in a grid containing Walls `#`, Keys `a-j`, and Locked Doors `A-J`.

**Input:**
1.  Grid dimensions `M, N`.
2.  The Maze Grid.

**Output:** Minimum steps or `-1`.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `S.a`<br>`#A#`<br>`..T` | `6` | 1. Move right to get key `a`.<br>2. Backtrack left.<br>3. Open door `A`.<br>4. Reach `T`. |

---

## 🧠 The Engineering Approach

### 🚫 Why Standard BFS Fails
Standard BFS marks a cell `(r, c)` as visited the moment you step on it. In this problem, you might need to visit a cell, leave to grab a key, and **return** to that same cell later. Standard BFS would block this "backtracking".

### ✅ BFS with State (Bitmasking)
We expand the definition of a "visited state". You haven't truly visited a state unless you have been to **Position (r, c)** holding **Specific Set of Keys**.



[Image of Depth First Search Algorithm]


**State Definition:** `visited[State{row, col, keyMask}]`

**The Bitmask Trick:**
Instead of a `Set<Character>` or `List` to track keys (which is slow/heavy), we use a single **Integer**.
* Key `a` = 1st bit (`0001`)
* Key `b` = 2nd bit (`0010`)

**Algorithm:**
1.  **Queue:** Stores `{row, col, keys_mask, steps}`.
2.  **Visited Map:** Tracks `State{row, col, keys}`.
3.  **Transitions:**
    * If `.` or `S`: Move normally.
    * If `a-j`: Update bitmask (`key | (1 << bit)`), add to queue.
    * If `A-J`: Check bitmask (`key & (1 << bit)`). If allowed, move.
    * If `T`: Return steps.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(M \times N \times 2^K)$
    * $K$ is the number of keys (max 10).
* **Space Complexity:** $O(M \times N \times 2^K)$
    * To store the visited states.