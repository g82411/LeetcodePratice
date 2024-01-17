# 2916. Subarrays Distinct Element Sum of Squares II

### 極難推

* 設nums = X X X X X X k X X X X i 
* dp[i] = 以nums[i]為結尾, 所有subarray 的distincts nums ^ 2
* nums[k] == nums[i] 所以對於k之後到i-1之前 distincts nums部會因為i而變化 
* 並且square[k:i-1] 也應該等於square[k:i] (因為i == k)

1. dp[i] = sum{square[j:i]} for j in [0, i]
2. = sum{square[j:i]} for j in [0, k] + sum{square[j:i]} for j in [k+1, i]
3. = sum{square[j:i-1]} for j in [0, k] + sum{square[j:i]} for j in [k+1, i]
4. = sum{square[j:i-1]} for j in [0, k] + sum{(count[j:i-1] + 1) ^ 2} for j in [k+1, i-1]
5. = sum{square[j:i-1]} for j in [0, k] +sum{count[j:i-1] ^ 2 + 2 * count[j:i-1] + 1} for j in [k+1,  i-1]
6. = sum{square[j:i-1]} for j in [0, k]+sum{square[j:i-1] + 2 * count[j:i-1] + 1} for j in [k+1, i i-1]
7. = sum{square[j:i-1]} for j in [0, k] + sum{square[j:i-1]} for j in [k+1,  i-1]+ sum{ 2 * count[j:i-1] + 1} for j in [k+1,  i-1] + 1
8. = dp[i-1] +  2 * sum{count[j:i-1]} + i - k for j in [k+1, i-1] + 1