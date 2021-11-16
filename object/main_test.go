package main

import (
	"reflect"
	"testing"
)

func TestMiniTrie(t *testing.T) {
	runeToInt := func(r rune) int {
		if !('a' <= r && r <= 'c') {
			panic("specify a or b or c")
		}
		return int(r - 'a')
	}
	trie := newTrie(runeToInt, 3)
	trie.addWord("bbc", 1)
	trie.addWord("cbc", 2)
	trie.addWord("cc", 3)

	tests := []struct {
		title  string
		word   string
		output []int
	}{
		{
			title:  "leaf",
			word:   "bbc",
			output: []int{1},
		},
		{
			title:  "multiple ids",
			word:   "c",
			output: []int{2, 3},
		},
		{
			title:  "non existing",
			word:   "a",
			output: []int{},
		},
		{
			title:  "too long chain",
			word:   "ccc",
			output: []int{},
		},
	}

	for _, td := range tests {
		node := trie.node(td.word)
		ids := node.ids()
		if !reflect.DeepEqual(ids, td.output) {
			t.Fatal(ids, "is not", td.output, "//test case for", td.title, "with word:", td.word)
		}
	}
}

func TestMediumTrie(t *testing.T) {
	runeToInt := func(r rune) int {
		if !(r == ' ' || ('a' <= r && r <= 'z')) {
			panic("specify a to z or space")
		}
		if r == ' ' {
			return 0
		}
		return int(r-'a') + 1
	}
	trie := newTrie(runeToInt, 28)
	trie.addWord("autumn", 1)
	trie.addWord("dock", 2)
	trie.addWord("indian summer", 3)
	trie.addWord("spring", 4)
	trie.addWord("sum", 5)
	trie.addWord("sumerian", 6)
	trie.addWord("summer", 7)
	trie.addWord("winter", 8)

	tests := []struct {
		title  string
		word   string
		output []int
	}{
		{
			title:  "leaf",
			word:   "indian summer",
			output: []int{3},
		},
		{
			title:  "leaf midway",
			word:   "indian su",
			output: []int{3},
		},
		{
			title:  "multiple ids",
			word:   "sum",
			output: []int{5, 6, 7},
		},
		{
			title:  "multiple ids midway",
			word:   "s",
			output: []int{4, 5, 6, 7},
		},
		{
			title:  "non existing",
			word:   "z",
			output: []int{},
		},
		{
			title:  "too long chain",
			word:   "docket",
			output: []int{},
		},
	}

	for _, td := range tests {
		node := trie.node(td.word)
		ids := node.ids()
		if !reflect.DeepEqual(ids, td.output) {
			t.Fatal(ids, "is not", td.output, "//test case for", td.title, "with word:", td.word)
		}
	}
}
