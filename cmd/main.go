package main

import (
	"openlogd.io/openlogd/pkg/indexes" 
)
func main(){
	index := indexes.BTree{}  
	index.Init(5) 
	index.Print()
	index.Add(10,10) 
	index.Print()
	index.Add(3,3) 
	index.Print()
	index.Add(2,2) 
	index.Print()
	index.Add(20,20) 
	index.Print()
	index.Add(7,7) 
	index.Print()
	index.Add(11,11) 
	index.Print()
}	
