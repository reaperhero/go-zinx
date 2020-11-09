package znet

import "go-zinx/ziface"

type Request struct {
	conn ziface.IConnection //已经和客户端建立好的 链接
	data []byte             //客户端请求的数据
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
