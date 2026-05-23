package embeddings

import "math/rand"

type PositionalEmbedding struct {
	Weights [][]float64
}

func NewPositionalEmbedding(maxLen, embedDim int) *PositionalEmbedding {
	weights := make([][]float64, maxLen)

	for i := range weights {
		weights[i] = make([]float64, embedDim)

		for j := range weights[i] {
			weights[i][j] = rand.Float64()*0.02 - 0.01
		}
	}

	return &PositionalEmbedding{
		Weights: weights,
	}
}

func AddEmbeddings(
	tokenEmbeds [][]float64,
	posEmbeds [][]float64,
) [][]float64 {

	output := make([][]float64, len(tokenEmbeds))

	for i := range tokenEmbeds {
		output[i] = make([]float64, len(tokenEmbeds[i]))

		for j := range tokenEmbeds[i] {
			output[i][j] =
				tokenEmbeds[i][j] +
					posEmbeds[i][j]
		}
	}

	return output
}
