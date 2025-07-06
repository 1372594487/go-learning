# go-learning
I want to transfer to become a Golang engineer!


# Go IM 即时通讯系统

一个基于 Go 语言开发的即时通讯系统，支持多用户在线聊天、私聊、用户管理等功能。

## 📋 项目概述

这是一个完整的 IM 聊天系统，包含服务端和客户端，实现了以下核心功能：

### 🌟 主要特性

- **多用户在线聊天** - 支持多人同时在线交流
- **私聊功能** - 用户间一对一私密聊天
- **智能聊天模式** - 公聊/私聊模式无缝切换
- **用户管理** - 用户重命名、在线用户查询
- **智能客户端** - 功能丰富的命令行客户端
- **超时管理** - 自动清理闲置用户
- **并发安全** - 基于 goroutine 的高并发处理
- **优雅退出** - 支持信号处理和安全断连

### 📁 项目结构

```
go-learning/
├── README.md                 # 项目说明文档
├── day1/                     # Go 基础示例
│   ├── 01-var/              # Go 变量声明
│   ├── 02-const_iota/       # Go 常量以及iota
│   ├── 03-func/             # Go 函数声明
│   ├── 04-init/             # Go 匿名以及别名导包方式
│   └── 05-pointer/          # Go 语言中的指针
├── day2/                     # Go 进阶示例
│   ├── 01-struct/           # Go 结构体
│   ├── 02-slice/            # Go 切片
│   ├── 03-map/              # Go 映射
│   ├── 04-channel/          # Go 通道
│   ├── 05-goroutine/        # Go 协程
│   └── 06-select/           # Go 选择器
├── day3/                     # Go 高级示例
│   ├── 01-interface/        # Go 接口
│   ├── 02-error/            # Go 错误处理
│   ├── 03-context/          # Go 上下文
│   ├── 04-sync/             # Go 同步原语
│   └── 05-web/              # Go Web 编程
└── day4/                    # 实战项目
    └── 13-golang-IM-System/    # IM 即时通讯系统
        ├── go.mod               # 项目模块配置
        ├── main.go              # 服务器启动入口
        ├── server/              # 服务器核心代码
        │   ├── server.go        # 服务器主要逻辑
        │   └── user.go          # 用户管理模块  
        └── client/              # 客户端代码
            └── client.go        # 智能客户端程序

```
### Go 基础示例

项目中包含了 Go 语言学习示例：

