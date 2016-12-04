package indexes 
import (
		"fmt"
		"strings"
)

const (
		LeafNode = 0
		IndexNode = 1
) 

type Node struct {
	Height int
	Keys []int 
	Children []interface{}
	NodeType int 
	NumObjects int 
}

func (n *Node)Init(size int){
	n.Keys = make([]int,size) 
	n.Children = make([]interface{},size) 
	n.NumObjects = 0 
}
func (n *Node)Print(){
	fmt.Printf("%s %+v\n", strings.Repeat("__",n.Height),n.Children)
	for _,child := range n.Children {
			switch v:=child.(type) {
			case Node: v.Print() 
		}
	}
}

func (n *Node)IsLeafNode() bool {
		return n.NodeType == LeafNode 
}

func (n *Node)IsIndexNode() bool {
		return n.NodeType == IndexNode 
}

func (n *Node)InsertData(node *Node,key int,data interface{}){
	if node.NumObjects == 0 {
			node.Keys[0] = key 
			node.Children[0] = data 
			node.NumObjects = node.NumObjects + 1
			return 
	}
		
	insertionPoint := 0  
	for i,_ := range node.Keys {
		if key < node.Keys[i]{
			insertionPoint = i
			break 
		}
		if i == node.NumObjects-1 {
			insertionPoint = i+1 
			break 
		}
	}
	copy(node.Keys[insertionPoint+1:],node.Keys[insertionPoint:])
	node.Keys[insertionPoint] = key 
	copy(node.Children[insertionPoint+1:],node.Children[insertionPoint:])
	node.Children[insertionPoint] = data 
	node.NumObjects = node.NumObjects + 1 
}



