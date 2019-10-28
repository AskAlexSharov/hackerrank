package tree

// This file describing Morris non-recursion traversal algo of binary search tree
// Morris Algo: http://www.learn4master.com/algorithms/morris-traversal-inorder-tree-traversal-without-recursion-and-without-stack

import (
	"context"
	"errors"
	"fmt"
)

func GtNthBiggest(root *Node, n int) (int, error) {
	var result *Node

	ctx, cancel := context.WithCancel(context.Background())
	for node := range root.PostorderMorris(ctx) {
		fmt.Printf("%d\n", node.Value) // if there is no left child, visit current node
		n--
		if n == 0 { // found
			result = &node
			cancel()
		}
	}

	if result == nil {
		return 0, errors.New("not found")
	}

	return result.Value, nil
}

type Visitor = func(node *Node) (goOn bool)

func (tree *Node) LeftRightMostChild() *Node {
	child := tree.Left

	// when pre.right == null, it means we go to the right most leaf
	// when pre.right == cur, it means the right most leaf has been visited in the last round
	for child.Right != nil && child.Right != tree {
		child = child.Right
	}
	return child
}

func (tree *Node) RightLeftMostChild() *Node {
	child := tree.Right
	// when pre.right == null, it means we go to the right most leaf
	// when pre.right == cur, it means the right most leaf has been visited in the last round
	for child.Left != nil && child.Left != tree {
		child = child.Left
	}

	return child
}

// Inorder Traverse without recursion and stack
// Morris algo
//  http://www.learn4master.com/algorithms/morris-traversal-inorder-tree-traversal-without-recursion-and-without-stack
func (tree *Node) InorderMorris(ctx context.Context) chan Node {
	out := make(chan Node, 10)
	go func() {
		defer close(out)
		tree.inorderMorris(ctx, out)
	}()
	return out
}

func (tree *Node) inorderMorris(ctx context.Context, out chan Node) {
	cur := tree
	for cur != nil {
		if cur.Left == nil { // visit leaf node
			select {
			case out <- *cur:
			case <-ctx.Done():
				return
			}

			cur = cur.Right // then we go the right branch
			continue
		}

		// find the right most leaf of root.left node.
		pre := cur.LeftRightMostChild()

		if pre.Right != cur { // the fist time to visit the pre node, make its right child point to current node
			pre.Right = cur // set soft link
			cur = cur.Left
			continue
		}

		// this means the pre.right has been set, it's time to go to current node
		pre.Right = nil

		// means the current node is pointed by left right most child
		// the left branch has been visited, it's time to print the current node
		select {
		case out <- *cur:
		case <-ctx.Done():
			return
		}

		cur = cur.Right // go to parent
	}
}

func (tree *Node) PostorderMorris(ctx context.Context) chan Node {
	out := make(chan Node, 10)
	go func() {
		defer close(out)
		tree.postorderMorris(ctx, out)
	}()
	return out
}

func (tree *Node) postorderMorris(ctx context.Context, out chan Node) {
	cur := tree
	for cur != nil {

		if cur.Right == nil { // visit leaf node
			select {
			case out <- *cur:
			case <-ctx.Done():
				return
			}

			cur = cur.Left // then we go the right branch
			continue
		}

		pre := cur.RightLeftMostChild()

		if pre.Left != cur { // the fist time to visit the pre node, make its right child point to current node
			pre.Left = cur // set soft link
			cur = cur.Right
			continue
		}

		// this means the pre.right has been set, it's time to go to current node
		pre.Left = nil

		// means the current node is pointed by left right most child
		// the left branch has been visited, it's time to print the current node
		select {
		case out <- *cur:
		case <-ctx.Done():
			return
		}

		cur = cur.Left // go to parent
	}
}
