/*
* @Author: zywOo 1372594487@qq.com
* @Date: 2025-07-04 13:58:38
  - @LastEditors: zywOo 1372594487@qq.com
  - @LastEditTime: 2025-07-05 19:44:48
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
	"io"
	"net"
	"strings"
	"sync"
	"time"
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
	// 添加超时配置
	UserTimeout time.Duration // 用户超时时间
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
		Message:     make(chan string),
		UserTimeout: 10 * time.Minute, // 设置默认超时时间为10分钟
	}
	return server
}
func (this *Server) DoHandler(conn net.Conn) {
	// fmt.Println("连接成功")
	// fmt.Println(conn.RemoteAddr().String())
	user := NewUser(conn, this)
	// 广播用户上线
	user.Online()

	// 接受用户消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			// 提取用户消息
			msg := string(buf[:n])
			// 广播用户消息
			user.DoMessage(msg)
		}
	}()
}

// 帮助消息
func (this *Server) Help(user *User) {
	helpMsg := "当前支持以下命令：\n" +
		"rename 新用户名 或 rename:新用户名 或 rename=新用户名\n" +
		"who 查询在线用户\n" +
		"to 用户名 消息内容 - 发送私聊消息\n" +
		"help - 显示帮助信息\n" +
		"exit 退出程序\n" +
		"注意: 10分钟无活动将被自动下线\n"
	user.SendMsg(helpMsg)
}

// 查询在线用户
func (this *Server) QueryOnlineUsers(queryUser *User) {
	this.mapLock.RLock()
	defer this.mapLock.RUnlock()

	//构建在线用户列表信息
	var onlineMsg strings.Builder
	onlineMsg.WriteString("当前在线用户列表：\n")

	if len(this.OnlineMap) == 0 {
		onlineMsg.WriteString("当前无用户在线\n")
	} else {
		for _, user := range this.OnlineMap {
			onlineMsg.WriteString("[" + user.Addr + "]" + user.Name + "\n")
		}
	}

	onlineMsg.WriteString(fmt.Sprintf("总计: %d 人在线\n", len(this.OnlineMap)))
	onlineMsg.WriteString("==================\n")

	// 针对查询用户发送
	queryUser.SendMsg(onlineMsg.String())
}

// 超时检查和清理用户
func (this *Server) CheckTimeoutUsers() {
	ticker := time.NewTicker(1 * time.Minute) // 每分钟检查一次
	defer ticker.Stop()

	for range ticker.C {
		// 记录检查开始时间
		checkTime := time.Now()
		fmt.Printf("[%s] 开始检查超时用户...\n", checkTime.Format("2006-01-02 15:04:05"))

		this.mapLock.Lock()
		var timeoutUsers []*User
		totalUsers := len(this.OnlineMap) // 记录总用户数

		// 查找超时用户
		for name, user := range this.OnlineMap {
			if user.IsTimeout(this.UserTimeout) {
				timeoutUsers = append(timeoutUsers, user)
				delete(this.OnlineMap, name)
			}
		}
		remainingUsers := len(this.OnlineMap) // 剩余用户数
		this.mapLock.Unlock()

		// 打印检查结果
		if len(timeoutUsers) > 0 {
			fmt.Printf("[%s] 发现 %d 个超时用户，总用户数: %d -> %d\n",
				checkTime.Format("15:04:05"), len(timeoutUsers), totalUsers, remainingUsers)

			// 打印超时用户详情
			for _, user := range timeoutUsers {
				lastActiveTime := time.Since(user.LastActiveTime)
				fmt.Printf("  - 用户 [%s] 超时下线 (闲置时间: %v)\n",
					user.Name, lastActiveTime.Round(time.Second))
			}
		} else {
			fmt.Printf("[%s] 无超时用户，当前在线: %d 人\n",
				checkTime.Format("15:04:05"), totalUsers)
		}

		// 处理超时用户
		for _, user := range timeoutUsers {
			user.SendMsg("由于长时间未活动，您已被自动下线")
			user.conn.Close()

			// 广播用户超时下线消息
			this.BroadCast(user, "超时下线")
		}

		// 打印处理完成信息
		if len(timeoutUsers) > 0 {
			fmt.Printf("[%s] 超时用户处理完成\n", time.Now().Format("15:04:05"))
		}
	}
}

// 私聊功能
func (this *Server) SendPrivateMessage(sender *User, targetName string, message string) {
	// 检查目标用户是否存在
	this.mapLock.RLock()
	targetUser, exists := this.OnlineMap[targetName]
	this.mapLock.RUnlock()

	if !exists {
		sender.SendMsg("用户 " + targetName + " 不存在")
		return
	}

	if targetUser == sender {
		sender.SendMsg("不能给自己发送消息")
		return
	}
	//构造私聊消息格式
	privateMsg := "[私聊][" + sender.Addr + "]" + sender.Name + " 对你说: " + message

	//发送给目标用户
	targetUser.SendMsg(privateMsg)

	// 给发送者确认信息
	confirmMsg := "[私聊已送达][" + targetUser.Addr + "]" + targetName + ": " + message
	sender.SendMsg(confirmMsg)
}

// 修改用户名
func (this *Server) RenameUser(user *User, newName string) {
	this.mapLock.Lock()
	// 判断用户名是否已经存在
	if _, ok := this.OnlineMap[newName]; ok {
		user.SendMsg("用户名已存在，请重新输入")
	} else {
		// 删除旧用户名
		delete(this.OnlineMap, user.Name)
		// 修改用户名
		user.Name = newName
		this.OnlineMap[newName] = user
		user.SendMsg("用户名修改成功")
	}
	this.mapLock.Unlock()

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
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
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

	// 启动超时检查的goroutine
	go this.CheckTimeoutUsers()

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
