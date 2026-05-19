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

type BPETokenizer struct {
	Vocab        map[string]int
	ReverseVocab map[int]string
	Merges       map[string]string
	NextID       int
}

func CountPairs(words [][]string) map[string]int {
	pairs := make(map[string]int)

	for _, word := range words {
		for i := 0; i < len(word)-1; i++ {
			pair := word[i] + " " + word[i+1]
			pairs[pair]++
		}
	}

	return pairs
}

func BestPair(pairs map[string]int) string {
	max := -1
	best := ""

	for pair, count := range pairs {
		if count > max {
			max = count
			best = pair
		}
	}

	return best
}

func MergePair(words [][]string, pair string) [][]string {
	parts := strings.Split(pair, " ")
	a, b := parts[0], parts[1]

	var result [][]string

	for _, word := range words {
		var merged []string

		i := 0
		for i < len(word) {
			if i < len(word)-1 &&
				word[i] == a &&
				word[i+1] == b {

				merged = append(merged, a+b)
				i += 2
			} else {
				merged = append(merged, word[i])
				i++
			}
		}

		result = append(result, merged)
	}

	return result
}

func (b *BPETokenizer) Train(texts []string, numMerges int) {
	var words [][]string

	for _, text := range texts {
		for _, word := range strings.Fields(text) {
			chars := strings.Split(word, "")
			words = append(words, chars)
		}
	}

	for i := 0; i < numMerges; i++ {
		pairs := CountPairs(words)

		best := BestPair(pairs)

		if best == "" {
			break
		}

		fmt.Println("Merging:", best)

		words = MergePair(words, best)

		token := strings.ReplaceAll(best, " ", "")
		b.Merges[best] = token

		if _, exists := b.Vocab[token]; !exists {
			b.Vocab[token] = b.NextID
			b.ReverseVocab[b.NextID] = token
			b.NextID++
		}
	}
}

func (b *BPETokenizer) Encode(word string) []string {
	tokens := strings.Split(word, "")

	for {
		changed := false

		for i := 0; i < len(tokens)-1; i++ {
			pair := tokens[i] + " " + tokens[i+1]

			if merged, exists := b.Merges[pair]; exists {
				newTokens := append([]string{}, tokens[:i]...)
				newTokens = append(newTokens, merged)
				newTokens = append(newTokens, tokens[i+2:]...)

				tokens = newTokens
				changed = true
				break
			}
		}

		if !changed {
			break
		}
	}

	return tokens
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
