# 488. Zuma Game

純暴力破解
少數可以剪枝的地方是放過去不連續
然後cache要注意要做成uint8[5]
不然很容易MLE