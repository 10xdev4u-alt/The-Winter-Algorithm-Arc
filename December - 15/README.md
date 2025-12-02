# Day 15: Royal Family Seating (Complete Binary Tree)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Determine if a royal family hierarchy (Binary Tree) forms a **Complete Binary Tree**.

A **Complete Binary Tree** is a binary tree in which every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible.

**Input:** Root of a binary tree.
**Output:** Boolean `true` (if complete) or `false` (if incomplete).

| Input (Level Order) | Output | Explanation |
| :--- | :--- | :--- |
| `1 2 3 4 5 6` | `true` | All levels filled, last level filled left-to-right. |
| `1 2 3 4 5 -1 7` | `false` | Node 3 has a Right child (7) but missing Left child. |

---

## 🧠 The Engineering Approach

### ✅ Breadth-First Search (BFS)
We verify the tree level-by-level using a **Queue**.
The defining characteristic of a complete tree is that in a Level-Order Traversal, once you encounter a `null` (empty seat), you should **never** see a valid node again. If you do, it means there is a "gap" in the seating arrangement.



[Image of Complete vs Incomplete Binary Tree]


**Algorithm:**
1.  Initialize a Queue with the root.
2.  Set a flag `seenNull = false`.
3.  **Iterate:**
    * Pop a node.
    * **If Node is NULL:** Set `seenNull = true`. (We have reached the potential end).
    * **If Node is VALID:**
        * If `seenNull` is *already* true, it means we found a node **after** a gap. Return `false`.
        * Add Left and Right children to the queue (even if they are `nil`).

### 🚫 Why not DFS?
Depth-First Search (DFS) dives deep into branches before visiting neighbors. To check for completeness with DFS, you would need to track the index of every node and compare it to the total node count. BFS is much more intuitive because it naturally follows the "filling order" of the tree.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We visit every node exactly once to check its validity.
* **Space Complexity:** $O(N)$
    * We use a Queue to store the nodes of the current level. In a binary tree, the maximum width is approximately $N/2$.