# Day 27: Signal Propagation (Dijkstra's Algorithm)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(E_log_V)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find the minimum time required for a signal sent from a source node `S` to reach **all** other nodes in a network.
The network is a directed graph where edges represent transmission delays.

**Input:**
1.  Integers `N` (Nodes) and `M` (Links).
2.  List of directed edges `u v t` (from, to, time).
3.  Integer `S` (Source).

**Output:** The time when the last node receives the signal, or `-1` if unreachable.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `0->1(2)`, `1->2(1)`, `0->2(4)` | `3` | Shortest path to 2 is via 1 (2+1=3). Max(dist) = 3. |

---

## 🧠 The Engineering Approach

### ✅ Dijkstra's Algorithm
This is the standard solution for **Single-Source Shortest Path** in graphs with non-negative weights.



[Image of Depth First Search Algorithm]


**Algorithm:**
1.  **Priority Queue:** Use a Min-Heap to always process the node with the current shortest travel time.
2.  **Relaxation:** For every neighbor `v` of current node `u`:
    * If `dist[u] + weight < dist[v]`:
        * Update `dist[v]`.
        * Push `v` to the heap.
3.  **Result Aggregation:** Once the heap is empty, iterate through the `dist` array.
    * If any node has `dist == Infinity`, return `-1`.
    * Otherwise, the answer is the **Maximum** value in the `dist` array (the time the last tower gets the signal).

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(E \log V)$
    * We process each edge once, and each priority queue operation takes logarithmic time relative to the number of vertices.
* **Space Complexity:** $O(V + E)$
    * Adjacency list and distance array.