package treemap

import (
	"github.com/zytekaron/structs"
	"golang.org/x/exp/constraints"
)

// TreeMap is an implementation of a red-black tree.
//
// The key comparison function must be provided to
// use most methods, but the value comparison
// function may be omitted (nil) if no methods that
// are called depend on value equality (ie ContainsValue).
type TreeMap[K, V any] struct {
	keyCmp structs.CompareFunc[K]
	valEq  structs.EqualFunc[V]
	root   *treeNode[K, V]
	size   int
}

func New[K, V any](keyCompare structs.CompareFunc[K], valueEqual structs.EqualFunc[V]) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		keyCmp: keyCompare,
		valEq:  valueEqual,
	}
}

func NewOrdered[K, V constraints.Ordered]() *TreeMap[K, V] {
	return &TreeMap[K, V]{
		keyCmp: structs.CompareOrdered[K],
		valEq:  structs.EqualOrdered[V],
	}
}

func NewOrderedKeys[K constraints.Ordered, V any](valueEqual structs.EqualFunc[V]) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		keyCmp: structs.CompareOrdered[K],
		valEq:  valueEqual,
	}
}

func NewOrderedValues[K any, V constraints.Ordered](keyCompare structs.CompareFunc[K]) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		keyCmp: keyCompare,
		valEq:  structs.EqualOrdered[V],
	}
}

func (t *TreeMap[K, V]) ContainsKey(key K) bool {
	return t.getNode(key) != nil
}

func (t *TreeMap[K, V]) ContainsValue(value V) bool {
	return t.findNode(value) != nil
}

func (t *TreeMap[K, V]) Get(key K) V {
	node := t.getNode(key)
	if node == nil {
		var null V
		return null
	}
	return node.Value
}

func (t *TreeMap[K, V]) GetKey(value V) K {
	node := t.findNode(value)
	if node == nil {
		var null K
		return null
	}
	return node.Key
}

func (t *TreeMap[K, V]) Put(key K, value V) V {
	if t.root == nil {
		t.root = newNode(key, value, true)
		t.size = 1
		var null V
		return null
	}

	var inserted *treeNode[K, V]
	node := t.root
loop:
	for {
		cmp := t.keyCmp(key, node.Key)
		switch {
		case cmp == 0:
			oldValue := node.Value
			node.Key = key
			node.Value = value
			return oldValue
		case cmp < 0:
			if node.Left == nil {
				node.Left = newNode(key, value, false)
				inserted = node.Left
				break loop
			}
			node = node.Left
		case cmp > 0:
			if node.Right == nil {
				node.Right = newNode(key, value, false)
				inserted = node.Right
				break loop
			}
			node = node.Right
		}
	}
	inserted.Parent = node

	t.insertFixup(inserted)
	t.size++
	var null V
	return null
}

func (t *TreeMap[K, V]) Remove(key K) V {
	node := t.getNode(key)
	if node == nil {
		var null V
		return null
	}
	deletedValue := node.Value
	t.removeNode(node)
	return deletedValue
}

func (t *TreeMap[K, V]) Size() int {
	return t.size
}

