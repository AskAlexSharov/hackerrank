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
	walkDepthRight(root, n, result)
	if result == nil {
		return 0, errors.New("not found")
	}

	return result.Value, nil
}

func walkDepthRight(n *Node, searchingFor int, result *Node) {

	node := n
	for node != nil {
		fmt.Printf("Node: %d \n", n.Value)

		if node.AmountOfChildren == searchingFor {
			result = node
			break
		}

		if node.Right != nil && node.Right.AmountOfChildren >= searchingFor {
			node = node.Right
			continue
		}

		if node.Left == nil {
			break
		}

		node = node.Left
	}
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

type Visitor = func(node *Node)

// Inorder Traverse without recursion and stack
// Morris algo
//  http://www.learn4master.com/algorithms/morris-traversal-inorder-tree-traversal-without-recursion-and-without-stack
func (tree *Node) Inorder(visitor Visitor) {
	cur := tree
	for cur != nil {
		if cur.Left == nil {
			visitor(cur)
			cur = cur.Right // then we go the right branch
		} else {
			// find the right most leaf of root.left node.
			pre := cur.Left

			// when pre.right == null, it means we go to the right most leaf
			// when pre.right == cur, it means the right most leaf has been visited in the last round
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			// this means the pre.right has been set, it's time to go to current node
			if pre.Right == cur {
				pre.Right = nil

				// means the current node is pointed by left right most child
				// the left branch has been visited, it's time to print the current node
				visitor(cur)
				cur = cur.Right
			} else {
				// the fist time to visit the pre node, make its right child point to current node
				pre.Right = cur
				//fmt.Printf("----set pre for cur=%v\n", cur.Value)
				cur = cur.Left
			}
		}
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
	tree.Inorder(func(node *Node) {
		fmt.Printf("%d\n", node.Value) // if there is no left child, visit current node
	})

	// TODO: get N'tt biggest
	// TODO: get N'th smallest
	// TODO: walk Postorder
	// TODO: walk Preorder
	// TODO: tests

	//r := GtNthBiggest(tree, 10)
	//fmt.Printf("Result: %v\n", r)
}
