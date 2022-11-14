package treemap

type treeNode[K, V any] struct {
	Key    K
	Value  V
	Black  bool
	Parent *treeNode[K, V]
	Left   *treeNode[K, V]
	Right  *treeNode[K, V]
}

func newNode[K, V any](key K, value V, black bool) *treeNode[K, V] {
	return &treeNode[K, V]{
		Key:    key,
		Value:  value,
		Black:  black,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}
}

func (n *treeNode[K, V]) grandparent() *treeNode[K, V] {
	if n == nil || n.Parent == nil {
		return nil
	}
	return n.Parent.Parent
}

func (n *treeNode[K, V]) uncle() *treeNode[K, V] {
	if n == nil || n.Parent == nil || n.Parent.Parent == nil {
		return nil
	}
	return n.Parent.sibling()
}

func (n *treeNode[K, V]) sibling() *treeNode[K, V] {
	if n == nil || n.Parent == nil {
		return nil
	}
	if n == n.Parent.Left {
		return n.Parent.Right
	}
	return n.Parent.Left
}

func (n *treeNode[K, V]) maximumNode() *treeNode[K, V] {
	if n == nil {
		return nil
	}
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *treeNode[K, V]) predecessor() *treeNode[K, V] {
	if n == nil {
		return nil
	}

	if n.Left != nil {
		p := n.Left
		for p.Right != nil {
			p = p.Right
		}
		return p
	}

	parent := n.Parent
	tmp := n
	for tmp != nil && tmp == parent.Left {
		tmp = parent
		parent = parent.Parent
	}

	return parent
}

func (n *treeNode[K, V]) successor() *treeNode[K, V] {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		p := n.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	parent := n.Parent
	tmp := n
	for tmp != nil && tmp == parent.Right {
		tmp = parent
		parent = parent.Parent
	}

	return parent
}

func (n *treeNode[K, V]) isBlack() bool {
	return n == nil || n.Black
}
