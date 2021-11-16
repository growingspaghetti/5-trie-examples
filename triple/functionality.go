package main

type Stack struct {
	data   []*node
	ordial []int
}

func newStack() *Stack {
	return &Stack{
		data:   make([]*node, 0),
		ordial: make([]int, 0),
	}
}

func (s *Stack) push(n *node, order int) {
	s.data = append(s.data, n)
	s.ordial = append(s.ordial, order)
}

func (s *Stack) pop() (*node, int) {
	last := len(s.data) - 1
	if last < 0 {
		panic("stack is empty")
	}
	n := s.data[last]
	s.data = s.data[:last]
	o := s.ordial[last]
	s.ordial = s.ordial[:last]
	return n, o
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *sparse) addParentInfo() {
	initialRow := 0
	rootNode := &node{
		OffspringRow: &initialRow,
	}
	stack := newStack()
	stack.push(rootNode, -1)
	for !stack.isEmpty() {
		parent, order := stack.pop()
		if parent.OffspringRow == nil {
			continue
		}
		for i := s.nOfLetters - 1; i >= 0; i-- {
			r := int(*parent.OffspringRow)
			node := &s.Nodes[r+i]
			if node.isPhantom() {
				continue
			}
			node.Check = &order
			stack.push(node, r+i)
		}
	}
}

func (t *trie) accumBase(row int) int {
	acc := 0
	for i := 0; i <= row; i++ {
		acc += t.Base[i]
	}
	return acc
}

func (t *trie) node(word string) (*node, int) {
	initialRow := 0
	rootNode := &node{
		OffspringRow: &initialRow,
	}
	stack := newStack()
	stack.push(rootNode, -1)
	for _, r := range word {
		if stack.isEmpty() {
			return nil, -1
		}
		parent, order := stack.pop()
		if parent.OffspringRow == nil {
			return nil, -1
		}
		i := int(*parent.OffspringRow)
		row := i / t.nOfLetters
		dest := t.accumBase(row) + t.atoi(r)
		if dest > len(t.Nodes)-1 {
			return nil, -1
		}
		node := &t.Nodes[dest]
		if node.Check == nil {
			return nil, -1
		}
		if *node.Check != order {
			return nil, -1
		}
		stack.push(node, i+t.atoi(r))
	}
	if stack.isEmpty() {
		return nil, -1
	}
	node, order := stack.pop()
	if node.isPhantom() {
		return nil, -1
	}
	return node, order
}

func (t *trie) ids(n *node, order int) []int {
	ids := make([]int, 0)
	if n == nil {
		return ids
	}
	stack := newStack()
	stack.push(n, order)
	for !stack.isEmpty() {
		parent, ord := stack.pop()
		if parent.Id != nil {
			id := int(*parent.Id)
			ids = append(ids, id)
		}
		if parent.OffspringRow == nil {
			continue
		}
		r := int(*parent.OffspringRow)
		base := t.accumBase(r / t.nOfLetters)
		for i := t.nOfLetters - 1; i >= 0; i-- {
			dest := base + i
			if dest > len(t.Nodes)-1 {
				continue
			}
			node := &t.Nodes[dest]
			if node.Check == nil {
				continue
			}
			if *node.Check != ord {
				continue
			}
			stack.push(node, r+i)
		}
	}
	return ids
}
