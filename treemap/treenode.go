package treemap

type TreeNode[K, V any] struct {
	Key    K
	Value  V
	Black  bool
	Parent *TreeNode[K, V]
	Left   *TreeNode[K, V]
	Right  *TreeNode[K, V]
}

func newNode[K, V any](key K, value V, black bool) *TreeNode[K, V] {
	return &TreeNode[K, V]{
		Key:    key,
		Value:  value,
		Black:  black,
		Parent: nil,
		Left:   nil,
		Right:  nil,
	}
}

func (n *TreeNode[K, V]) grandparent() *TreeNode[K, V] {
	if n == nil || n.Parent == nil {
		return nil
	}
	return n.Parent.Parent
}

func (n *TreeNode[K, V]) uncle() *TreeNode[K, V] {
	if n == nil || n.Parent == nil || n.Parent.Parent == nil {
		return nil
	}
	return n.Parent.sibling()
}

func (n *TreeNode[K, V]) sibling() *TreeNode[K, V] {
	if n == nil || n.Parent == nil {
		return nil
	}
	if n == n.Parent.Left {
		return n.Parent.Right
	}
	return n.Parent.Left
}

func (n *TreeNode[K, V]) maximumNode() *TreeNode[K, V] {
	if n == nil {
		return nil
	}
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *TreeNode[K, V]) predecessor() *TreeNode[K, V] {
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

func (n *TreeNode[K, V]) successor() *TreeNode[K, V] {
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
