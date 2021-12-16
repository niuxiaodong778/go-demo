package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/hello/pkt"
type HelloServiceInterface = interface {
	Hello(req string,reply * string)error
	Hello2(req string,reply *string)error
}

type HelloSVC struct {
}
func (h *HelloSVC) Hello(req string,reply *string) error{
	*reply = "hello " + req
	return nil
}

func (h *HelloSVC) Hello2(req string,reply *string) error{
	*reply = "hello2 " + req
	return nil
}

func RegHelloService(svc HelloServiceInterface)error  {
	return rpc.RegisterName(HelloServiceName,svc)
}
func main() {
	RegHelloService(new(HelloSVC))
	listen,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("listen tcp 1234 err: ",err)
	}
	log.Println("listen on 1234")
	for  {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("accept err: ",err)
		}
		go rpc.ServeConn(conn)
	}
}
