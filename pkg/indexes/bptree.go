package indexes
import (
		"fmt"
)

type BTree struct {
	Root *Node
	PageSize  int 
}


func (b *BTree)Init(pageSize int){
	b.PageSize=pageSize
	b.Root = &Node {  Height:0, NodeType:LeafNode } 
	b.Root.Init(b.PageSize) 
}

func (b *BTree)Add(key int, data interface{}){

	b.Insert(b.Root,key,data)
	

}

func (b *BTree)Insert(node *Node,key int, data interface{}){
	if node.IsLeafNode() {
			if b.IsFull(node)==false{
				node.InsertData(node,key,data)
			}
	}
}

func (b *BTree)IsFull(node *Node) bool {
	return b.PageSize == node.NumObjects 
} 

func (b *BTree) Print(){
	b.Root.Print() 
}
func Moo(){
	fmt.Println("moo mare")
} 