func (t *TreeMap[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

func (t *TreeMap[K, V]) EntryIterator() *EntryIterator[K, V] {
	return &EntryIterator[K, V]{
		iter: t.nodeIterator(),
	}
}

func (t *TreeMap[K, V]) DescendingEntryIterator() *EntryIterator[K, V] {
	return &EntryIterator[K, V]{
		iter: t.descendingNodeIterator(),
	}
}

func (t *TreeMap[K, V]) KeyIterator() *KeyIterator[K, V] {
	return &KeyIterator[K, V]{
		iter: t.nodeIterator(),
	}
}

func (t *TreeMap[K, V]) DescendingKeyIterator() *DescendingKeyIterator[K, V] {
	return &DescendingKeyIterator[K, V]{
		iter: t.descendingNodeIterator(),
	}
}

func (t *TreeMap[K, V]) Iterator() *ValueIterator[K, V] {
	return &ValueIterator[K, V]{
		iter: t.nodeIterator(),
	}
}

func (t *TreeMap[K, V]) DescendingIterator() *ValueIterator[K, V] {
	return &ValueIterator[K, V]{
		iter: t.descendingNodeIterator(),
	}
}

func (t *TreeMap[K, V]) insertFixup(node *treeNode[K, V]) {
	if node.Parent == nil {
		node.Black = true
		return
	}

	if node.Parent.isBlack() {
		return
	}

	uncle := node.uncle()
	grandparent := node.grandparent()
	if !uncle.isBlack() {
		node.Parent.Black = true
		uncle.Black = true
		grandparent.Black = false
		t.insertFixup(grandparent)
		return
	}

	if node == node.Parent.Right && node.Parent == grandparent.Left {
		t.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		t.rotateRight(node.Parent)
		node = node.Right
	}

	node.Parent.Black = true
	grandparent.Black = false
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		t.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		t.rotateLeft(grandparent)
	}
}

func (t *TreeMap[K, V]) removeFixup(node *treeNode[K, V]) {
	if node.Parent == nil {
		return
	}

	sibling := node.sibling()
	if !sibling.isBlack() {
		node.Parent.Black = false
		sibling.Black = true
		if node == node.Parent.Left {
			t.rotateLeft(node.Parent)
		} else {
			t.rotateRight(node.Parent)
		}
	}

	if sibling.isBlack() && sibling.Left.isBlack() && sibling.Right.isBlack() {
		if node.Parent.isBlack() {
			sibling.Black = false
			t.removeFixup(node.Parent)
		} else {
			sibling.Black = false
			node.Parent.Black = true
		}
		return
	}

	if node == node.Parent.Left && sibling.isBlack() && !sibling.Left.isBlack() && sibling.Right.isBlack() {
		sibling.Black = false
		sibling.Left.Black = true
		t.rotateRight(sibling)
	} else if node == node.Parent.Right && sibling.isBlack() && !sibling.Right.isBlack() && sibling.Left.isBlack() {
		sibling.Black = false
		sibling.Right.Black = true
		t.rotateLeft(sibling)
	}

	sibling.Black = node.Parent.isBlack()
	node.Parent.Black = true
	if node == node.Parent.Left && !sibling.Right.isBlack() {
		sibling.Right.Black = true
		t.rotateLeft(node.Parent)
	} else if !sibling.Left.isBlack() {
		sibling.Left.Black = false
		t.rotateRight(node.Parent)
	}
}

func (t *TreeMap[K, V]) rotateLeft(node *treeNode[K, V]) {
	right := node.Right
	t.replace(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (t *TreeMap[K, V]) rotateRight(node *treeNode[K, V]) {
	left := node.Left
	t.replace(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (t *TreeMap[K, V]) replace(oldNode *treeNode[K, V], newNode *treeNode[K, V]) {
	if oldNode.Parent == nil {
		t.root = newNode
	} else {
		if oldNode == oldNode.Parent.Left {
			oldNode.Parent.Left = newNode
		} else {
			oldNode.Parent.Right = newNode
		}
	}
	if newNode != nil {
		newNode.Parent = oldNode.Parent
	}
}

func (t *TreeMap[K, V]) getNode(key K) *treeNode[K, V] {
	node := t.root
	for node != nil {
		cmp := t.keyCmp(key, node.Key)
		if cmp == 0 {
			break
		}
		if cmp < 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

func (t *TreeMap[K, V]) getLastNode() *treeNode[K, V] {
	node := t.root
	if node != nil {
		for node.Right != nil {
			node = node.Right
		}
	}
	return node
}

func (t *TreeMap[K, V]) findNode(value V) *treeNode[K, V] {
	node := t.root
	for node != nil {
		if t.valEq(value, node.Value) {
			break
		}
		node = node.successor()
	}
	return node
}

func (t *TreeMap[K, V]) findLastNode(value V) *treeNode[K, V] {
	node := t.getLastNode()
	for node != nil {
		if t.valEq(value, node.Value) {
			break
		}
		node = node.predecessor()
	}
	return node
}

func (t *TreeMap[K, V]) removeNode(node *treeNode[K, V]) {
	var child *treeNode[K, V]
	if node.Left != nil && node.Right != nil {
		prev := node.Left.maximumNode()
		node.Key = prev.Key
		node.Value = prev.Value
		node = prev
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.Black {
			node.Black = child.isBlack()
			t.removeFixup(node)
		}
		t.replace(node, child)
		if node.Parent == nil && child != nil {
			child.Black = true
		}
	}
	t.size--
}

func (t *TreeMap[K, V]) nodeIterator() *nodeIterator[K, V] {
	return t.nodeIteratorAt(t.root)
}

func (t *TreeMap[K, V]) descendingNodeIterator() *nodeIterator[K, V] {
	return t.nodeIteratorAt(t.getLastNode())
}

func (t *TreeMap[K, V]) nodeIteratorAt(node *treeNode[K, V]) *nodeIterator[K, V] {
	return &nodeIterator[K, V]{
		treemap: t,
		next:    node,
		last:    nil,
	}
}
