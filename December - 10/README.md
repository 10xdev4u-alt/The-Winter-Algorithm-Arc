# Day 10: Zig-Zag Linked List (Reorder List)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Reorder a singly linked list $L_0 \rightarrow L_1 \rightarrow \dots \rightarrow L_{n-1} \rightarrow L_n$ into a zig-zag pattern: $L_0 \rightarrow L_n \rightarrow L_1 \rightarrow L_{n-1} \rightarrow L_2 \dots$

**Input:**
1.  Integer `N`.
2.  `N` node values.

**Output:** The reordered list values.

| Input | Output |
| :--- | :--- |
| `1 2 3 4` | `1 4 2 3` |
| `1 2 3 4 5` | `1 5 2 4 3` |

---

## 🧠 The Engineering Approach



### 🚫 The Naive Approach (Array Dump)
Store all nodes in an array, access them by index to create the new order, and rebuild the list.
* **Drawback:** Requires $O(N)$ extra space to store the nodes.

### ✅ The Optimized Approach (Split, Flip, Zip)
We perform the operation in-place using three linear steps:

1.  **Split:** Use the **Tortoise and Hare** (Slow/Fast pointers) strategy to find the middle of the list. Split the list into two halves.
2.  **Flip:** Reverse the second half of the list ($O(N)$ time, $O(1)$ space).
3.  **Zip:** Merge the two lists together, alternating nodes from the first half and the reversed second half.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * Finding middle + Reversing + Merging are all linear operations.
* **Space Complexity:** $O(1)$
    * We modify the pointers in place without allocating new data structures.