package main

type Node struct {
	Depth         int     // 深度
	ChildrenSlice []*Node // 子節點
	Parent        *Node   // 父節點
	Char          string
	AttrMap       map[string]string
	Key           string
}

type Root struct {
	rootPtr    *Node
	currentPtr *Node
}

func (r *Root) FindParent() *Node {
	return r.currentPtr.Parent
}

func (r *Root) AppendChildren(n *Node) {
	r.currentPtr.ChildrenSlice = append(r.currentPtr.ChildrenSlice, n)
	n.Parent = r.currentPtr
}
