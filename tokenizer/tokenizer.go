package main

import (
	"strings"
)

type Tokenizer struct {
	Vocab        map[string]int
	ReverseVocab map[int]string
	NextID       int
}

func NewTokenizer() *Tokenizer {
	t := &Tokenizer{
		Vocab:        make(map[string]int),
		ReverseVocab: make(map[int]string),
		NextID:       0,
	}

	t.AddToken("<PAD>")
	t.AddToken("<UNK>")

	return t
}

func (t *Tokenizer) AddToken(token string) {
	t.Vocab[token] = t.NextID
	t.ReverseVocab[t.NextID] = token
	t.NextID++
}

func (t *Tokenizer) Fit(texts []string) {
	for _, text := range texts {
		words := strings.Fields(strings.ToLower(text))

		for _, word := range words {
			if _, exists := t.Vocab[word]; !exists {
				t.AddToken(word)
			}
		}
	}
}

func (t *Tokenizer) Encode(text string) []int {
	words := strings.Fields(strings.ToLower(text))

	var tokens []int

	for _, word := range words {
		if id, exists := t.Vocab[word]; exists {
			tokens = append(tokens, id)
		} else {
			tokens = append(tokens, t.Vocab["<UNK>"])
		}
	}

	return tokens
}
