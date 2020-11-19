package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"

	"chat/server"

	"github.com/fatih/color"
)

var ccS = color.New(color.FgGreen, color.Bold)
var ccF = color.New(color.FgRed, color.Bold)

// go run ./chatServer.go -port 8000
func main() {
	port := flag.Int("port", 7777, "设置服务端口")
	flag.Parse()
	fmt.Printf("服务器正在[:%s]启动", strconv.Itoa(*port))
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa((*port)))
	if err != nil {
		ccF.Println("服务器启动失败：", err)
		return
	}
	//启动一个新携程监听信息
	go server.TransmitMsg()
	ccS.Println("已成功启动，正等待客户连接……")
	for {
		//不听循环1000
		time.Sleep(20 * time.Millisecond)
		conn, err := listener.Accept()
		if err != nil {
			ccF.Println("接受客户端链接失败", err)
			return
		}
		ccS.Printf("[%v]的客户端已经连接成功\n", conn.RemoteAddr())
		//并发聊天，一个客户端一个协程
		go server.HandleConnect(conn)
	}

}
