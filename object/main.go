package main

import (
	"fmt"
)

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
	fmt.Println(bbcNode.ids())

	cnodes := t.node("c")
	fmt.Println(prettyPrint(cnodes))
	fmt.Println(cnodes.ids())
}

type trie struct {
	Children   []trie `json:",omitempty"`
	Id         *int   `json:",omitempty"`
	Debug      string `json:",omitempty"`
	atoi       func(rune) int
	nOfLetters int
}

func newTrie(atoi func(rune) int, nOfLetters int) *trie {
	return &trie{
		Children:   make([]trie, nOfLetters),
		Id:         nil,
		Debug:      string('/'),
		atoi:       atoi,
		nOfLetters: nOfLetters,
	}
}

func (t *trie) initialize(debug rune, rest []rune, id int, parent *trie) {
	t.Debug = string(debug)
	t.atoi = parent.atoi
	t.nOfLetters = parent.nOfLetters
	if len(rest) == 0 {
		id := int(id)
		t.Id = &id
		return
	}
	t.addChars(rest, id)
}

func (t *trie) addChars(chars []rune, id int) {
	if len(chars) == 0 {
		return
	}
	r := chars[0]
	i := t.atoi(r)
	if len(t.Children) != t.nOfLetters {
		t.Children = make([]trie, t.nOfLetters)
	}
	child := &t.Children[i]
	child.initialize(r, chars[1:], id, t)
}

func (t *trie) isPhantom() bool {
	return len(t.Children) == 0 && t.Id == nil
}

func (t *trie) addWord(word string, id int) {
	t.addChars([]rune(word), id)
}
