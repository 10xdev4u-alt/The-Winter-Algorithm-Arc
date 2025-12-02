# Day 29: Minimum Weight Cycle (Graph Theory)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(E^2_log_V)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find the cycle in an undirected, weighted graph that has the minimum total sum of edge weights.

**Input:** Vertices `V` and a list of weighted edges `[u, v, w]`.
**Output:** The weight of the minimum cycle, or `-1` if no cycle exists.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `V=5`, `(0,1,2)...` | `6` | Cycle `0-1-4` has weights `2+1+3 = 6`. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive DFS
DFS can detect cycles, but finding the *shortest* one by weight requires exhaustive backtracking, which is computationally expensive ($O(V!)$ in worst case).

### ✅ Edge Removal + Dijkstra
We rely on the definition of a simple cycle: **A cycle is an edge $(u, v)$ plus a path between $u$ and $v$ that doesn't use that edge.**

To find the smallest cycle involving a specific edge $(u, v)$:
1.  **Ignore** the direct edge between $u$ and $v$.
2.  Find the **Shortest Path** between $u$ and $v$ using the rest of the graph (Dijkstra's Algorithm).
3.  Calculate `CycleCost = Weight(u,v) + DijkstraDist(u,v)`.

We iterate this process for every edge in the graph and keep the minimum `CycleCost` found.

**Optimization:**
If the weight of the current edge is already greater than the best cycle found so far, we skip the Dijkstra step entirely.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(E \times (E \log V))$
    * We run Dijkstra ($E \log V$) for potentially every edge ($E$).
* **Space Complexity:** $O(V + E)$
    * Adjacency list storage.