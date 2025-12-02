# Day 22: AquaNet (Multi-Source BFS)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(V+E)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Determine the minimum time required to fill all reservoirs in a city network.
* Water flows from filled reservoirs to connected neighbors in **1 minute**.
* Multiple reservoirs can have water initially.
* If a reservoir remains unreachable, return `-1`.

**Input:**
1.  Integer `V` (Vertices) and `E` (Edges).
2.  List of edges `u v`.
3.  Initial state array (1 = Wet, 0 = Dry).

**Output:** Minimum time (minutes) or `-1`.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `0-1-2-3-4`<br>Start at `0` | `4` | 0->1(1m)->2(2m)->3(3m)->4(4m) |

---

## 🧠 The Engineering Approach

### ✅ Multi-Source Breadth-First Search (BFS)
This problem asks for the shortest propagation time from *any* source. This is a classic application of BFS.

**Why Multi-Source?**
Instead of running BFS multiple times (once for each starting node), we initialize the Queue with **all** starting nodes at once ($t=0$). This allows the search to naturally find the "meeting points" between different water sources efficiently.

**Algorithm:**
1.  **Graph Setup:** Convert edge list to Adjacency List.
2.  **Initialization:** Push all nodes with `water[i] == 1` into the Queue. Set `time[i] = 0` and mark as visited.
3.  **Process:**
    * While Queue is not empty, pop `node`.
    * For every unvisited `neighbor`, set `time[neighbor] = time[node] + 1`, mark visited, and push to queue.
    * Track the maximum time encountered.
4.  **Verification:** After the queue empties, check if all nodes were visited. If not, return `-1`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(V + E)$
    * We visit every node and traverse every edge exactly once.
* **Space Complexity:** $O(V + E)$
    * Adjacency list stores edges ($E$), Queue stores nodes ($V$).