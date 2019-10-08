package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

// get N-th biggest element

type Node struct {
	Left             *Node
	Right            *Node
	Value            int
	AmountOfChildren int
}

func GtNthBiggest(root *Node, n int) (int, error) {
	var result *Node

	root.Postorder(func(node *Node) bool {
		fmt.Printf("%d\n", node.Value) // if there is no left child, visit current node
		n--
		if n == 0 {
			result = node
			return true
		}

		return true
	})

	if result == nil {
		return 0, errors.New("not found")
	}

	return result.Value, nil
}

func (tree *Node) Insert(num int) {
	n := tree
	for n != nil {
		if n.Value == num {
			break
		}

		n.AmountOfChildren++
		if n.Value > num {
			if n.Left == nil {
				n.Left = &Node{Value: num}
				break
			}

			n = n.Left
		}

		if n.Right == nil {
			n.Right = &Node{Value: num}
			break
		}

		n = n.Right
	}
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
func (tree *Node) Inorder(visitor Visitor) {
	cur := tree
	for cur != nil {
		if cur.Left == nil { // visit leaf node
			if goOn := visitor(cur); !goOn {
				break
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
		if goOn := visitor(cur); !goOn {
			break
		}

		cur = cur.Right // go to parent
	}
}

func (tree *Node) Postorder(visitor Visitor) {
	cur := tree
	for cur != nil {

		if cur.Right == nil { // visit leaf node
			if goOn := visitor(cur); !goOn {
				break
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
		if goOn := visitor(cur); !goOn {
			break
		}

		cur = cur.Left // go to parent
	}
}

// BuildBinarySearchTree assadfasdf
func BuildBinarySearchTree(n int) *Node {
	tree := &Node{
		Value: int(math.Ceil(float64(n) / 2)),
	}

	rand.Seed(1650)

	for i := 1; i < n; i++ {
		rnd := rand.NormFloat64()
		r := math.Ceil(math.Abs(rnd) * float64(n))
		tree.Insert(int(r))
	}

	return tree
}

func main() {
	tree := BuildBinarySearchTree(10)
	tree.Inorder(func(node *Node) bool {
		fmt.Printf("%d\n", node.Value) // if there is no left child, visit current node
		return true
	})

	// TODO: get N'tt biggest
	// TODO: get N'th smallest
	// TODO: walk Postorder
	// TODO: walk Preorder
	// TODO: tests

	r, err := GtNthBiggest(tree, 3)
	fmt.Printf("Result: %v %v\n", r, err)
}
