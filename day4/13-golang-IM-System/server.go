/*
* @Author: zywOo 1372594487@qq.com
* @Date: 2025-07-04 13:58:38
  - @LastEditors: zywOo 1372594487@qq.com
  - @LastEditTime: 2025-07-04 19:07:36
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
	"sync"
)

// 定义一个Server结构体，包含Ip、Port、OnlineMap和Message四个字段
type Server struct {
	Ip   string
	Port int
	// 在服务器端维护一个map，用来保存用户id和用户对象
	OnlineMap map[string]*User
	//读写锁
	mapLock sync.RWMutex
	// 创建一个channel，用来保存上线消息
	Message chan string
}

/* API
 */
// 创建一个Server的API
func NewServer(ip string, port int) *Server {
	// 创建一个Server对象，并初始化Ip、Port、OnlineMap和Message字段
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		// 创建一个channel消息广播，用来保存上线消息
		Message: make(chan string),
	}
	return server
}
func (this *Server) DoHandler(conn net.Conn) {
	// fmt.Println("连接成功")
	// fmt.Println(conn.RemoteAddr().String())
	user := NewUser(conn, this)
	// 广播用户上线
	user.Online()

	// 当前handler阻塞
	select {}
}

// 广播功能
func (this *Server) BroadCast(user *User, msg string) {
	// 格式化消息

	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg
	// 将消息发送到服务器的Message channel
	this.Message <- sendMsg
}

// 监听功能
func (this *Server) ListenMessage() {
	// 循环监听Message channel
	for {
		msg := <-this.Message
		// 将消息发送给全部用户
		this.mapLock.Lock()
		for username, cli := range this.OnlineMap {
			fmt.Printf("向用户 %s 发送消息\n :", username) //向用户 zhangsan 发送消息
			fmt.Println("服务器端显示:", msg)             // 调试用
			// 确保消息是UTF-8编码
			// sendMsgBytes := []byte(msg)
			// fmt.Printf("发送字节: %v\n", sendMsgBytes) // 调试用
			cli.C <- msg
		}
		// for _, cli := range this.OnlineMap {
		// 	cli.C <- msg
		// }
		this.mapLock.Unlock()
	}
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
	//select

	//close
	defer listener.Close()
	//启动监听消息的goroutine
	go this.ListenMessage()

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

}
