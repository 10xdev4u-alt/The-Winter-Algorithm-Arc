# Day 23: Shortest Path in a Warehouse (Grid BFS)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(M*N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find the minimum number of steps to travel from the top-left `(0,0)` to the bottom-right `(m-1, n-1)` of a grid.
* `0` = Path
* `1` = Wall
* Movement: Up, Down, Left, Right.

**Input:** Dimensions `M, N` and the Grid.
**Output:** Minimum steps or `-1`.

| Input | Output |
| :--- | :--- |
| `3 3`<br>`0 1 0`<br>`0 1 0`<br>`0 0 0` | `5` |

---

## 🧠 The Engineering Approach

### ✅ Breadth-First Search (BFS)
BFS is the standard algorithm for finding the shortest path in an **unweighted graph** (or grid). It explores neighbors layer by layer.

**Algorithm:**
1.  **Queue:** Initialize with the start node `(0,0)` and `steps = 0`.
2.  **Visited Matrix:** Keep track of visited cells to avoid infinite loops.
3.  **Process:**
    * Pop the current cell `(r, c)`.
    * If `(r, c) == (target)`, return `steps`.
    * Look at all 4 neighbors `(nr, nc)`.
    * If `(nr, nc)` is valid (in bounds, not a wall, not visited), push it to the queue with `steps + 1`.

### 🚫 Why not DFS?
DFS plunges deep into one path before trying others. It might find a path that is 100 steps long before realizing a 5-step path existed right next to the start. DFS does not guarantee the shortest path without checking *every* path.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(M \times N)$
    * In the worst case, we visit every cell in the grid once.
* **Space Complexity:** $O(M \times N)$
    * `visited` array takes $M \times N$. The Queue can hold roughly $O(min(M, N))$ elements (the diagonal).