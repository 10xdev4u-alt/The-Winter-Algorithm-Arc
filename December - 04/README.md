# Day 04: Target Subarray Finder (Prefix Sum)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Given an array of integers (which may include negative numbers) and a target sum `K`, find the starting and ending indices (0-based) of a contiguous subarray that sums up to `K`.

**Input:**
1. Integers `N` (size) and `Target`.
2. Array of `N` integers.

**Output:** The start and end indices `[start end]`, or `[-1 -1]` if not found.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `7 15`<br>`1 2 3 7 5 1 2` | `[2 4]` | Subarray `[3, 7, 5]` sums to 15. |
| `5 100`<br>`10 20 30 40 50` | `[-1 -1]` | No subarray sums to 100. |

---

## 🧠 The Engineering Approach

### 🚫 The Trap: Sliding Window
A common mistake is trying to use the "Sliding Window" (Two Pointer) technique.
* **Why it fails:** The input allows **negative numbers**.
* **The Issue:** Sliding Window assumes that adding an element *increases* the sum and removing one *decreases* it. With negative numbers, this monotonic property breaks, making Sliding Window invalid.

### ✅ The Optimized Approach: Prefix Sum + Hash Map
We rely on basic algebra. If the sum of elements from index `0` to `i` is `CurrSum`, and we want a subarray ending at `i` with sum `Target`, we need to find an earlier prefix sum `PrevSum` such that:

$$CurrSum - PrevSum = Target$$

Rearranging for the lookup:
$$PrevSum = CurrSum - Target$$



**Algorithm:**
1.  Initialize a Map: `SumMap = {0: -1}` (Handles subarrays starting at index 0).
2.  Iterate through the array, calculating `CurrSum`.
3.  Check if `(CurrSum - Target)` exists in the map.
    * **If Yes:** We found the subarray! It starts at `Map[Value] + 1` and ends at `current index`.
    * **If No:** Add `CurrSum` to the map (only if it's not already there, to keep the shortest/earliest index).

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * *Why?* We traverse the array exactly once. Map lookups and insertions are average $O(1)$.
* **Space Complexity:** $O(N)$
    * *Why?* In the worst case (no solution found), we store `N` distinct prefix sums in the Hash Map.