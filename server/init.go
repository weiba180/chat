package server

import (
	"github.com/fatih/color"
)

//全局chan，处理各个服务端收到的信息
var message = make(chan []byte)

//储存在线用户信息
var onlineUsers = make(map[string]userInfo)
var ccS = color.New(color.FgGreen, color.Bold)
var ccF = color.New(color.FgRed, color.Bold)

type userInfo struct {
	name    string
	perC    chan []byte
	AddUser chan []byte //广播用户进入或退出
}
