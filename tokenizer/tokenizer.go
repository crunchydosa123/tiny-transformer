package tokenizer

import (
	"fmt"
	"strings"
)

type Tokenizer struct {
	Vocab    map[string]int
	RevVocab map[int]string
	NextID   int
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		Vocab:    make(map[string]int),
		RevVocab: make(map[int]string),
		NextID:   0,
	}
}

func (t *Tokenizer) Fit(texts []string) {
	for _, text := range texts {
		words := strings.Fields(strings.ToLower(text))
		for _, word := range words {
			if _, exists := t.Vocab[word]; !exists {
				t.Vocab[word] = t.NextID
				t.RevVocab[t.NextID] = word
				t.NextID++
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
		}
	}

	return tokens
}

func (t *Tokenizer) Decode(tokens []int) string {
	var words []string
	for _, token := range tokens {
		if word, exists := t.RevVocab[token]; exists {
			words = append(words, word)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	tokenizer := NewTokenizer()

	data := []string{
		"list files",
		"delete temp folder",
		"find all txt files",
	}

	tokenizer.Fit(data)

	encoded := tokenizer.Encode("find all txt files")
	fmt.Println(encoded)

	decoded := tokenizer.Decode(encoded)
	fmt.Println(decoded)
}
