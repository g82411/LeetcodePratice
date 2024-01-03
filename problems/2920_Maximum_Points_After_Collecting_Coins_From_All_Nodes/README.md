# 2920. Maximum Points After Collecting Coins From All Nodes


## 樹的經典處理

1. 給定幾個參數，描樹樹的Edges, 以及每個Node, 然後node有個value etc. 通則都是以下處理
```go
var neighbors = [][]int
for _, e := range edges {
    neighbors[e[0]] = append(neighbors[e[0]], e[1])
    neighbors[e[1]] = append(neighbors[e[1]], e[0])
}
```
2. DFS pattern
   1. DFS宣告
      2. ```var dfs func (cur, parent interface{})```
      3. 回傳有關 ```var dfs func (cur, parent interface{}) interface{}``` 
      4. 紀錄資訊往下傳遞 ```var dfs func (cur, parent, value, interface{}) interface{}```
   2. 記憶化
      3. 基本```var memo = make(map[interface{}]interface{})```



差不多醬子，基本上這題就是DFS Model

