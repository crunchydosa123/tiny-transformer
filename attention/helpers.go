package attention

import "math"

func MatMul(a, b [][]float64) [][]float64 {
	rows := len(a)
	cols := len(b[0])
	inner := len(b)

	result := make([][]float64, rows)

	for i := range result {
		result[i] = make([]float64, cols)

		for j := 0; j < cols; j++ {
			sum := 0.0

			for k := 0; k < inner; k++ {
				sum += a[i][k] * b[k][j]
			}

			result[i][j] = sum
		}
	}

	return result
}

func Softmax(x []float64) []float64 {
	maxVal := x[0]

	for _, v := range x {
		if v > maxVal {
			maxVal = v
		}
	}

	expVals := make([]float64, len(x))
	sum := 0.0

	for i, v := range x {
		expVals[i] = math.Exp(v - maxVal)
		sum += expVals[i]
	}

	for i := range expVals {
		expVals[i] /= sum
	}

	return expVals

}
