package main

import (
	"fmt"
)

/*
	Найти 2 числа в слайсе data, сумма который даёт число t
*/

func findTarget(data []int, t int) (int, int) {

}

func main() {
	fmt.Println(findTarget([]int{1, 4, 2}, 5))
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return `()`
	}
	return fmt.Sprintf(`%d(%s)(%s)`, root.Val, tree2str(root.Left), tree2str(root.Right))
}

func main() {
	fmt.Println(tree2str(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	))
}
