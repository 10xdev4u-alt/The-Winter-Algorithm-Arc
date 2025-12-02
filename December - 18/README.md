# Day 18: Mirror Necklace Check (Two Pointers)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Determine if a sequence of numbers (necklace beads) forms a palindrome.
A sequence is a palindrome if it reads the same forwards and backwards.

**Input:**
1.  Integer `N`.
2.  Array of `N` integers.

**Output:**
* "The necklace is mirrored."
* "The necklace is not mirrored."
* "The necklace is empty." (if N=0)

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `5`, `1 2 3 2 1` | `Mirrored` | Forward: 1-2-3-2-1. Backward: 1-2-3-2-1. Match. |
| `4`, `10 20 20 30` | `Not Mirrored` | Forward: 10... Backward: 30... Mismatch. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Reverse & Compare)
Create a copy of the array, reverse it, and compare it to the original.
* **Drawback:** Requires $O(N)$ extra space to store the copy.

### ✅ The Optimized Approach (Two Pointers)
We iterate from both ends simultaneously towards the center.
1.  Initialize `Left = 0` and `Right = N-1`.
2.  Compare `Arr[Left]` vs `Arr[Right]`.
    * If they differ, return `false` immediately.
    * If equal, increment `Left` and decrement `Right`.
3.  Stop when `Left >= Right`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We visit each element at most once (technically $N/2$ comparisons).
* **Space Complexity:** $O(1)$
    * No extra arrays are created; we simply move indices.