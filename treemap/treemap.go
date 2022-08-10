package treemap

import (
	"github.com/zytekaron/structs"
)

type TreeMap[K, V any] struct {
	keyCmp structs.CompareFunc[K]
	valCmp structs.CompareFunc[V]
	root   *TreeNode[K, V]
	size   int
}

type TreeNode[K, V any] struct {
	Key    K
	Value  V
	Black  bool
	Parent *TreeNode[K, V]
	Left   *TreeNode[K, V]
	Right  *TreeNode[K, V]
}

func New[K, V any](keyCompare structs.CompareFunc[K], valueCompare structs.CompareFunc[V]) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		keyCmp: keyCompare,
		valCmp: valueCompare,
	}
}

func (t *TreeMap[K, V]) ContainsKey(key K) bool {
	return t.GetNode(key) != nil
}

func (t *TreeMap[K, V]) ContainsValue(value V) bool {
	return t.FindNode(value) != nil
}

func (t *TreeMap[K, V]) GetValue(key K) V {
	node := t.GetNode(key)
	if node == nil {
		var null V
		return null
	}
	return node.Value
}

func (t *TreeMap[K, V]) GetKey(value V) K {
	node := t.FindNode(value)
	if node == nil {
		var null K
		return null
	}
	return node.Key
}

func (t *TreeMap[K, V]) GetNode(key K) *TreeNode[K, V] {
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

func (t *TreeMap[K, V]) FindNode(value V) *TreeNode[K, V] {
	node := t.root
	for node != nil {
		cmp := t.valCmp(value, node.Value)
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

func (t *TreeMap[K, V]) Put(key K, value V) V {
	node := t.root
	if node == nil {
		t.root = &TreeNode[K, V]{
			Key:   key,
			Value: value,
		}
		t.size = 1
		var null V
		return null
	}

	var cmp int
	var parent *TreeNode[K, V]
	// duplicate code since go doesn't have a do/for
	parent = node
	cmp = t.keyCmp(key, node.Key)
	if cmp == 0 {
		node.Value = value
		var null V
		return null
	}
	if cmp < 0 {
		node = node.Left
	} else {
		node = node.Right
	}
	for parent != nil {
		// duplicate code since go doesn't have a do/for
		parent = node
		cmp = t.keyCmp(key, node.Key)
		if cmp == 0 {
			node.Value = value
			var null V
			return null
		}
		if cmp < 0 {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	newNode := &TreeNode[K, V]{
		Parent: parent,
		Key:    key,
		Value:  value,
	}
	if cmp < 0 {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	t.insertFixup(newNode)
	t.size++
	var null V
	return null
}

func (t *TreeMap[K, V]) Size() int {
	return t.size
}

func (t *TreeMap[K, V]) Clear() {
	t.root = nil
	t.size = 0
}

func (t *TreeMap[K, V]) insertFixup(x *TreeNode[K, V]) {
	x.Black = x == t.root
	for x != t.root && !x.Parent.Black {
		if x.Parent == x.Parent.Parent.Left {
			y := x.Parent.Parent.Right
			if y != nil && !y.Black {
				y.Black = true
				x = x.Parent
				x.Black = true
				x = x.Parent
				x.Black = x == t.root
			} else {
				if x != x.Parent.Left {
					x = x.Parent
					rotateLeft(x)
				}
				x = x.Parent
				x.Black = true
				x = x.Parent
				x.Black = false
				rotateRight(x)
				break
			}
		} else {
			y := x.Parent.Parent.Left
			if y != nil && !y.Black {
				y.Black = true
				x = x.Parent
				x.Black = true
				x = x.Parent
				x.Black = x == t.root
			} else {
				if x == x.Parent.Left {
					x = x.Parent
					rotateRight(x)
				}
				x = x.Parent
				x.Black = true
				x = x.Parent
				x.Black = false
				rotateLeft(x)
				break
			}
		}
	}
}

func successor[K, V any](node *TreeNode[K, V]) *TreeNode[K, V] {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		p := node.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	parent := node.Parent
	n := node
	for n == nil && n == parent.Right {
		n = parent
		parent = parent.Parent
	}
	return parent
}

func predecessor[K, V any](node *TreeNode[K, V]) *TreeNode[K, V] {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		p := node.Left
		for p.Right != nil {
			p = p.Right
		}
		return p
	}
	parent := node.Parent
	n := node
	for n == nil && n == parent.Left {
		n = parent
		parent = parent.Parent
	}
	return parent
}

func rotateLeft[K, V any](x *TreeNode[K, V]) {
	y := x.Right
	x.Right = y.Left
	if x.Right != nil {
		x.Right.Parent = x
	}
	y.Parent = x.Parent
	if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func rotateRight[K, V any](x *TreeNode[K, V]) {
	y := x.Left
	x.Left = y.Right
	if x.Left != nil {
		x.Left.Parent = x
	}
	y.Parent = x.Parent
	if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Right = x
	x.Parent = y
}
