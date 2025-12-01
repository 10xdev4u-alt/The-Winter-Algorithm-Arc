# Day 02: Total of Distinct Values

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Given an integer `N`, display numbers from `1` to `N` in four different numerical bases (Decimal, Octal, Hexadecimal, Binary) in a clean, aligned tabular format.

**Input:** `N` (Integer, $1 \le N \le 99$)
**Output:** Four columns aligned dynamically based on the width of the binary representation.

| Input | Output Preview |
| :--- | :--- |
| `2` | ` 1  1  1  1`<br>` 2  2  2 10` |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Fixed Spacing)
Using hardcoded tabs (`\t`) or fixed integer widths (e.g., `%4d`).
* **Issue:** Binary representations grow much faster than decimal.
* **Result:** Columns become misaligned as soon as the binary string exceeds the fixed padding (e.g., `1` takes 1 space, `100` takes 3 spaces).
* **Verdict:** ❌ UI/UX Fail (Visual misalignment).

### ✅ The Optimized Approach (Dynamic Padding)
We determine the "Master Width" before printing. Since Binary (Base-2) is always the longest representation, the width of `binary(N)` dictates the padding for the entire table.

$$Width = \text{length}(\text{binary}(N))$$

**Algorithm:**
1.  **Calculate Width:** Convert `N` to a binary string and measure its length.
    * *Go Trick:* `width := len(fmt.Sprintf("%b", n))`
2.  **Iterate:** Loop `i` from `1` to `N`.
3.  **Dynamic Format:** Use Go's `fmt.Sprintf` with the `%*` verb to inject the calculated width dynamically.
    * `%*d` $\rightarrow$ Decimal
    * `%*o` $\rightarrow$ Octal
    * `%*X` $\rightarrow$ Hex (Uppercase)
    * `%*b` $\rightarrow$ Binary

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * *Why?* We iterate through the numbers exactly once. The base conversion cost is negligible for small integers.
* **Space Complexity:** $O(1)$
    * *Why?* We calculate the string on the fly and print it immediately; no arrays or lists are stored in memory.
