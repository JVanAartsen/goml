package utils

import (
	"strings"
	"strconv"
	"log"
)


type Instance struct {
	Attributes []float64
	ClassArray []int
}

func LineToInstance(line string, delim string, numAttr int, numClassVars int) *Instance {
	stringAttributes := strings.Split(line, delim)

	attributes := make([]float64, numAttr)
	classVars := make([]int, numClassVars)

	var i int
	var err error

	for i = 0; i < numAttr; i++ {
		attributes[i], err = strconv.ParseFloat(stringAttributes[i], 64)
		if err != nil { log.Fatal(err) }
	}
	for j := 0; j < numClassVars; j++ {
		classVars[j], err = strconv.Atoi(stringAttributes[i])
		if err != nil { log.Fatal(err) }
		i++
	}
	return &Instance{attributes, classVars}
}
