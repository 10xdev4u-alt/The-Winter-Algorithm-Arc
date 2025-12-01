# Day 05: Island Counter (Graph Traversal)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(R*C)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** You are given a map represented as a 2D grid where `1` represents land and `0` represents water. Your task is to count the number of islands. An island is formed by connecting adjacent lands horizontally or vertically.

**Input:**
1. Dimensions `R` (rows) and `C` (columns).
2. A grid of `0`s and `1`s.

**Output:** A single integer representing the total number of islands.

| Input | Output | Visual Explanation |
| :--- | :--- | :--- |
| `3 3`<br>`1 0 1`<br>`0 1 0`<br>`1 0 1` | `5` | Five separate 1s. No connections exist. |
| `4 5`<br>`1 1 0 0 0`<br>`1 1 0 0 1`<br>`0 0 0 1 1`<br>`0 0 0 0 0` | `2` | One large chunk on the top-left, one block on the right. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (External State)
A common first thought is to use a separate `visited[][]` boolean matrix to keep track of where we have been.
* **Drawback:** This increases space complexity significantly ($O(R \times C)$ extra memory).
* **Complexity:** Managing two separate grids (data grid + visited grid) increases code verbosity and error potential.

### ✅ The Optimized Approach (DFS "Sinking")
We treat the grid as an undirected graph and use **Flood Fill**. To save space, we modify the input grid in-place to track our state.

**The "Sinking" Strategy:**
When we find a `1` (Land), we increment our counter and immediately trigger a **DFS (Depth First Search)** to "sink" the island.
1.  **Visit:** Mark the current cell as `0` (Water).
2.  **Explore:** Recursively call DFS on all 4 neighbors (Up, Down, Left, Right).
3.  **Result:** By the time the recursion finishes, the *entire* connected island has been turned to water. The next time our loop encounters these cells, it sees `0` and skips them.



**Algorithm:**
```go
func dfs(r, c) {
    // 1. Boundary & Water Check
    if outOfBounds(r,c) || grid[r][c] == 0 { return }
    
    // 2. Sink the land
    grid[r][c] = 0 
    
    // 3. Visit Neighbors
    dfs(r+1, c); dfs(r-1, c); ... 
}
````

-----

### 📊 Complexity Analysis

  * **Time Complexity:** $O(R \times C)$
      * *Why?* We visit every cell at most twice: once during the main loop iteration, and once during the DFS recursion.
  * **Space Complexity:** $O(R \times C)$ (Worst Case)
      * *Why?* In the worst case (the entire grid is one long snake-like island), the recursion stack depth could equal the total number of cells.
