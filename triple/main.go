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
	s := newSparse(runeToInt, 3)
	s.addWord("bbc", 1)
	s.addWord("cbc", 2)
	s.addWord("cc", 3)
	fmt.Println(prettyPrint(s))
	t := newTrie(s)
	fmt.Println(prettyPrint(t))

	bbcNode, ord := t.node("bbc")
	fmt.Println(prettyPrint(bbcNode))
	fmt.Println(t.ids(bbcNode, ord))

	cnodes, ord := t.node("c")
	fmt.Println(prettyPrint(cnodes))
	fmt.Println(t.ids(cnodes, ord))
}

type sparse struct {
	Nodes      []node
	atoi       func(rune) int
	nOfLetters int
}

type trie struct {
	Nodes      []node
	Base       []int
	atoi       func(rune) int
	nOfLetters int
}

type node struct {
	OffspringRow *int   `json:",omitempty"`
	Id           *int   `json:",omitempty"`
	Debug        string `json:",omitempty"`
	Check        *int   `json:",omitempty"`
}

func (t *trie) mergeablePadding(row []node) (int, error) {
	base := len(t.Nodes) - t.nOfLetters
Shift:
	for i := 0; base+i < len(t.Nodes); i++ {
		for j := 0; base+i+j < len(t.Nodes); j++ {
			if !t.Nodes[base+i+j].isPhantom() && !row[j].isPhantom() {
				continue Shift
			}
		}
		return i, nil
	}
	return -1, errors.New("no position found")
}

func (t *trie) merge(row []node, padding int) {
	t.Nodes = append(t.Nodes, make([]node, padding)...)
	base := len(t.Nodes) - t.nOfLetters
	for j, node := range row {
		if !node.isPhantom() {
			t.Nodes[base+j] = node
		}
	}
}

func newTrie(s *sparse) *trie {
	s.addParentInfo()
	t := &trie{
		Nodes:      make([]node, s.nOfLetters),
		atoi:       s.atoi,
		nOfLetters: s.nOfLetters,
	}
	for i := 0; i < len(s.Nodes); i += s.nOfLetters {
		row := s.Nodes[i : i+s.nOfLetters]
		padding, err := t.mergeablePadding(row)
		if err != nil {
			t.Base = append(t.Base, t.nOfLetters)
			t.Nodes = append(t.Nodes, row...)
			continue
		}
		t.merge(row, padding)
		t.Base = append(t.Base, padding)
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
	if len(rest) == 0 {
		id := int(id)
		n.Id = &id
	}
	if n.OffspringRow == nil {
		row := len(t.Nodes)
		n.OffspringRow = &row
		t.Nodes = append(t.Nodes, make([]node, t.nOfLetters)...)
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
