package data_structures

type graphType string

const (
	DirectedGraph         graphType = "DIRECTED"
	UndirectedGraph       graphType = "UNDIRECTED"
	DirectedWeightedGraph graphType = "DIRECTED_WEIGHTED"
)

type Node struct {
	Value    interface{}
	Children []*Node
}

type Edge struct {
	Weight interface{}
}

func GetSampleGraph(name graphType, nodeEle interface{}) *Node {
	switch name {
	case DirectedWeightedGraph:
		return SampleDirectedWeightedGraph(nodeEle)
	default:
		return nil
	}
}
func SampleDirectedWeightedGraph(nodeEle interface{}) *Node {
	root := &Node{
		Value:    nodeEle,
		Children: nil,
	}
	return root
}
func Main2() {

}
