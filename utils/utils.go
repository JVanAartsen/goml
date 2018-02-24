package utils

import "math/rand"
// TrainTestSplit splits training and testing instances with the given testSize ratio
func TrainTestSplit(set []Instance, testSize float64) ([]Instance, []Instance) {
	if testSize == 0.0 { testSize = 0.2 }
	splitIndex := int(float64(len(set)) * 0.8)
	return set[0:splitIndex], set[splitIndex:]
}

func Shuffle(set []Instance) []Instance {
	var r1, r2 int
	for range set {
		r1, r2 = rand.Intn(len(set)), rand.Intn(len(set))
		set[r1], set[r2] = set[r2], set[r1]
	}
	return set
}
