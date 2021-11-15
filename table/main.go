package main

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func main() {
	runeToInt := func(r rune) int {
		if !('a' <= r && r <= 'c') {
			panic("specify a or b or c")
		}
		return int(r - 'a')
	}
	t := newTrie(runeToInt, 3)
	t.addWord("bbc", 1)
	t.addWord("cbc", 2)
	t.addWord("cc", 3)

	fmt.Println(prettyPrint(t))

	bbcNode := t.node("bbc")
	fmt.Println(prettyPrint(bbcNode))
	fmt.Println(t.ids(bbcNode))

	cnodes := t.node("c")
	fmt.Println(prettyPrint(cnodes))
	fmt.Println(t.ids(cnodes))
}

type trie struct {
	Nodes      [][]node
	atoi       func(rune) int
	nOfLetters int
}

type node struct {
	OffspringRow *int   `json:",omitempty"`
	Id           *int   `json:",omitempty"`
	Debug        string `json:",omitempty"`
}

func newTrie(atoi func(rune) int, nOfLetters int) *trie {
	return &trie{
		Nodes: [][]node{
			make([]node, nOfLetters),
		},
		atoi:       atoi,
		nOfLetters: nOfLetters,
	}
}

func (n *node) initialize(debug rune, rest []rune, id int, t *trie) {
	n.Debug = string(debug)
	if n.OffspringRow == nil {
		t.Nodes = append(t.Nodes, make([]node, t.nOfLetters))
		row := len(t.Nodes) - 1
		n.OffspringRow = &row
	}
	if len(rest) == 0 {
		id := int(id)
		n.Id = &id
		return
	}
	n.addChars(rest, id, t)
}

func (n *node) addChars(chars []rune, id int, t *trie) {
	if len(chars) == 0 {
		return
	}
	r := chars[0]
	row := *n.OffspringRow
	child := &t.Nodes[row][t.atoi(r)]
	child.initialize(r, chars[1:], id, t)
}

func (t *trie) addChars(chars []rune, id int) {
	initialRow := 0
	rootNode := node{
		OffspringRow: &initialRow,
	}
	rootNode.addChars(chars, id, t)
}

func (n *node) isPhantom() bool {
	return n.OffspringRow == nil && n.Id == nil
}

func (t *trie) addWord(word string, id int) {
	t.addChars([]rune(word), id)
}
