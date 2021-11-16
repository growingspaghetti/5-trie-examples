package main

import (
	"errors"
	"fmt"
)

func main() {
	runeToInt := func(r rune) int {
		if !('a' <= r && r <= 'c') {
			panic("specify a or b or c")
		}
		return int(r - 'a')
	}
	sparse := newSparse(runeToInt, 3)
	sparse.addWord("bbc", 1)
	sparse.addWord("cbc", 2)
	sparse.addWord("cc", 3)
	fmt.Println(prettyPrint(sparse))

	t := newTrie(sparse)
	fmt.Println(prettyPrint(t))

	bbcNode, ord := t.node("bbc")
	fmt.Println(prettyPrint(bbcNode))
	fmt.Println(t.ids(bbcNode, ord))

	cnodes, ord := t.node("c")
	fmt.Println(prettyPrint(cnodes))
	fmt.Println(t.ids(cnodes, ord))

	aa, ord := t.node("a")
	fmt.Println(prettyPrint(aa))
	fmt.Println(t.ids(aa, ord))
}

type sparse struct {
	Nodes      []node
	atoi       func(rune) int
	nOfLetters int
}

type trie struct {
	Nodes      []node
	atoi       func(rune) int
	nOfLetters int
}

type node struct {
	OffspringRow *int   `json:",omitempty"`
	Id           *int   `json:",omitempty"`
	Debug        string `json:",omitempty"`
	Check        *int   `json:",omitempty"`
}

func (t *trie) mergeablePosition(row []node) (int, error) {
Shift:
	for i := 0; i <= len(t.Nodes)-t.nOfLetters; i++ {
		for j, node := range row {
			if !node.isPhantom() && !t.Nodes[i+j].isPhantom() {
				continue Shift
			}
		}
		return i, nil
	}
	return -1, errors.New("no position found")
}

func (t *trie) merge(row []node, position int) {
	for j, node := range row {
		if !node.isPhantom() {
			t.Nodes[position+j] = node
		}
	}
}

func newTrie(s *sparse) *trie {
	s.addParentInfo()
	t := &trie{
		Nodes:      make([]node, 0),
		atoi:       s.atoi,
		nOfLetters: s.nOfLetters,
	}
	positionMap := make(map[int]int)
	for i := 0; i < len(s.Nodes); i += s.nOfLetters {
		row := s.Nodes[i : i+s.nOfLetters]
		pos, err := t.mergeablePosition(row)
		if err != nil {
			positionMap[i] = len(t.Nodes)
			t.Nodes = append(t.Nodes, row...)
			continue
		}
		t.merge(row, pos)
		positionMap[i] = pos
	}
	for i := 0; i < len(t.Nodes); i++ {
		n := &t.Nodes[i]
		if n.OffspringRow != nil {
			p := positionMap[*n.OffspringRow]
			n.OffspringRow = &p
		}
		if n.Check != nil {
			shift := *n.Check % t.nOfLetters
			p := positionMap[*n.Check-shift]
			c := p + shift
			n.Check = &c
		}
	}
	return t
}

func newSparse(atoi func(rune) int, nOfLetters int) *sparse {
	return &sparse{
		Nodes:      make([]node, nOfLetters),
		atoi:       atoi,
		nOfLetters: nOfLetters,
	}
}

func (n *node) initialize(debug rune, rest []rune, id int, t *sparse) {
	n.Debug = string(debug)
	if n.OffspringRow == nil {
		row := len(t.Nodes)
		n.OffspringRow = &row
		t.Nodes = append(t.Nodes, make([]node, t.nOfLetters)...)
	}
	if len(rest) == 0 {
		id := int(id)
		n.Id = &id
		return
	}
	n.addChars(rest, id, t)
}

func (n *node) addChars(chars []rune, id int, t *sparse) {
	if len(chars) == 0 {
		return
	}
	r := chars[0]
	row := *n.OffspringRow
	child := &(t.Nodes[row+t.atoi(r)])
	child.initialize(r, chars[1:], id, t)
}

func (t *sparse) addChars(chars []rune, id int) {
	initialRow := 0
	rootNode := &node{
		OffspringRow: &initialRow,
	}
	rootNode.addChars(chars, id, t)
}

func (n *node) isPhantom() bool {
	return n.OffspringRow == nil && n.Id == nil
}

func (t *sparse) addWord(word string, id int) {
	t.addChars([]rune(word), id)
}
