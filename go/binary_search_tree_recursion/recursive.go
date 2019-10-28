package tree

import (
	"context"
	"math"
	"math/rand"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (tree *Node) Insert(num int) {
	if num == tree.Value {
		return
	}

	if num > tree.Value {
		if tree.Right == nil {
			tree.Right = &Node{Value: num}
		} else {
			tree.Right.Insert(num)
		}
	}

	if num < tree.Value {
		if tree.Left == nil {
			tree.Left = &Node{Value: num}
		} else {
			tree.Left.Insert(num)
		}
	}
}

func (tree *Node) inorderRecursive(ctx context.Context, out chan Node) {
	if tree == nil {
		return
	}

	if tree.Left != nil {
		tree.Left.inorderRecursive(ctx, out)
	}

	out <- *tree

	if tree.Right != nil {
		tree.Right.inorderRecursive(ctx, out)
	}
}

func (tree *Node) postorderRecursive(ctx context.Context, out chan Node) {
	if tree == nil {
		return
	}

	if tree.Right != nil {
		tree.Right.postorderRecursive(ctx, out)
	}

	select {
	case out <- *tree:
	case <-ctx.Done():
		return
	}

	if tree.Left != nil {
		tree.Left.postorderRecursive(ctx, out)
	}
}

func (tree *Node) preorderRecursive(ctx context.Context, out chan Node) {
	if tree == nil {
		return
	}

	select {
	case out <- *tree:
	case <-ctx.Done():
		return
	}

	if tree.Left != nil {
		tree.Left.preorderRecursive(ctx, out)
	}

	if tree.Right != nil {
		tree.Right.preorderRecursive(ctx, out)
	}
}

func (tree *Node) Inorder(ctx context.Context) chan Node {
	out := make(chan Node, 10)
	go func() {
		defer close(out)
		tree.inorderRecursive(ctx, out)
	}()
	return out
}

func (tree *Node) Postorder(ctx context.Context) chan Node {
	out := make(chan Node, 10)
	go func() {
		defer close(out)
		tree.postorderRecursive(ctx, out)
	}()
	return out
}

func (tree *Node) Preorder(ctx context.Context) chan Node {
	out := make(chan Node, 10)
	go func() {
		defer close(out)
		tree.preorderRecursive(ctx, out)
	}()
	return out
}

// New
func New(seed int64, n int) *Node {
	tree := &Node{
		Value: int(math.Ceil(float64(n) / 2)),
	}

	rand.Seed(seed)

	for i := 1; i < n; i++ {
		rnd := rand.NormFloat64()
		r := math.Ceil(math.Abs(rnd) * float64(n))
		tree.Insert(int(r))
	}

	return tree
}

func Ch2Arr(ch chan Node) []int {
	res := []int{}
	for node := range ch {
		res = append(res, node.Value)
	}
	return res
}
