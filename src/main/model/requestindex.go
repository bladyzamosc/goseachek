package model

import "fmt"

type RequestIndex struct {
	Text string
	What string
	Id   uint64
}

func (ri RequestIndex) PrintMe() {
	fmt.Println("Id : ", ri.Id)
	fmt.Println("Text : ", ri.Text)
	fmt.Println("What : ", ri.What)
}
