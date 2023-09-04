
# LeetCode 練習專案

就 LeetCode的Golang練習＆吐槽

## 目錄

- [差分數列](#差分數列)
- [DP](#DP)
- [Greedy](#Greedy)
- [雙指針](#雙指針)
- [DFS](#DFS)
- [Graph](#Graph)
- [Tree](#Tree)


---
### Tree
#### 共同祖先
- **[#2846_Minimum_Edge_Weight_Equilibrium_Queries_in_a_Tree](./problems/2846_Minimum_Edge_Weight_Equilibrium_Queries_in_a_Tree/solution.go)** - [LeetCode Link](https://leetcode.com/problems/minimum-edge-weight-equilibrium-queries-in-a-tree/)

### 雙指針
#### 滑窗
- **[#2024_Maximize_the_Confusion_of_an_Exam](./problems/2024_Maximize_the_Confusion_of_an_Exam/solution.go)** - [LeetCode Link](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/) 
- **[#2747_Count_Zero_Request_Servers](./problems/2747_Count_Zero_Request_Servers/solution.go)** - [LeetCode Link](https://leetcode.com/problems/count-zero-request-servers/description/)
- **[#2516_Take_K_of_Each_Character_From_Left_and_Right](./problems/2516_Take_K_of_Each_Character_From_Left_and_Right/solution.go)** - [LeetCode Link](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/)

### 差分數列

- **[#1094_Car_Pooling](./problems/1094_Car_Pooling/solution.go)** - [LeetCode Link](https://leetcode.com/problems/car-pooling/)
- **[#759_Employee_Free_Time](./problems/759_Employee_Free_Time/solution.go)** - [LeetCode Link](https://leetcode.com/problems/employee-free-time/)
- **[#731_My_Calendar_II](./problems/732_My_Calendar_III/solution.go)** - [LeetCode Link](https://leetcode.com/problems/my-calendar-ii/)

### DP
#### 依賴前一個的簡單型
- **[#256_Paint_House](./problems/256_Paint_House)** - [LeetCode Link](https://leetcode.com/problems/paint-house/)
- **[#265_Paint_House_II](./problems/265_Paint_House_II/solution.go)** - [LeetCode Link](https://leetcode.com/problems/paint-house-ii/)

- **[#801 Minimum Swaps To Make Sequences Increasing](./problems/801_Minimum_Swaps_To_Make_Sequences_Increasing/solution.go)** - [LeetCode Link](https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing)
- **[#790 Domino and Tromino Tiling](./problems/790_Domino_and_Tromino_Tiling/solution.go)** - [LeetCode Link](https://leetcode.com/problems/domino-and-tromino-tiling)
- **[#2318 Number of Distinct Roll Sequences](./problems/2318_Number_of_Distinct_Roll_Sequences/solution.go)** - [LeetCode Link](https://leetcode.com/problems/number-of-distinct-roll-sequences/)
- **[#552_Student_Attendance_Record_II](./problems/552_Student_Attendance_Record_II)** - [LeetCode Link](https://leetcode.com/problems/student-attendance-record-ii/)
- **[#790_Domino_and_Tromino_Tiling](./problems/790_Domino_and_Tromino_Tiling)** - [LeetCode Link](https://leetcode.com/problems/domino-and-tromino-tiling/)
- **[#1262_Greatest_Sum_Divisible_by_Three](./problems/1262_Greatest_Sum_Divisible_by_Three)** - [LeetCode Link](https://leetcode.com/problems/greatest-sum-divisible-by-three/)
- **[#1419_Minimum_Number_of_Frogs_Croaking](./problems/1419_Minimum_Number_of_Frogs_Croaking)** - [LeetCode Link](https://leetcode.com/problems/minimum-number-of-frogs-croaking/)
- **[#2318_Number_of_Distinct_Roll_Sequences](./problems/2318_Number_of_Distinct_Roll_Sequences)** - [LeetCode Link](https://leetcode.com/problems/number-of-distinct-roll-sequences/)
- **[#818_Race_Car](./problems/818_race_car/solution.go)** - [LeetCode Link](https://leetcode.com/problems/race-car/)
- **[#2361_Minimum_Costs_Using_the_Train_Line](./problems/2361_Minimum_Costs_Using_the_Train_Line/solution.go)** - [LeetCode_Link](https://leetcode.com/problems/minimum-costs-using-the-train-line/)

### Greedy
- **[#134_Gas_Station](./problems/134_Gas_Statio/solution.go)** - [LeetCode Link](https://leetcode.com/problems/gas-station/)
- **[#881_Boats_to_Save_People](./problems/881_Boats_to_Save_People/solution.go)** - [LeetCode Link](https://leetcode.com/problems/boats-to-save-people)
- **[#2171_Removing_Minimum_Number_of_Magic_Beans](./problems/2171_Removing_Minimum_Number_of_Magic_Beans/solution.go)** - [LeetCode Link](https://leetcode.com/problems/removing-minimum-number-of-magic-beans/)
- **[#1727_Largest_Submatrix_With_Rearrangements](./problems/1727_Largest_Submatrix_With_Rearrangements/soultion.go)** - [LeetCode Link](https://leetcode.com/problems/largest-submatrix-with-rearrangements/)
- **[#1657_Determine_if_Two_Strings_Are_Close](./problems/1657_Determine_if_Two_Strings_Are_Close/solution.go)** - [LeetCode Link](https://leetcode.com/problems/determine-if-two-strings-are-close/)
- **[#1578_Minimum_Time_to_Make_Rope_Colorful](./problems/1578_Minimum_Time_to_Make_Rope_Colorful)** - [LeetCode Link](https://leetcode.com/problems/minimum-time-to-make-rope-colorful/editorial/)

### DFS
- **[#489_Robot_Room_Cleaner](./problems/489_Robot_Room_Cleaner/solution.go)** - [LeetCode Link](https://leetcode.com/problems/robot-room-cleaner/)

### Graph

#### Dijkstra + pq
- **[#1514_Path_with_Maximum_Probability](./problems/1514_Path_with_Maximum_Probability/solution.go)** - [LeetCode Link](https://leetcode.com/problems/path-with-maximum-probability/)

---

## 使用方法

1. 克隆此專案。
2. 在 `problems/` 資料夾中找到你感興趣的問題。
3. 運行測試：

```bash
go test ./problems/問題資料夾名稱/
```

例如：

```bash
go test ./problems/0001_two_sum/
```

