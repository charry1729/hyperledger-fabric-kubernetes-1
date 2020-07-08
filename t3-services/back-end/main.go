package main

import (
	"github.com/charry1729/hyperledger-fabric-kubernetes-1/t3-services/back-end/server"
	 
)
func main()  {
	s := server.NewServer()
	if err := s.Init(8080);
	err!=nil{
		panic(err)
	}
	s.Start()
}