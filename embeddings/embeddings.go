package embeddings

import (
	"math/rand"
)

type Embedding struct {
	Weights [][]float64
}

func NewEmbedding(vocabSize, embedDim int) *Embedding {
	weights := make([][]float64, vocabSize)

	for i := range weights {
		weights[i] = make([]float64, embedDim)

		for j := range weights[i] {
			weights[i][j] = rand.Float64()*0.02 - 0.01
		}
	}

	return &Embedding{
		Weights: weights,
	}
}

func (e *Embedding) Forward(tokens []int) [][]float64 {
	output := make([][]float64, len(tokens))

	for i, token := range tokens {
		output[i] = e.Weights[token]
	}

	return output
}
