package main

type Stack struct {
	data []*trie
}

func (s *Stack) push(t *trie) {
	s.data = append(s.data, t)
}

func (s *Stack) pop() *trie {
	last := len(s.data) - 1
	if last < 0 {
		panic("stack is empty")
	}
	t := s.data[last]
	s.data = s.data[:last]
	return t
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (t *trie) node(word string) *trie {
	stack := Stack{
		data: []*trie{t},
	}
	for _, r := range word {
		if stack.isEmpty() {
			return nil
		}
		parent := stack.pop()
		if len(parent.Children) == 0 {
			return nil
		}
		child := &parent.Children[t.atoi(r)]
		stack.push(child)
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

func (t *trie) ids() []int {
	ids := make([]int, 0)
	if t == nil {
		return ids
	}
	stack := Stack{
		data: []*trie{t},
	}
	for !stack.isEmpty() {
		parent := stack.pop()
		if parent.Id != nil {
			id := int(*parent.Id)
			ids = append(ids, id)
		}
		for i := len(parent.Children) - 1; i >= 0; i-- {
			stack.push(&parent.Children[i])
		}
	}
	return ids
}
