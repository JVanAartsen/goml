package utils

import (
	"math/rand"
	"os"
	"log"
	"bufio"
)
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

type DatasetFactory struct {
	DataType string
	Delim string
	NumAttr int
	NumClassVars int
}


func (datasetFactory *DatasetFactory) FileToSet(filePath string) []Instance {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	instanceFactory := InstanceFactory{
		datasetFactory.DataType,
		datasetFactory.Delim,
		datasetFactory.NumAttr,
		datasetFactory.NumClassVars}

	var set []Instance

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		set = append(set, instanceFactory.LineToInstance(scanner.Text()))
	}
	return set
}
