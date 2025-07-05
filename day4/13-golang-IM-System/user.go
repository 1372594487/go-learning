/*
 * @Author: zywOo 1372594487@qq.com
 * @Date: 2025-07-04 16:10:49
 * @LastEditors: zywOo 1372594487@qq.com
 * @LastEditTime: 2025-07-05 20:10:52
 * @FilePath: \go-learning\day4\13-golang-IM-System\user.go
 * @Description:
 *
 */
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	Name string
	Addr string
	// 通道，用于发送和接收消息
	C chan string
	// 连接，用于网络通信
	conn net.Conn
	// 服务器引用，用于访问服务器功能
	server *Server
	// 最后活跃时间
	// 添加最后活跃时间
	LastActiveTime time.Time
}

/* API */
// 创建用户
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

// 更新用户活跃时间
func (this *User) UpdateActiveTime() {
	this.LastActiveTime = time.Now()
}

// 检查用户是否超时
func (this *User) IsTimeout(timeout time.Duration) bool {
	return time.Since(this.LastActiveTime) > timeout
}

// 监听当前User channel的方法，一旦有消息就发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		_, err := this.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("write to client failed:", err)
			// 可以选择关闭连接或进行其他错误处理
			// this.conn.Close()
			return
		}
	}
}

// 用户上线方法
func (this *User) Online() {

	// 初始化活跃时间
	this.UpdateActiveTime()

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

// 给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	_, err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("send msg to client failed:", err)
		return
	}
}

// 处理用户消息
func (this *User) DoMessage(msg string) {

	// 更新用户活跃时间
	this.UpdateActiveTime()

	msg = strings.TrimSpace(msg)

	if msg == "who" {
		// 查询当前在线用户
		this.server.QueryOnlineUsers(this)
	} else if this.parseRenameCommand(msg) {
		// 在 parseRenameCommand 中处理重命名
	} else if strings.HasPrefix(msg, "to ") {
		// 私聊功能
		this.PrivateChat(msg)
	} else if msg == "exit" {
		// 下线功能
		this.Offline()
		this.conn.Close()
	} else if msg == "help" {
		// 查看帮助信息
		this.server.Help(this)
	} else {
		// 广播消息
		this.server.BroadCast(this, msg)
	}

}

// 解析重命名命令，支持多格式
func (this *User) parseRenameCommand(msg string) bool {
	var newName string

	// 支持多种格式
	if strings.HasPrefix(msg, "rename ") {
		// rename newname
		parts := strings.Fields(msg)
		if len(parts) >= 2 {
			newName = parts[1]
		}
	} else if strings.Contains(msg, "rename:") {
		// rename:newname
		parts := strings.SplitN(msg, ":", 2)
		if len(parts) == 2 {
			newName = strings.TrimSpace(parts[1])
		}
	} else if strings.Contains(msg, "rename=") {
		// rename=newname
		parts := strings.SplitN(msg, "=", 2)
		if len(parts) == 2 {
			newName = strings.TrimSpace(parts[1])
		}
	} else {
		return false // 不是重命名命令
	}

	if newName != "" {
		this.server.RenameUser(this, newName)
		return true
	} else {
		this.SendMsg("用法: rename 新用户名 或 rename:新用户名 或 rename=新用户名")
		return true
	}
}

// 处理私聊消息
func (this *User) PrivateChat(msg string) {
	//解析私聊命令：to 用户名 消息内容
	parts := strings.SplitN(msg, " ", 3)
	if len(parts) < 3 {
		this.SendMsg("私聊格式错误，正确格式：to 用户名 消息内容")
		return
	}
	targetName := parts[1]
	message := parts[2]

	this.server.SendPrivateMessage(this, targetName, message)
}
