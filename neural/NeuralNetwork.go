package neural

import (
	"fmt"
	"math/rand"
	"github.com/jvanaartsen/goml/utils"
)

const LEARN_RATE = 1
const MAX_EPOCHS = 1
// lets only support 1 hidden layer for now
type NeuralNetwork struct {
	inputNodes []Node
	hiddenNodes []Node
	outputNodes []Node
}

func New(numFeatures int, numHidden int, numClassVars int) *NeuralNetwork {

	inputNodes := make([]Node, numFeatures+1)
	hiddenNodes := make([]Node, numHidden+1)
	outputNodes := make([]Node, numClassVars)

	var nodeWeightPair NodeWeightPair // re assignable nodeweightpair

	// set up inputNodes
	for i := 0 ; i < numFeatures; i++ {
		// parent node will just be a bias node i think
		inputNodes[i] = CreateNode(INPUT)
	}
	inputNodes[numFeatures] = CreateNode(BIAS_TO_HIDDEN)

	// set up hiddenNodes
	for j := 0 ; j < numHidden; j++ {
		hiddenNodes[j] = CreateNode(HIDDEN)
		// set up parents
		for i := range inputNodes {
			nodeWeightPair = NodeWeightPair{&inputNodes[i], rand.Float64()/float64(numFeatures)}
			hiddenNodes[j].Parents = append(hiddenNodes[j].Parents, nodeWeightPair)
		}
	}
	hiddenNodes[numHidden] = CreateNode(BIAS_TO_OUTPUT)

	// set up outputNodes
	for k := range outputNodes {
		outputNodes[k] = CreateNode(OUTPUT)
		for j := range hiddenNodes {
			nodeWeightPair = NodeWeightPair{&hiddenNodes[j], rand.Float64()/float64(numHidden)}
			outputNodes[k].Parents = append(outputNodes[k].Parents, nodeWeightPair)
		}
	}

	return &NeuralNetwork{inputNodes, hiddenNodes, outputNodes}
}


func (nn *NeuralNetwork) Train(trainingSet []utils.Instance) {

	// between input and hidden
	// hidden nodes don't have parents btw
	weights_InputHidden := make([][]float64, len(nn.inputNodes))
	for i := range nn.inputNodes {
		// input node i's weights to hidden nodes, there is no input->hiddenBias weight
		weights_InputHidden[i] = make([]float64, len(nn.hiddenNodes)-1)
		for j := range weights_InputHidden[i] {
			// weight between input node i and hidden node j
			weights_InputHidden[i][j] = nn.hiddenNodes[j].Parents[i].Weight
		}
	}

	// between hidden and outputs
	weights_HiddenOutput := make([][]float64, len(nn.hiddenNodes))
	for j := range nn.hiddenNodes {
		// hidden node j's various weights to output nodes
		weights_HiddenOutput[j] = make([]float64, len(nn.outputNodes))
		for k := range nn.outputNodes {
			// weight between hidden node j and output node k
			weights_HiddenOutput[j][k] = nn.outputNodes[k].Parents[j].Weight
		}
	}

	// now, for each instance in the set
	for _, instance := range trainingSet {

		output := nn.calculateOutputForInstance(instance)

		outputWeightedErrors := make([]float64, len(nn.outputNodes))
		for k := range nn.outputNodes {
			outputError := float64(instance.ClassArray[k]) - output[k]
			// weighting the error by the derivative of the activaition function
			// accounts for cost function being more "sensitive" to weight updates
			weightedError := outputError * D_Sigmoid(nn.outputNodes[k].inputValue)
			outputWeightedErrors[k] = weightedError
		}

		hiddenWeightedErrors := make([]float64, len(nn.hiddenNodes))
		for j := range nn.hiddenNodes {
			layerWeightedErrorSum := 0.0
			for k := range nn.outputNodes {
				layerWeightedErrorSum += weights_HiddenOutput[j][k] * outputWeightedErrors[k]
			}
			hiddenWeightedErrors[j] = layerWeightedErrorSum * D_Sigmoid(nn.hiddenNodes[j].inputValue)
		}

		// matrices for weight updates :)
		deltaWeights_IH := make([][]float64, len(weights_InputHidden))
		for i := range deltaWeights_IH {
			deltaWeights_IH[i] = make([]float64, len(weights_InputHidden[0]))
			for j := range deltaWeights_IH[i] {
				deltaWeights_IH[i][j] = LEARN_RATE * nn.inputNodes[i].Output() * hiddenWeightedErrors[j]
			}
		}

		deltaWeights_HO := make([][]float64, len(weights_HiddenOutput))
		for j := range deltaWeights_HO {
			deltaWeights_HO[j] = make([]float64, len(weights_HiddenOutput[0]))
			for k := range deltaWeights_HO[j] {
				deltaWeights_HO[j][k] = LEARN_RATE * nn.hiddenNodes[j].Output() * outputWeightedErrors[k]
			}
		}

		// now apply the delta matrices to the weights matrices
		for i := range weights_InputHidden {
			for j := range weights_InputHidden[i] {
				weights_InputHidden[i][j] = weights_InputHidden[i][j] + deltaWeights_IH[i][j]
				nn.hiddenNodes[j].Parents[i].Weight = weights_InputHidden[i][j]
			}
		}

		for j := range weights_HiddenOutput {
			for k := range weights_HiddenOutput[j] {
				weights_HiddenOutput[j][k] = weights_HiddenOutput[j][k] + deltaWeights_HO[j][k]
				nn.outputNodes[k].Parents[j].Weight = weights_HiddenOutput[j][k]
			}
		}

	}


}

