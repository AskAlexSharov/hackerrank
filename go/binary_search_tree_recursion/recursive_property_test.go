package tree_test

import (
	"context"
	"testing"
	"testing/quick"

	"github.com/askalexsharov/hackerrank/go/tree"
	"github.com/k0kubun/pp"
)

const MaxNodes = 100

func TestDebug(t *testing.T) {
	tr := tree.New(1650, 10)
	pp.Println(tr)
	pp.Println("--- Inorder:", tree.Ch2Arr(tr.Inorder(context.Background())))
	pp.Println("--- Postorder:", tree.Ch2Arr(tr.Postorder(context.Background())))
	pp.Println("--- Preorder:", tree.Ch2Arr(tr.Preorder(context.Background())))
}

func TestInorderProps(t *testing.T) {
	comm := func(n int, seed int64) bool {
		ch := tree.New(seed, n%MaxNodes).Inorder(context.Background())
		return arrIsIncreasing(tree.Ch2Arr(ch))
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

func TestPostorderProps(t *testing.T) {
	comm := func(n int, seed int64) bool {
		ch := tree.New(seed, n%MaxNodes).Postorder(context.Background())
		return arrIsDecreasing(tree.Ch2Arr(ch))
	}

	if err := quick.Check(comm, nil); err != nil {
		t.Error(err)
	}
}

func arrIsDecreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}

	return true
}

func arrIsIncreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}

	return true
}
