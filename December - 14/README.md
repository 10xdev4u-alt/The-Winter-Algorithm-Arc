# Day 14: Magical Treasure Chest Parser (Recursive Descent)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Parse a serialized string representing a nested treasure chest.
The string contains integers and nested compartments enclosed in brackets `[]`.
* Integers are "Treasures".
* Lists `[...]` are "Compartments".

**Input:** String `s` (e.g., `"[123,[456]]"`).
**Output:** A structured `NestedChest` object hierarchy.

| Input | Output |
| :--- | :--- |
| `"324"` | `324` |
| `"[1,[2,3]]"` | `[1,[2,3]]` |

---

## 🧠 The Engineering Approach

### ✅ Top-Down Recursive Parser
We treat the string as a stream of characters and maintain a global index (`idx`) to track our position.
We define two specific parsing functions:

1.  **`parseList`:**
    * Triggered by `[`.
    * Iterates through the string.
    * If it sees `[`, it recursively calls itself.
    * If it sees a number, it calls `parseInteger`.
    * Stops when it sees `]`.

2.  **`parseInteger`:**
    * Triggered by a digit or `-`.
    * Consumes characters until it hits a non-digit.
    * Returns the integer object.

This separation of concerns makes the code modular and easier to debug than a single giant function.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We traverse the string exactly once.
* **Space Complexity:** $O(D)$
    * $D$ is the depth of nesting (Recursion Stack).