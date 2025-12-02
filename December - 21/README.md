# Day 21: Efficient Parcel Sorting (Simulation)

![Language](https://img.shields.io/badge/Language-Go_1.23-00ADD8?style=flat-square&logo=go)
![Difficulty](https://img.shields.io/badge/Difficulty-Easy-green?style=flat-square)
![Complexity](https://img.shields.io/badge/Time-O(N^2)-orange?style=flat-square)

## 📜 Problem Statement
**Goal:** Sort a queue of parcels by weight using a machine that can:
1.  **Rotate:** Move elements to bring a specific parcel to the front.
2.  **Pick:** Remove the lightest parcel from the queue.

**Input:** Array of parcel weights.
**Output:** Sorted array.

| Input | Output |
| :--- | :--- |
| `4 2 1 5 3` | `1 2 3 4 5` |

---

## 🧠 The Engineering Approach

### ✅ Selection Sort Simulation
The problem describes a mechanical version of **Selection Sort**.
* In standard Selection Sort, we swap the minimum element with the current index.
* In this "Machine" version, we find the minimum element, "rotate" the queue to bring it to the front (effectively removing it from the middle), and place it in the output bin.

**Algorithm:**
1.  **Find Min:** Loop through the slice to find the index of the smallest item.
2.  **Extract:** Append that item to the `SortedList`.
3.  **Remove:** Delete that item from the `InputList` (simulating the machine picking it up).
    * Go Trick: `slice = append(slice[:i], slice[i+1:]...)`
4.  **Repeat:** Continue until the input list is empty.

---

### 📊 Complexity Analysis
* **Time Complexity:** $O(N^2)$
    * We perform $N$ iterations. In each iteration, we scan for the minimum ($O(k)$) and shift elements to remove ($O(k)$).
* **Space Complexity:** $O(N)$
    * We create a new slice to store the sorted result.