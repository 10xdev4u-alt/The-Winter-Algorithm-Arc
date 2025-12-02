# Day 16: Treasure Trail Adjustment (Linked List Removal)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Remove the $n^{th}$ node from the **end** of a linked list and return the head of the updated list.

**Input:**
1.  Integer `TotalNodes`.
2.  Integer `N` (Target position from end).
3.  Array of values to build the list.

**Output:** The modified linked list.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `5 2`<br>`1 2 3 4 5` | `1 -> 2 -> 3 -> 5 -> nil` | The 2nd node from end is `4`. It is removed. |
| `1 1`<br>`1` | `nil` | The 1st node from end is `1`. List becomes empty. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Two Pass)
1.  Iterate through the list once to count the total length $L$.
2.  Iterate a second time to position $L - n$ to perform the deletion.
* **Drawback:** Requires traversing the list twice ($2N$ operations).

### ✅ The Optimized Approach (Runner Technique)
We use two pointers, `Slow` and `Fast`, to accomplish this in a **single pass**. We maintain a constant gap of size `n` between them.

**Algorithm:**
1.  **Dummy Node:** Create a dummy node pointing to `head`. This simplifies edge cases (e.g., deleting the head itself).
2.  **Create Gap:** Move the `Fast` pointer `n + 1` steps forward.
3.  **The Chase:** Move both `Slow` and `Fast` one step at a time until `Fast` reaches the end (`nil`).
    * At this point, `Slow` is sitting exactly at the node **before** the one we want to delete.
4.  **Surgery:** Update pointers to skip the target: `slow.next = slow.next.next`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We traverse the list exactly once.
* **Space Complexity:** $O(1)$
    * We modify the pointers in-place and only use a few variables (`slow`, `fast`, `dummy`).