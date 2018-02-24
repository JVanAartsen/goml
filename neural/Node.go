package neural


const INPUT = 1
const BIAS_TO_HIDDEN = 2
const HIDDEN = 3
const BIAS_TO_OUTPUT = 4
const OUTPUT = 5

type Node struct {
	Type int
	Parents []NodeWeightPair // slice containing the parents including bias node
	inputValue float64
	outputValue float64
}

type NodeWeightPair struct {
	Node *Node
	Weight float64
}

func CreateNode(nodeType int) Node {

	var parents []NodeWeightPair
	var inputValue float64
	var outputValue float64
	//
	if nodeType == INPUT || nodeType == HIDDEN || nodeType == OUTPUT {
		parents = make([]NodeWeightPair, 0)
		inputValue = float64(-1)
		outputValue = float64(-1)
	} else { // bias nodes (and input)
		parents = nil
		inputValue = float64(1)
		outputValue = inputValue // no activation fn
	}

	return Node{nodeType, parents, inputValue, outputValue}
}

func (node *Node) SetInput(inputValue float64) {
	// should only be called for input nodes
	if node.Type != INPUT { panic("cannot call node.SetInput on anything other than Input nodes") }

	node.inputValue = inputValue
	// inputs have no activation, their input is their output
	node.outputValue = inputValue
	// reset hidden, output

}

func (node *Node) Output() float64 {
	// only hidden and output nodes need to sum their weighted inputs
	if node.Type == HIDDEN || node.Type == OUTPUT {
		weightedInputSum := 0.0
		for i := range node.Parents {
			weightedInputSum += node.Parents[i].Weight * node.Parents[i].Node.Output()
		}
		node.inputValue = weightedInputSum
		node.outputValue = Sigmoid(node.inputValue)
	} else {
		node.outputValue = node.inputValue
	}
	return node.outputValue
}
