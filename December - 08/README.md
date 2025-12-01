# Day 08: Cafeteria Queue Challenge (Queue Simulation)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Determine how many students are unable to eat lunch.
Students stand in a queue and sandwiches are in a stack. If the student at the front likes the sandwich on top (0 for Circle, 1 for Square), they take it and leave. If not, they move to the back of the queue. This continues until no one wants the top sandwich.

**Input:**
1. Number of students `N`.
2. Array `students` (preferences).
3. Array `sandwiches` (stack).

**Output:** Integer count of students remaining.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `[1,1,0,0]`<br>`[0,1,0,1]` | `0` | Everyone finds a match eventually. |
| `[1,1,1,0,0,1]`<br>`[1,0,0,0,1,1]` | `3` | The stack has `0` on top, but all remaining students want `1`. Deadlock. |

---

## 🧠 The Engineering Approach

### 🚫 The Naive Approach (Simulation)
Actually moving elements from the front of a list to the back (`append` + `slice`) to simulate the rotation.
* **Drawback:** Time complexity degrades to $O(N^2)$ because in the worst case, the queue rotates $N$ times for every sandwich.

### ✅ The Optimized Approach (Bucket Counting)
We realize that the **order of students in the queue does not matter**. If a circle sandwich is on top, *any* student who likes circles will eventually rotate to the front and eat it.
The deadlock only happens when the top sandwich is `X`, but **count(students wanting X) == 0**.

**Algorithm:**
1.  **Count Preferences:** Create a frequency map (or array) of how many students want `0` vs `1`.
2.  **Iterate Sandwiches:** Loop through the stack from top to bottom.
    * If `count[current_sandwich] > 0`: Decrement that count (someone eats it).
    * If `count[current_sandwich] == 0`: **Stop**. No one left wants this.
3.  **Result:** Sum of remaining counts.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N)$
    * We iterate through the students once to count, and the sandwiches once to serve.
* **Space Complexity:** $O(1)$
    * We only use two integer variables (or an array of size 2) to store counts.