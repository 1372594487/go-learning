/*
  - @Author: zywOo 1372594487@qq.com

  - @Date: 2025-07-04 14:02:28

* @LastEditors: zywOo 1372594487@qq.com

* @LastEditTime: 2025-07-04 15:34:13

* @FilePath: \go-learning\day4\13-golang-IM-System\main.go

  - @Description: 主入口文件
    运行：go build -o server main.go server.go
    启动：./server
    nc 127.0.0.1 8888
    gitbash 没有nc(netcat)，需要下载
    下载预编译的nc.exe：
    下载地址：https://nmap.org/ncat/ 或 https://eternallybored.org/misc/netcat/
    放置位置：C:\Program Files\Git\usr\bin\nc.exe

    *
*/
package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
