# Day 24: First Non-Repeating Character (Frequency Map)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Find the *first* character in a string that appears exactly once. If no such character exists, report it.

**Input:** A string of lowercase English letters.
**Output:** The character or a "not found" message.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `swiss` | `w` | 's' repeats. 'w' is the first one that implies count=1. |
| `aabbcc` | `Not Found` | Every character appears twice. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Nested Loops)
For every character `s[i]`, loop through `s[j]` to see if it appears again.
* **Complexity:** $O(N^2)$.
* **Verdict:** Too slow for strings length $10^5$.

### ✅ The Optimized Approach (Frequency Map)
We split the problem into two linear passes:
1.  **Data Collection:** Iterate through the string and increment the count of each character in a fixed-size array (`[26]int`).
2.  **Verification:** Iterate through the **string** again (to preserve the "first" order). For each character, check the array. The first character with `count == 1` is our answer.

**Why iterate the string in Pass 2?**
If we iterated the *array* (a..z), we would find the alphabetically first unique character (e.g., 'i' in "swiss"), not the positionally first unique character ('w').

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * Pass 1: $N$ operations. Pass 2: Max $N$ operations. Total $2N$ which is linear.
* **Space Complexity:** $O(1)$
    * The frequency array is fixed at size 26, regardless of input size.