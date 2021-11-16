package main

type Stack struct {
	data []*node
}

func (s *Stack) push(n *node) {
	s.data = append(s.data, n)
}

func (s *Stack) pop() *node {
	last := len(s.data) - 1
	if last < 0 {
		panic("stack is empty")
	}
	n := s.data[last]
	s.data = s.data[:last]
	return n
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (t *trie) node(word string) *node {
	initialRow := 0
	rootNode := node{
		OffspringRow: &initialRow,
	}
	stack := Stack{
		data: []*node{
			&rootNode,
		},
	}
	for _, r := range word {
		if stack.isEmpty() {
			return nil
		}
		parent := stack.pop()
		if parent.OffspringRow != nil {
			i := int(*parent.OffspringRow)
			row := &t.Nodes[i]
			stack.push(&(*row)[t.atoi(r)])
		}
	}
	if stack.isEmpty() {
		return nil
	}
	node := stack.pop()
	if node.isPhantom() {
		return nil
	}
	return node
}

func (t *trie) ids(n *node) []int {
	ids := make([]int, 0)
	if n == nil {
		return ids
	}
	stack := Stack{
		data: []*node{
			n,
		},
	}
	for !stack.isEmpty() {
		parent := stack.pop()
		if parent.Id != nil {
			id := int(*parent.Id)
			ids = append(ids, id)
		}
		if parent.OffspringRow == nil {
			continue
		}
		for i := t.nOfLetters - 1; i >= 0; i-- {
			r := int(*parent.OffspringRow)
			row := &t.Nodes[r]
			stack.push(&(*row)[i])
		}
	}
	return ids
}
