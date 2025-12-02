# Day 30: Burn the Binary Tree (Graph Conversion)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Hard-red?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Given a Binary Tree and a Target Node, simulate a fire starting at the target. The fire spreads to the **Parent**, **Left Child**, and **Right Child** simultaneously every second.

**Input:**
1.  Integer `N` (Size of level-order array).
2.  Array of `N` integers (Tree representation, `-1` for null).
3.  Integer `Target`.

**Output:** A list of nodes burned at every second.

| Input | Output |
| :--- | :--- |
| `12 13 10 -1 -1 14 24`<br>Target: `14` | `14`<br>`10`<br>`12, 24`<br>`13` |

---

## 🧠 The Engineering Approach



[Image of Depth First Search Algorithm]


### 🚫 The "One-Way" Trap
A standard Binary Tree is a Directed Graph (Parent $\to$ Child). You cannot traverse "Up" to the parent easily. This makes burning upwards impossible with standard traversal.

### ✅ Tree to Graph Conversion
We treat the Tree as an **Undirected Graph**.
1.  **Parent Map:** Since `TreeNode` usually lacks a `parent` pointer, we traverse the tree once ($O(N)$) and build a Map: `parentMap[Node] -> ParentNode`.
2.  **BFS (Fire Spread):** We treat the Target Node as the center of a ripple.
    * **Queue:** Start with `[Target]`.
    * **Loop:** At every step, we attempt to move to `node.Left`, `node.Right`, and `parentMap[node]`.
    * **Visited:** Essential to prevent the fire from burning a node twice.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We traverse the tree once to map parents ($N$).
    * We burn every node exactly once using BFS ($N$).
* **Space Complexity:** $O(N)$
    * We store the `parentMap` ($N$) and the Queue ($N$).