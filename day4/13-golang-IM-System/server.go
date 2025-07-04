/*
* @Author: zywOo 1372594487@qq.com
* @Date: 2025-07-04 13:58:38
  - @LastEditors: zywOo 1372594487@qq.com
  - @LastEditTime: 2025-07-04 15:38:22
  - @FilePath: \go-learning\day4\13-golang-IM-System\server.go
  - @Description: server文件 创建server实例 包含api：
    1 监听端口
    2 调用DoHandler方法处理客户端请求
    3 启动服务器

*
*/
package main

import (
	"fmt"
	"net"
)

// 定义一个Server结构体，包含Ip、Port、OnlineMap和Message四个字段
type Server struct {
	Ip   string
	Port int
	// 在服务器端维护一个map，用来保存用户id和用户对象
	// OnlineMap map[string]*User
	// 创建一个channel，用来保存上线消息
	// Message chan string
}

// 编写一个方法，用来处理用户上线
// func (this *Server) BroadCast(user *User, msg string) {
// 	// 将用户发送的消息格式化为：[用户地址]用户名：消息内容
// 	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
// }

// 创建一个Server的API
func NewServer(ip string, port int) *Server {
	// 创建一个Server对象，并初始化Ip、Port、OnlineMap和Message字段
	server := &Server{
		Ip:   ip,
		Port: port,
		// OnlineMap: make(map[string]*User),
		// Message:   make(chan string),
	}
	return server
}
func (this *Server) DoHandler(conn net.Conn) {
	fmt.Println("连接成功")
	fmt.Println(conn.RemoteAddr().String())
	// 创建一个User对象
	// user := NewUser(conn, this)
	// // 用户上线
	// user.Online()
	// // 监听用户消息
	// isLive := make(chan bool)
	// // 接收用户发送的消息
	// go func() {
	// 	buf := make([]byte, 4096)
	// 	for {
	// 		n, err := conn.Read(buf)
	// 		if n == 0 {
	// 			user.Offline()
	// 			return
	// 		}
	// 		if err != nil {
	// 			fmt.Println("conn read err:", err)
	// 			return
	// 		}
	// 		// 提取用户消息
	// 		msg := string(buf[:n])
	// 		// 用户针对msg进行处理
	// 		user.DoMessage(msg)
	// 		// 读取用户是否活跃的channel
	// 		isLive <- true
	// 	}
}

// 启动服务器
func (this *Server) Start() {
	// 打印服务器启动信息
	fmt.Printf("Server %s:%d is starting\n", this.Ip, this.Port)
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("listen err:", err)
	}

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}
		//do handler
		go this.DoHandler(conn)
	}

	//select

	//close
	defer listener.Close()
}
