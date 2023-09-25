package _764_is_Array_a_Preorder_of_Some_Binary_Tree

// 啥米係preorder => R|Left Sub Tree|Right Sub Tree
// SubTree的規則是一RLeft Sub | Left Sub Sub | Right Sub Sub
// 所以is array preorder iff left sub tree found parent
// 然後go連stack 都要自己實作我哭死
type Stack []int

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s Stack) Top() int {
	return s[s.Len()-1]
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() int {
	prev := *s
	top := prev.Top()
	*s = prev[0 : prev.Len()-1]
	return top
}

func isPreorder(nodes [][]int) bool {
	root := nodes[0]
	parents := Stack{}
	parents.Push(root[0])
	for i := 1; i < len(nodes); i++ {
		parent, child := nodes[i][1], nodes[i][0]
		for !parents.IsEmpty() && parents.Top() != parent {
			parents.Pop()
		}
		// 有孤兒 顯然他不是preorder
		if parents.IsEmpty() {
			return false
		}
		parents.Push(child)
	}
	return true
}
