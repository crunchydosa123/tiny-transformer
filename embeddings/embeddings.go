package embeddings

import "math/rand/v2"

type Embedding struct {
	VocabSize int
	EmbedDim  int
	Weights   [][]float64
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
		VocabSize: vocabSize,
		EmbedDim:  embedDim,
		Weights:   weights,
	}
}
