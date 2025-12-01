# Day 11: Counting Primes (Sieve of Eratosthenes)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N_log_log_N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Count the number of prime numbers strictly less than a non-negative integer `N`.
A prime number is a number greater than 1 that has no divisors other than 1 and itself.

**Input:** Integer `N`.
**Output:** The count of primes $< N$.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `10` | `4` | Primes: 2, 3, 5, 7 |
| `20` | `8` | Primes: 2, 3, 5, 7, 11, 13, 17, 19 |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Trial Division)
Iterate from 2 to $N$. For each number, check if it is divisible by any number up to its square root.
* **Time Complexity:** $O(N \sqrt{N})$.
* **Verdict:** For $N=1,000,000$, this is too slow (~1 Billion operations).

### ✅ The Optimized Approach (Sieve of Eratosthenes)
Instead of testing each number, we generate non-primes (composites) and eliminate them.
1.  **Assume all numbers are prime** (or track them as "not visited").
2.  Start with the first prime, **2**. Mark all its multiples (4, 6, 8...) as composite.
3.  Move to the next unmarked number (**3**). Mark all its multiples (9, 12, 15...) as composite.
4.  Repeat until $\sqrt{N}$.
5.  Count the numbers that remain unmarked.

**Key Optimization:** When marking multiples of `p`, start from `p * p`. Smaller multiples (like `2 * p`) would have already been marked by the smaller prime (`2`).

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N \log (\log N))$
    * This is almost linear, extremely fast.
* **Space Complexity:** $O(N)$
    * We store a boolean flag for every number up to `N`. For $N=10^6$, this consumes ~1MB of memory.