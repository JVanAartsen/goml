package main

import (
	"github.com/jvanaartsen/goml/utils"
	"github.com/jvanaartsen/goml/neural"
)


const NUM_ATTRIBUTES = 256
const NUM_HIDDEN = 20
const NUM_CLASS_VARS = 10


func main() {

	datasetFactory := utils.DatasetFactory{"float", " ", NUM_ATTRIBUTES, NUM_CLASS_VARS}

	set := datasetFactory.FileToSet("/Users/jovanaartsen/play/datasets/semeion-digits.txt")

	trainingSet, testSet := utils.TrainTestSplit(utils.Shuffle(set), 0.2)


	myNN := neural.New(NUM_ATTRIBUTES, NUM_HIDDEN, NUM_CLASS_VARS)
	myNN.Train(trainingSet)
	myNN.Test(testSet)
	//PrintNeuralNetwork(myNN)
}
