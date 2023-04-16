package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//k - кол-во общих вершин
//O(n + m - k) (k < n, m)-> O(n + m) - память
//O(k) - память (k - кол-во общих вершин) - если структура данных будет персист.

// Можем улететь в stackoverflow (рекурсия не хвостовая)
func mergeTreesWithOutPersist(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	var v1, v2 int
	var t1Left, t2Left *TreeNode
	var t1Right, t2Right *TreeNode
	if t1 != nil {
		v1 = t1.Val
		t1Left = t1.Left
		t1Right = t1.Right
	}
	if t2 != nil {
		v2 = t2.Val
		t2Left = t2.Left
		t2Right = t2.Right
	}

	root := &TreeNode{
		Val: v1 + v2,
	}

	root.Left = mergeTreesWithOutPersist(t1Left, t2Left)
	root.Right = mergeTreesWithOutPersist(t1Right, t2Right)

	return root
}

func mergeTreesWithPersist(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	f1, f2 := t1 != nil, t2 != nil
	if f1 && f2 {
		root := &TreeNode{
			Val: t2.Val + t1.Val,
		}
		root.Left = mergeTreesWithOutPersist(t1.Left, t2.Left)
		root.Right = mergeTreesWithOutPersist(t1.Right, t2.Right)
		return root
	} else if f1 {
		return t1
	} else if f2 {
		return t2
	} else {
		return nil
	}
}

func main() {

}
