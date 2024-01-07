# 808. Soup Servings

這題看解答真的超困惑，什麼叫做同時清空的機率是0.5
後來看影片才知道是同時清空時，機率算做一半
「plus half the probability that A and B become empty at the same time」
爆出來之後就是dfs裸題
注意數值範圍 超過5000以上的時候 1 - p <= 1e5
所以m >= 5000 return 1
