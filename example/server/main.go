package main

import (
	"fmt"
	"go-zinx/ziface"
	"go-zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client : msgId = ", request.GetMsgID(), "data = ", string(request.GetData()))
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")

	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendBuffMsg(0, []byte("hello...hello...hello"))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("DoConneciotnLost is Called ... ")
}

func main() {
	s := znet.NewServer()
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
