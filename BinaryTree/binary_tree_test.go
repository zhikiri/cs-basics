package binarytree

import (
	"testing"
)

func getNodes(values ...int) []*Node {
	nodes := make([]*Node, len(values))
	for _, value := range values {
		nodes = append(nodes, NewNode(value, nil, nil))
	}
	return nodes
}

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	if tree.Root != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, tree.Root)
	}
}

func TestInsertRoot(t *testing.T) {
	tree := NewBinaryTree()
	node := NewNode(10, nil, nil)
	tree.Insert(node)

	if tree.Root != node {
		t.Errorf("Expected '%+v' actual '%+v'", node, tree.Root)
	}
}

func TestInsertRootChildren(t *testing.T) {
	tree := NewBinaryTree(getNodes(10, 5, 15)...)

	node := tree.Root
	for _, value := range [3]int{10, 5, 15} {
		if node == nil {
			t.Errorf("Expected '%+v' to not be nil", node)
		}
		if node.Value != value {
			t.Errorf("Expected '%+v' actual '%+v'", value, node.Value)
		}
		if node.Value == 10 {
			node = tree.Root.Left
			continue
		}
		node = tree.Root.Right
	}
}

func TestContainsOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if contains := tree.Contains(2); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}

func TestContainsOnExistingChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	for _, val := range [3]int{2, 4, 8} {
		if contains := tree.Contains(val); contains != true {
			t.Errorf("Expected '%+v' actual '%+v'", contains, true)
		}
	}
}

func TestContainsOnMissingChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	if contains := tree.Contains(5); contains != false {
		t.Errorf("Expected '%+v' actual '%+v'", contains, false)
	}
}

func TestFindOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if node := tree.Find(5); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", node, nil)
	}
}

func TestFindOnMissingChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	if node := tree.Find(5); node != nil {
		t.Errorf("Expected '%+v' actual '%+v'", node, nil)
	}
}

func TestFindOnExistingChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(4, 8, 2)...)
	for _, value := range [3]int{2, 4, 8} {
		if node := tree.Find(value); node == nil {
			t.Errorf("Expected '%+v' to not be nil", node)
		}
	}
}

func TestMinOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if min := tree.Min(); min != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, min)
	}
}

func TestMinOnRootOnly(t *testing.T) {
	node := NewNode(2, nil, nil)
	tree := NewBinaryTree(node)
	if min := tree.Min(); min != node {
		t.Errorf("Expected '%+v' actual '%+v'", node, min)
	}
}

func TestMinOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(20, 6, 8, 43, 12, 53)...)
	if min := tree.Min(); min.Value != 6 {
		t.Errorf("Expected '%+v' actual '%+v'", min.Value, 6)
	}
}

func TestMaxOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if min := tree.Min(); min != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, min)
	}
}

func TestMaxOnRootOnly(t *testing.T) {
	node := NewNode(2, nil, nil)
	tree := NewBinaryTree(node)
	if min := tree.Min(); min != node {
		t.Errorf("Expected '%+v' actual '%+v'", node, min)
	}
}

func TestMaxOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(20, 6, 8, 43, 12, 53)...)
	if max := tree.Max(); max.Value != 53 {
		t.Errorf("Expected '%+v' actual '%+v'", max.Value, 53)
	}
}

func TestFindParentOnEmptyTree(t *testing.T) {
	tree := NewBinaryTree()
	if parent := tree.FindParent(5); parent != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, parent)
	}
}

func TestFindParentOnRoot(t *testing.T) {
	tree := NewBinaryTree(getNodes(5)...)
	if parent := tree.FindParent(5); parent != nil {
		t.Errorf("Expected '%+v' actual '%+v'", nil, parent)
	}
}

// 32
// 1 62
//   54 71
//  34    83
func TestFindParentOnMultipleChild(t *testing.T) {
	tree := NewBinaryTree(getNodes(32, 62, 1, 54, 71, 34, 83)...)
	tests := map[int]int{1: 32, 62: 32, 54: 62, 71: 62, 34: 54, 83: 71}
	for value, expected := range tests {
		if par := tree.FindParent(value); par.Value != expected {
			t.Errorf("Expected '%+v' actual '%+v'", expected, par.Value)
		}
	}
}
