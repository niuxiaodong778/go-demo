package main

import (
	"fmt"
	"log"
	"net/rpc"
)
const HelloServiceNameClient = "path/hello/pkt"

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network,addr string) (*HelloServiceClient,error) {
	c,err := rpc.Dial(network,addr)
	if err != nil {
		log.Fatal(err)
	}
	return &HelloServiceClient{Client:c},err
}
func (c *HelloServiceClient)Hello(req string,reply * string) error {
	return c.Client.Call(HelloServiceNameClient+".Hello",req,reply)
}

func (c * HelloServiceClient)Hello2(req string,reply * string)error  {
	return c.Client.Call(HelloServiceNameClient+".Hello2",req,reply)
}
func main() {
	client,err := DialHelloService("tcp",":1234")
	if err != nil {
		log.Println(err)
	}

	var reply string
	var reply2 string
	err = client.Hello("xiaodong",&reply)
	if err != nil {
		log.Fatal(err)
	}
	client.Hello2("ruiteng",&reply2)
	fmt.Println(reply,reply2)
}
