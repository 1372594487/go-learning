/*
- @Author: zywOo 1372594487@qq.com

- @Date: 2025-07-04 14:02:28

  - @LastEditors: zywOo 1372594487@qq.com

  - @LastEditTime: 2025-07-05 14:51:51

  - @FilePath: \go-learning\day4\13-golang-IM-System\main.go

  - @Description: 主入口文件
    运行：go build -o server main.go server.go
    启动：./server
    客户端(cmd)：
    chcp 65001
    nc 127.0.0.1 8888
    gitbash 没有nc(netcat)，需要下载
    下载预编译的nc.exe：
    下载地址：https://nmap.org/ncat/ 或 https://eternallybored.org/misc/netcat/
    放置位置：C:\Program Files\Git\usr\bin\nc.exe
    注：Git Bash 和 CMD 对字符编码的默认处理方式不同。CMD 默认使用 GBK 编码，
    而 Git Bash 默认使用 UTF-8 编码。因此，在 CMD 中运行 nc.exe 时，可能会出现乱码问题。
    为了解决这个问题，您可以在 CMD 中使用 chcp 65001 命令将命令行窗口的字符编码设置为 UTF-8，然后再运行 nc.exe。
    或 GitBash 运行：
    curl telnet://127.0.0.1:8888

    curl 的缓冲机制：curl 在 telnet 模式下会缓存接收到的数据
    没有换行符刷新：服务端发送的消息可能没有正确的换行符来触发缓冲区刷新
    连接关闭时才显示：只有当连接断开时，所有缓存的数据才会被显示

    最佳选择是使用你已有的 nc 工具，因为：
    它专为网络通信设计
    支持实时双向通信
    没有 HTTP 相关的缓冲问题
    你的 Git Bash 已经包含了它
    curl 主要是为 HTTP 请求设计的，用于原始 TCP 连接时会有各种限制。

    流程：

    nc连接 → Accept → DoHandler → NewUser → Online → BroadCast → Message通道
    ↓
    TCP连接建立 ← conn.Write ← ListenMessage ← C通道 ← ListenMessage(服务器)

    0ms: nc发起连接
    1ms: 服务器Accept连接，创建User对象
    2ms: User上线，消息进入Message通道
    3ms: 服务器ListenMessage取出消息，分发给所有用户C通道
    4ms: 用户ListenMessage从C通道取出消息
    5ms: 通过TCP连接发送给您的nc客户端
    6ms: 您看到上线消息
    整个过程是异步并发的，多个goroutine同时工作，确保消息能够实时传递。

    *
*/
package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
