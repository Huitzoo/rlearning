package structdata

import (
	"fmt"
	"reinforcement/tools"
)

type Node struct {
	ID     string
	Coords tools.Coords
	Val    int
}

func NewNode(id string, val int, coords tools.Coords) *Node {
	return &Node{ID: id, Val: val, Coords: coords}
}

type Arist struct {
	Value      int
	Connection []*Node
	ID         string
}

func NewArist(N1, N2 *Node, val int, ID string) *Arist {
	return &Arist{Connection: []*Node{N1, N2}, Value: val, ID: ID}
}

type Graph struct {
	Arists      []*Arist
	ListOfNodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{ListOfNodes: make(map[string]*Node)}
}

func (g *Graph) CreateConnection(nodes [][]int, value int) {
	for i := 0; i < len(nodes); i += 2 {
		Node1 := g.CreateNode(nodes[i])
		Node2 := g.CreateNode(nodes[i+1])
		ID := fmt.Sprintf("%s%s", Node1.ID, Node2.ID)
		Arist := NewArist(Node1, Node2, value, ID)
		g.Arists = append(g.Arists, Arist)
	}
}

func (g *Graph) CreateNode(node []int) *Node {
	ID := fmt.Sprintf("%d", node)
	if _, ok := g.ListOfNodes[ID]; !ok {
		coords := tools.Coords{
			X: node[0],
			Y: node[1],
		}
		g.ListOfNodes[ID] = NewNode(ID, 0, coords)
	}
	return g.ListOfNodes[ID]
}
