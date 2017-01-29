package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Usage : %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	//获取一个TCP地址信息,TCPAddr
	//解析地址和端口号
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	//创建一个TCP连接:TCPConn
	//建立链接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	//发送HTTP请求头
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	//获得返回数据
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(0)
	}
}
