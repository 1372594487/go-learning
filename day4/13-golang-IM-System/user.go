/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-04 16:10:49
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-04 18:38:41
 * @FilePath: \go-learning\day4\13-golang-IM-System\user.go
 * @Description:
 *
 */
package main

import "net"

type User struct {
	Name string
	Addr string
	// 通道，用于发送和接收消息
	C chan string
	// 连接，用于网络通信
	conn net.Conn
	// 服务器引用，用于访问服务器功能
	server *Server
}

/* API */
// 用户上线
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前User channel消息的goroutine
	go user.ListenMessage()
	return user
}

// 监听当前User channel的方法，一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线方法
func (this *User) Online() {
	// 加锁保护OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线方法
func (this *User) Offline() {
	// 加锁保护OnlineMap
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播用户下线消息
	this.server.BroadCast(this, "已下线")
}

// 处理用户消息
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}
