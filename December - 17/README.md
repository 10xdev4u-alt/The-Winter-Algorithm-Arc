# Day 17: Racing Turtles (Sorting & Physics)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Medium-yellow?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N_log_N)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Calculate the number of distinct "fleets" of turtles reaching a finish line.
Turtles start at different positions with different speeds. A faster turtle catches up to a slower one and forms a fleet (the fleet moves at the slower turtle's speed). Turtles cannot overtake one another.

**Input:**
1.  Integer `Target` (Finish line).
2.  Integer `N` (Number of turtles).
3.  Array `Position` (Starting locations).
4.  Array `Speed`.

**Output:** Integer count of fleets.

| Input | Output | Explanation |
| :--- | :--- | :--- |
| `Target: 12`<br>`Pos: [10,8,0,5,3]`<br>`Spd: [2,4,1,1,3]` | `3` | Turtles starting at 10 and 8 merge.<br>Turtle at 5 forms its own fleet.<br>Turtles at 0 and 3 merge. |

---

## 🧠 The Engineering Approach

### 🚫 The Simulation Trap
Simulating the race second-by-second is inefficient because the `Target` distance can be very large ($10^6$), leading to Time Limit Exceeded (TLE).

### ✅ The Optimized Approach (Sort + Time Calculation)
We solve this using **Physics** and **Sorting**.
Since turtles cannot overtake, the turtle closest to the finish line dictates the pace for anyone stuck behind them.

**Algorithm:**
1.  **Sort:** We combine position and speed into a struct and sort all turtles by **Position Descending** (closest to the target first).
2.  **Calculate Time:** For each turtle, calculate the time to arrive:
    $$Time = \frac{\text{Target} - \text{Position}}{\text{Speed}}$$
3.  **Linear Scan:** We track the `CurrentFleetTime`.
    * If a turtle behind takes **more** time than the current fleet, it can never catch up. It starts a **New Fleet**.
    * If a turtle behind takes **less or equal** time, it catches up and joins the current fleet (we ignore it).

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N \log N)$
    * The most expensive operation is sorting the turtles by position. The subsequent scan is $O(N)$.
* **Space Complexity:** $O(N)$
    * We allocate a slice of structs to store paired positions and speeds.