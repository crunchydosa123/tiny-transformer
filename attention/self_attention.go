package attention

type SelfAttention struct {
	Wq [][]float64
	Wk [][]float64
	Wv [][]float64

	EmbedDim int
}

func NewSelfAttention(embedDim int) *SelfAttention
