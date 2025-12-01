# Day 03: Bridge Crossing Challenge (Jump Game)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** You are given an array `stones` where each element represents the **maximum** jump length from that position. Determine if it is possible to reach the last index starting from index `0`.

**Input:** An integer array `stones`.
**Output:** `true` if the last index is reachable, otherwise `false`.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `[2, 3, 1, 1, 4]` | `true` | Jump 1 step from 0 to 1, then 3 steps to the end. |
| `[3, 2, 1, 0, 4]` | `false` | You will always arrive at index 3, which has a jump value of 0. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Backtracking)
Trying every possible combination of jumps creates a decision tree.
* **Complexity:** Exponential $O(2^N)$.
* **Verdict:** ❌ Time Limit Exceeded on large inputs.

### ✅ The Optimized Approach (Greedy Strategy)
Instead of finding a specific path, we track the **Reachable Boundary**.
We ask: *"At current position `i`, what is the furthest index I can possibly reach?"*

**Algorithm:**
1.  Initialize `maxReach = 0`.
2.  Iterate through the array.
3.  **Safety Check:** If `i > maxReach`, it means the current stone is unreachable. Return `false`.
4.  **Push Boundary:** Update `maxReach` using the formula:
    $$\text{maxReach} = \max(\text{maxReach}, i + \text{stones}[i])$$
5.  **Victory Condition:** If `maxReach >= lastIndex`, return `true`.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * *Why?* We traverse the array exactly once (Single Pass).
* **Space Complexity:** $O(1)$
    * *Why?* We only store variables for tracking the current index and max reach; no extra arrays are created.
