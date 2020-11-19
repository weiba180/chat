package server

import (
	"net"
)

func connRead(conn net.Conn, overTime chan bool) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		//与服务器通信的客户端用户名
		thisUser := onlineUsers[conn.RemoteAddr().String()].name
		if n == 0 {
			for _, v := range onlineUsers {
				if thisUser != "" {
					v.AddUser <- []byte("用户[" + thisUser + "]已退出\n")
				}
			}
			delete(onlineUsers, conn.RemoteAddr().String())
			return
		}
		if err != nil {
			ccF.Println("连接读取失败：", err)
			return
		}
		//处理消息内容
		var msg []byte
		if buf[0] != 10 {
			msg = append([]byte("["+thisUser+"]说>:"), buf[:n-1]...)
		} else {
			msg = nil
		}
		overTime <- true
		message <- msg
		LogToDb(string(msg), thisUser)
	}
}
