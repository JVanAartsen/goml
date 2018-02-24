package neural

import "math"

func RectifiedLinearUnit(input float64) float64 {
	return math.Max(float64(0), input)
}

func Sigmoid(input float64) float64 {
	output := 1.0 / (1.0 + math.Exp(-input))
	return output
}

// x is the INPUT of the activation function
// this means the SUM of weighted parents for a node
// which is assigned to Node.inputValue in our impl
func D_ReLU(input float64) float64 {
	if input <= 0.0 {
		return 0.0
	} else {
		return 1.0
	}
}

func D_Sigmoid(input float64) float64 {
	d := Sigmoid(input) * (1 - Sigmoid(input))
	return d
}
