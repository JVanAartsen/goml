package main

import (
	"os"
	"bufio"
	"log"
	"github.com/jvanaartsen/goml/utils"
	"github.com/jvanaartsen/goml/neural"
)


const NUM_ATTRIBUTES = 256
const NUM_HIDDEN = 20
const NUM_CLASS_VARS = 10


func main() {

	file, err := os.Open("/Users/jovanaartsen/play/datasets/semeion-digits.txt")
	if err != nil {
		log.Fatal(err)
	}

	var set []utils.Instance

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instance := utils.LineToInstance(scanner.Text(), " ", NUM_ATTRIBUTES, NUM_CLASS_VARS)
		set = append(set, instance)
	}


	trainingSet, testSet := utils.TrainTestSplit(utils.Shuffle(set), 0.2)


	myNN := neural.New(NUM_ATTRIBUTES, NUM_HIDDEN, NUM_CLASS_VARS)
	myNN.Train(trainingSet)
	myNN.Test(testSet)
	//PrintNeuralNetwork(myNN)
}