- **day1/01-var/** - Go 变量声明
- **day1/02-const_iota/** - Go 常量以及iota
- **day1/03-func/** - Go 函数声明
- **day1/04-init/** - Go 匿名以及别名导包方式
- **day1/05-pointer/** - Go 语言中的指针

- **day2/01-struct/** - Go 结构体
- **day2/02-slice/** - Go 切片
- **day2/03-map/** - Go 映射
- **day2/04-channel/** - Go 通道
- **day2/05-goroutine/** - Go 协程
- **day2/06-select/** - Go 选择器

- **day3/01-interface/** - Go 接口
- **day3/02-error/** - Go 错误处理
- **day3/03-context/** - Go 上下文
- **day3/04-sync/** - Go 同步原语
- **day3/05-web/** - Go Web 编程

## 🚀 快速开始

### 前置条件

- Go 1.16 或更高版本
- 支持 TCP 连接的网络环境

### 初始化项目

```bash
# 进入项目目录
cd day4/13-golang-IM-System

# 初始化模块（如果还没有 go.mod）
go mod init golang-im-system

# 整理依赖
go mod tidy
```

### 运行服务器

```bash
# 在项目根目录
cd day4/13-golang-IM-System

# 启动服务器
go run main.go
```

服务器将在 `127.0.0.1:8888` 启动，输出信息：
```
Server 127.0.0.1:8888 is starting
```

### 运行客户端

#### 方式一：使用智能客户端（推荐）

```bash
# 进入客户端目录
cd day4/13-golang-IM-System/client

# 基本连接
go run client.go

# 指定服务器地址和端口
go run client.go -host 192.168.1.100 -port 9999

# 预设用户名并启用详细模式
go run client.go -user Alice -verbose

# 自动测试模式
go run client.go --auto

# 查看帮助
go run client.go -help
```

#### 方式二：使用 netcat（简单测试）

```bash
# Windows CMD (需要先设置编码)
chcp 65001
nc 127.0.0.1 8888

# Git Bash 或 Linux/Mac
nc 127.0.0.1 8888
```

### 客户端命令参数

| 参数 | 简写 | 说明 | 默认值 |
|------|------|------|--------|
| `-host` | `-h` | 服务器主机地址 | 127.0.0.1 |
| `-port` | `-p` | 服务器端口号 | 8888 |
| `-user` | `-u` | 预设用户名 | 无 |
| `-verbose` | `-v` | 详细输出模式 | false |
| `-auto` | | 自动测试模式 | false |
| `-timeout` | | 连接超时时间(秒) | 30 |
| `-help` | | 显示帮助信息 | |

## 💬 使用说明

### 客户端本地命令（以 / 开头）

| 命令 | 说明 | 示例 |
|------|------|------|
| `/help` | 显示客户端命令帮助 | `/help` |
| `/status` | 显示连接状态和统计信息 | `/status` |
| `/config` | 显示当前客户端配置 | `/config` |
| `/clear` | 清屏 | `/clear` |
| `/verbose` | 切换详细输出模式 | `/verbose` |
| `/public` | 切换到公聊模式 | `/public` |
| `/private <用户名>` | 切换到私聊模式 | `/private Alice` |
| `/pm <用户名>` | 私聊模式简写 | `/pm Bob` |
| `/mode` | 显示当前聊天模式 | `/mode` |

### 服务器命令（直接输入）

| 命令 | 说明 | 示例 |
|------|------|------|
| `help` | 显示服务器命令帮助 | `help` |
| `rename 用户名` | 修改用户名 | `rename Alice` |
| `who` | 查询在线用户 | `who` |
| `to 用户名 消息` | 发送私聊消息 | `to Bob 你好` |
| `exit` / `quit` | 退出程序 | `exit` |

### 聊天模式

#### 公聊模式（默认）
- 直接输入消息发送给所有在线用户
- 提示符：`💬> `
- 切换：使用 `/public` 命令

#### 私聊模式
- 使用 `/private 用户名` 切换到私聊模式
- 直接输入消息发送给指定用户
- 提示符：`💭[用户名]> `
- 使用 `/public` 切换回公聊模式

### 使用示例

```bash
# 启动客户端
go run client.go -user Alice

# 公聊模式
💬> 大家好，我是Alice

# 切换到私聊模式
💬> /private Bob
💭[Bob]> 你好Bob，这是私聊消息

# 切换回公聊模式
💭[Bob]> /public
💬> 我回到公聊了

# 查看在线用户
💬> who

# 修改用户名
💬> rename Alice_2023
```

## 🔧 构建说明

### 构建服务器

```bash
# 在项目根目录
cd day4/13-golang-IM-System

# 构建服务器可执行文件
go build -o server main.go

# 运行
./server
```

### 构建客户端

```bash
# 进入客户端目录
cd day4/13-golang-IM-System/client

# 构建客户端可执行文件
go build -o client client.go

# 运行
./client -host 127.0.0.1 -port 8888
```

### 交叉编译

```bash
# Windows 可执行文件
GOOS=windows GOARCH=amd64 go build -o server.exe main.go
GOOS=windows GOARCH=amd64 go build -o client.exe client/client.go

# Linux 可执行文件
GOOS=linux GOARCH=amd64 go build -o server main.go
GOOS=linux GOARCH=amd64 go build -o client client/client.go

# macOS 可执行文件
GOOS=darwin GOARCH=amd64 go build -o server main.go
GOOS=darwin GOARCH=amd64 go build -o client client/client.go
```

## 🛠️ 技术特性

### 服务器端
- **并发处理** - 每个用户连接使用独立 goroutine
- **线程安全** - 使用读写锁保护共享资源
- **消息广播** - 基于 channel 的消息分发机制
- **超时管理** - 自动清理 10 分钟无活动用户
- **优雅关闭** - 支持资源清理和连接管理

### 客户端
- **智能交互** - 支持公聊/私聊模式智能切换
- **命令补全** - 丰富的本地和服务器命令
- **状态管理** - 实时显示连接和聊天状态
- **错误处理** - 完善的错误提示和恢复机制
- **信号处理** - 支持 Ctrl+C 优雅退出
- **命令行参数** - 灵活的启动配置选项

## 📝 开发说明

### 运行流程

1. **服务器启动** - 监听 TCP 端口，等待客户端连接
2. **客户端连接** - 建立 TCP 连接，创建用户对象
3. **用户上线** - 用户信息加入在线列表，广播上线消息
4. **消息处理** - 实时处理用户输入，支持各种命令
5. **消息分发** - 通过 channel 机制分发消息给所有用户
6. **优雅退出** - 清理资源，广播下线消息

### 核心组件

- **Server** - 服务器主体，管理用户连接和消息分发
- **User** - 用户对象，处理单个用户的消息和状态
- **Client** - 智能客户端，提供丰富的交互功能

## 🐛 故障排除

### 常见问题

1. **连接失败**
   ```
   连接失败: dial tcp 127.0.0.1:8888: connect: connection refused
   ```
   - 检查服务器是否正常启动
   - 确认 IP 地址和端口号正确
   - 检查防火墙设置

2. **编译错误**
   ```
   undefined: xxx
   ```
   - 执行 `go mod tidy` 整理依赖
   - 检查 go.mod 文件是否正确
   - 确保所有文件在正确的包中

3. **中文乱码**
   - Windows CMD: 使用 `chcp 65001` 设置 UTF-8 编码
   - 推荐使用智能客户端而非 netcat

4. **用户超时**
   - 默认 10 分钟无活动会自动下线
   - 发送任意消息可重置活跃时间

## 📚 学习资源





## 📄 许可证

本项目基于 MIT 许可证开源。

## 🤝 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！

## 📞 联系方式

- 作者：zywOo
- 邮箱：1372594487@qq.com

---

⭐ 如果这个项目对你有帮助，请给它一个 Star！