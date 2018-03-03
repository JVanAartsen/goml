package utils

import (
	"strings"
	"strconv"
	"log"
)



type Instance struct {
	Attributes []interface{}
	ClassArray []interface{}
}
type InstanceFactory struct {
	DataType string
	Delim string
	NumAttr int
	NumClassVars int
}



func (instanceFactory *InstanceFactory) LineToInstance(line string) Instance {
	stringAttributes := strings.Split(line, instanceFactory.Delim)

	attributes := make([]interface{}, instanceFactory.NumAttr)
	classVars := make([]interface{}, instanceFactory.NumClassVars)
	var i int
	for i = 0; i < instanceFactory.NumAttr; i++ {
		attributes[i] = parse(instanceFactory.DataType)(stringAttributes[i])
	}
	for j := 0; j < instanceFactory.NumClassVars; i, j = i+1, j+1 {
		classVars[j] = parse(instanceFactory.DataType)(stringAttributes[i])
	}
	return Instance{attributes, classVars}
}


func parse(dataType string) func(str string) interface{} {
	switch dataType {
	case "float": return parseFloat
	case "int": return parseInt
	default: {
		log.Fatal("Invalid DataType :", dataType)
		return nil
	}
	}
}

func parseFloat(str string) interface{} {
	ret, err := strconv.ParseFloat(str, 64)
	if err != nil { log.Fatal(err) }
	return ret
}

func parseInt(str string) interface{} {
	ret, err := strconv.Atoi(str)
	if err != nil { log.Fatal(err) }
	return ret
}