func (nn *NeuralNetwork) calculateOutputForInstance(instance utils.Instance) []float64 {

	output := make([]float64, len(nn.outputNodes))
	for i, attributeValue := range instance.Attributes {
		nn.inputNodes[i].SetInput(attributeValue)
	}
	for i := range nn.outputNodes {
		// this will trigger output/hidden nodes to calc if they haven't yet :^)
		output[i] = nn.outputNodes[i].Output()
	}

	return output
}

func (nn *NeuralNetwork) Classify(instance utils.Instance) int {
	nn.calculateOutputForInstance(instance)
	bestOutput := 0.0
	bestIndex := -1
	for k := range nn.outputNodes {
		if (nn.outputNodes[k].outputValue > bestOutput) {
			bestOutput = nn.outputNodes[k].outputValue
			bestIndex = k
		}
	}

	outputArray := make([]int, len(nn.outputNodes))
	outputArray[bestIndex] = 1
	return bestIndex
}

func (nn *NeuralNetwork) Test(testSet []utils.Instance) {
	correct := 0
	for _, instance := range testSet {
		prediction := nn.Classify(instance)
		var actual int
		for c, val := range instance.ClassArray {
			if val == 1 { actual = c }
		}
		if prediction == actual {
			correct++
		}
		fmt.Printf("Predicted %d, actual %d\n", prediction, actual)


	}
	fmt.Printf("Accuracy: %f\n", float64(correct) / float64(len(testSet)))
}

func PrintNeuralNetwork(nn *NeuralNetwork) {
	fmt.Println("----- INPUT -----")
	for i := range nn.inputNodes {
		fmt.Printf("%d, ", int(nn.inputNodes[i].outputValue))
	}
	fmt.Println("\n----- HIDDEN -----")
	for i := range nn.hiddenNodes {
		fmt.Printf("%p\n", &nn.hiddenNodes[i])
		fmt.Printf("%d: input %f, output %f\n", i, nn.hiddenNodes[i].inputValue, nn.hiddenNodes[i].outputValue)
	}

	fmt.Println("----- OUTPUT -----")
	for i := range nn.outputNodes {
		fmt.Printf("%p\n", &nn.hiddenNodes[i])
		fmt.Printf("%d: input %f, output %f\n", i, nn.hiddenNodes[i].inputValue, nn.hiddenNodes[i].outputValue)
	}
}
