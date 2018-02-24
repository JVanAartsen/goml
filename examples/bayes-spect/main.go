package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"github.com/jvanaartsen/goml/utils"
)


func main() {

	file, err := os.Open("/Users/jovanaartsen/play/datasets/SPECT.csv")
	if err != nil {
		log.Fatal(err)
	}

	var set []utils.Instance

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		set = append(set, utils.LineToInstance(scanner.Text(), ",", 22, 1))
	}

	fmt.Println(set)
}
