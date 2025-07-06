/*
  - @Author: zywOo 1372594487@qq.com
  - @Date: 2025-07-06 16:52:40

* @LastEditors: zywOo 1372594487@qq.com
* @LastEditTime: 2025-07-06 19:00:58
* @FilePath: \go-learning\day4\13-golang-IM-System\client\client.go
  - @Description: client文件 创建client实例 包含api：
    1 连接服务器
    2 发送消息
    3 接收消息
    4 处理用户输入
    5 优雅退出
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// ChatMode 聊天模式类型
type ChatMode int

const (
	PublicMode  ChatMode = iota // 公聊模式
	PrivateMode                 // 私聊模式
)

// String 返回模式的字符串表示
func (m ChatMode) String() string {
	switch m {
	case PublicMode:
		return "公聊"
	case PrivateMode:
		return "私聊"
	default:
		return "未知"
	}
}

// ChatSession 私聊会话信息
type ChatSession struct {
	TargetUser string
	Mode       ChatMode
}

// Client 客户端结构体
type Client struct {
	ServerIp    string
	ServerPort  int
	Name        string
	conn        net.Conn
	connected   bool
	mu          sync.Mutex
	config      *Config      // 添加配置字段
	verbose     bool         // 详细输出模式
	chatMode    ChatMode     // 当前聊天模式
	chatSession *ChatSession // 私聊会话信息
}

// Config 客户端配置
type Config struct {
	ServerHost string
	ServerPort int
	Username   string
	AutoMode   bool
	Verbose    bool
	Timeout    int
}

// SwitchToPublicMode 切换到公聊模式
func (client *Client) SwitchToPublicMode() {
	client.mu.Lock()
	defer client.mu.Unlock()

	client.chatMode = PublicMode
	client.chatSession = nil

	client.logInfo("已切换到公聊模式")
	fmt.Println("💬 现在是公聊模式，直接输入消息发送给所有人")
}

// SwitchToPrivateMode 切换到私聊模式
func (client *Client) SwitchToPrivateMode(targetUser string) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if targetUser == "" {
		fmt.Println("❌ 请指定私聊对象")
		return
	}

	client.chatMode = PrivateMode
	client.chatSession = &ChatSession{
		TargetUser: targetUser,
		Mode:       PrivateMode,
	}

	client.logInfo("已切换到私聊模式，对象: %s", targetUser)
	fmt.Printf("💭 现在是私聊模式，正在与 [%s] 私聊\n", targetUser)
	fmt.Println("   直接输入消息发送私聊，输入 /public 切换回公聊")
}

// GetCurrentMode 获取当前聊天模式信息
func (client *Client) GetCurrentMode() (ChatMode, string) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.chatMode == PrivateMode && client.chatSession != nil {
		return client.chatMode, client.chatSession.TargetUser
	}
	return client.chatMode, ""
}

// GetPrompt 获取输入提示符
func (client *Client) GetPrompt() string {
	mode, target := client.GetCurrentMode()

	switch mode {
	case PublicMode:
		return "💬> "
	case PrivateMode:
		return fmt.Sprintf("💭[%s]> ", target)
	default:
		return "> "
	}
}

// parseFlags 解析命令行参数
func parseFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.ServerHost, "host", "127.0.0.1", "服务器主机地址")
	flag.StringVar(&config.ServerHost, "h", "127.0.0.1", "服务器主机地址 (简写)")
	flag.IntVar(&config.ServerPort, "port", 8888, "服务器端口号")
	flag.IntVar(&config.ServerPort, "p", 8888, "服务器端口号 (简写)")
	flag.StringVar(&config.Username, "user", "", "预设用户名")
	flag.StringVar(&config.Username, "u", "", "预设用户名 (简写)")
	flag.BoolVar(&config.AutoMode, "auto", false, "自动模式，用于测试")
	flag.BoolVar(&config.Verbose, "verbose", false, "详细输出模式")
	flag.BoolVar(&config.Verbose, "v", false, "详细输出模式 (简写)")
	flag.IntVar(&config.Timeout, "timeout", 30, "连接超时时间(秒)")

	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "IM聊天客户端 v1.0\n\n")
		fmt.Fprintf(os.Stderr, "使用方法:\n")
		fmt.Fprintf(os.Stderr, "  %s [选项]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "选项:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n示例:\n")
		fmt.Fprintf(os.Stderr, "  %s -host 192.168.1.100 -port 9999\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -u Alice -v\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --auto  # 自动测试模式\n", os.Args[0])
	}

	flag.Parse()
	return config
}

// NewClient 创建新的客户端实例
func NewClient(config *Config) *Client {
	return &Client{
		ServerIp:    config.ServerHost,
		ServerPort:  config.ServerPort,
		Name:        "",
		connected:   false,
		config:      config,
		verbose:     config.Verbose,
		chatMode:    PublicMode, // 默认公聊模式
		chatSession: nil,
	}
}

// Connect 连接到服务器
func (client *Client) Connect() error {
	// 建立TCP连接
	address := net.JoinHostPort(client.ServerIp, fmt.Sprintf("%d", client.ServerPort))
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("连接服务器失败: %v", err)
	}

	client.mu.Lock()
	client.conn = conn
	client.connected = true
	client.mu.Unlock()

	fmt.Printf("成功连接到服务器 %s:%d\n", client.ServerIp, client.ServerPort)

	// 如果预设了用户名，自动发送重命名命令
	if client.config.Username != "" {
		client.logVerbose("自动设置用户名: %s", client.config.Username)
		time.Sleep(100 * time.Millisecond) // 等待连接稳定
		err := client.SendMessage("rename " + client.config.Username)
		if err != nil {
			client.logVerbose("设置用户名失败: %v", err)
		}
	}

	return nil
}

// SendMessage 发送消息到服务器
func (client *Client) SendMessage(message string) error {
	client.mu.Lock()
	defer client.mu.Unlock()

	if !client.connected || client.conn == nil {
		return fmt.Errorf("未连接到服务器")
	}

	_, err := client.conn.Write([]byte(message))
	if err != nil {
		client.connected = false
		return fmt.Errorf("发送消息失败: %v", err)
	}

	return nil
}

// ReceiveMessage 接收服务器消息
func (client *Client) ReceiveMessage() {
	defer func() {
		client.mu.Lock()
		client.connected = false
		client.mu.Unlock()
	}()

	buf := make([]byte, 4096)
	for {
		client.mu.Lock()
		if !client.connected {
			client.mu.Unlock()
			break
		}
		conn := client.conn
		client.mu.Unlock()

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("\n服务器连接断开")
			break
		}

		if n > 0 {
			message := string(buf[:n])
			// 在新行显示收到的消息，避免与输入提示符混淆
			fmt.Printf("\r%s\n> ", strings.TrimSpace(message))
		}
	}
}

// IsConnected 检查连接状态
func (client *Client) IsConnected() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return client.connected
}

// Disconnect 断开连接
func (client *Client) Disconnect() {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.connected && client.conn != nil {
		client.conn.Close()
		client.connected = false
		fmt.Println("已断开与服务器的连接")
	}
}

// handleSendError 统一处理发送错误
func (client *Client) handleSendError(err error, operation string) bool {
	if err != nil {
		fmt.Printf("%s失败: %v\n", operation, err)
		// 如果是连接断开错误，返回false表示应该退出
		if !client.IsConnected() {
			fmt.Println("连接已断开，程序将退出")
			return false
		}
	}
	return true
}

// ShowLocalHelp 显示客户端本地命令帮助
func (client *Client) ShowLocalHelp() {
	fmt.Println("\n=== 客户端本地命令帮助 ===")
	fmt.Println("基础命令:")
	fmt.Println("  /help, /h         - 显示此帮助信息")
	fmt.Println("  /status           - 显示连接状态和统计信息")
	fmt.Println("  /config           - 显示当前客户端配置")
	fmt.Println("  /clear            - 清屏")
	fmt.Println("  /verbose          - 切换详细输出模式")
	fmt.Println()
	fmt.Println("聊天模式:")
	fmt.Println("  /public, /pub     - 切换到公聊模式")
	fmt.Println("  /private <用户名> - 切换到私聊模式")
	fmt.Println("  /priv <用户名>    - 私聊模式简写")
	fmt.Println("  /pm <用户名>      - 私聊模式简写")
	fmt.Println("  /mode             - 显示当前聊天模式")
	fmt.Println("  /history, /hist   - 显示聊天历史")
	fmt.Println()
	fmt.Println("=== 服务器命令 (直接输入，无需/) ===")
	fmt.Println("  help              - 显示服务器命令帮助")
	fmt.Println("  rename 用户名     - 修改用户名")
	fmt.Println("  who               - 查询在线用户")
	fmt.Println("  exit, quit        - 退出程序")
	fmt.Println()
	fmt.Println("💡 提示:")
	fmt.Println("  • 在公聊模式下，直接输入消息发送给所有人")
	fmt.Println("  • 在私聊模式下，直接输入消息发送给指定用户")
	fmt.Println("  • 使用 /public 随时切换回公聊模式")
	fmt.Println("==============================")
}

// ShowStatus 显示客户端状态
func (client *Client) ShowStatus() {
	client.mu.Lock()
	defer client.mu.Unlock()

	mode, target := client.GetCurrentMode()

	fmt.Println("\n=== 客户端状态 ===")
	fmt.Printf("服务器地址: %s:%d\n", client.ServerIp, client.ServerPort)
	fmt.Printf("连接状态: %v\n", client.connected)
	fmt.Printf("当前用户名: %s\n", client.Name)
	fmt.Printf("聊天模式: %s\n", mode)
	if mode == PrivateMode {
		fmt.Printf("私聊对象: %s\n", target)
	}
	fmt.Printf("详细模式: %v\n", client.verbose)
	if client.config != nil {
		fmt.Printf("自动模式: %v\n", client.config.AutoMode)
		fmt.Printf("连接超时: %d秒\n", client.config.Timeout)
	}
	fmt.Println("==================")
}

// ProcessUserInput 处理用户输入
func (client *Client) ProcessUserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	// 显示命令提示
	client.showCommandTips()

	for client.IsConnected() {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// 处理本地客户端命令（以/开头）
		if strings.HasPrefix(input, "/") {
			client.handleLocalCommand(input)
			continue
		}

		// 处理退出命令
		if input == "exit" || input == "quit" {
			fmt.Println("正在退出...")
			break
		}

		// 根据当前模式处理消息
		client.handleMessage(input)

	}
}

// handleMessage 根据当前模式处理消息
func (client *Client) handleMessage(message string) {
	mode, target := client.GetCurrentMode()

	var command string
	switch mode {
	case PublicMode:
		// 公聊模式：直接发送消息
		command = message
		client.logVerbose("发送公聊消息: %s", message)

	case PrivateMode:
		// 私聊模式：构造私聊命令
		command = fmt.Sprintf("to %s %s", target, message)
		client.logVerbose("发送私聊消息给 %s: %s", target, message)
	}

	err := client.SendMessage(command)
	if !client.handleSendError(err, "发送消息") {
		return
	}
}

// showCommandTips 显示命令提示
func (client *Client) showCommandTips() {
	fmt.Println("命令提示:")
	fmt.Println("  /help       - 显示客户端命令帮助")
	fmt.Println("  /status     - 显示客户端状态")
	fmt.Println("  /config     - 显示客户端配置")
	fmt.Println("  help        - 显示服务器命令帮助")
	fmt.Println("  其他命令    - 发送到服务器")
	fmt.Println("  exit/quit   - 退出程序")
	fmt.Println()
}

// ShowConfig 显示客户端配置
func (client *Client) ShowConfig() {
	if client.config == nil {
		fmt.Println("配置信息不可用")
		return
	}

	fmt.Println("\n=== 客户端配置 ===")
	fmt.Printf("服务器主机: %s\n", client.config.ServerHost)
	fmt.Printf("服务器端口: %d\n", client.config.ServerPort)
	fmt.Printf("预设用户名: %s\n", client.config.Username)
	fmt.Printf("自动模式: %v\n", client.config.AutoMode)
	fmt.Printf("详细输出: %v\n", client.config.Verbose)
	fmt.Printf("连接超时: %d秒\n", client.config.Timeout)
	fmt.Println("==================")
}

// ClearScreen 清屏
func (client *Client) ClearScreen() {
	// Windows
	if os.PathSeparator == '\\' {
		fmt.Print("\033[H\033[2J")
	} else {
		// Unix/Linux/Mac
		fmt.Print("\033[2J\033[H")
	}
}

// ToggleVerbose 切换详细输出模式
func (client *Client) ToggleVerbose() {
	client.verbose = !client.verbose
	if client.config != nil {
		client.config.Verbose = client.verbose
	}
	fmt.Printf("详细输出模式: %v\n", client.verbose)
}

// handleLocalCommand 处理本地客户端命令
func (client *Client) handleLocalCommand(input string) {
	command := strings.TrimPrefix(input, "/")
	parts := strings.Fields(command)

	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "help", "h":
		client.ShowLocalHelp()
	case "status":
		client.ShowStatus()
	case "config":
		client.ShowConfig()
	case "clear":
		client.ClearScreen()
	case "verbose":
		client.ToggleVerbose()
	case "public", "pub":
		client.SwitchToPublicMode()
	case "private", "priv", "pm":
		if len(parts) < 2 {
			fmt.Println("❌ 用法: /private <用户名>")
			fmt.Println("   例如: /private Alice")
		} else {
			client.SwitchToPrivateMode(parts[1])
		}
	case "mode":
		client.ShowCurrentMode()
	case "history", "hist":
		client.ShowChatHistory()
	default:
		fmt.Printf("❌ 未知的本地命令: %s\n", parts[0])
		fmt.Println("输入 /help 查看可用的本地命令")
	}
}

// ShowCurrentMode 显示当前聊天模式
func (client *Client) ShowCurrentMode() {
	mode, target := client.GetCurrentMode()

	fmt.Println("\n=== 当前聊天模式 ===")
	fmt.Printf("模式: %s\n", mode)
	if mode == PrivateMode {
		fmt.Printf("私聊对象: %s\n", target)
	}
	fmt.Println("==================")
}

// ShowChatHistory 显示聊天历史（简化版）
func (client *Client) ShowChatHistory() {
	fmt.Println("\n=== 聊天记录 ===")
	fmt.Println("💡 提示: 聊天记录功能开发中...")
	fmt.Println("================")
}

// SetupSignalHandler 设置信号处理器，优雅退出
func (client *Client) SetupSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n\n收到退出信号，正在安全退出...")
		client.Disconnect()
		os.Exit(0)
	}()
}

// Run 启动客户端主程序
func (client *Client) Run() {
	// 设置信号处理
	client.SetupSignalHandler()

	// 连接服务器
	err := client.Connect()
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}

	// 延迟断开连接
	defer client.Disconnect()

	// 启动接收消息的goroutine
	go client.ReceiveMessage()

	// 等待连接稳定
	time.Sleep(100 * time.Millisecond)

	// 显示欢迎信息
	fmt.Println("\n=== 欢迎使用 IM 聊天客户端 ===")
	fmt.Println("输入 'help' 查看帮助信息")
	fmt.Println("输入 'exit' 或 'quit' 退出程序")
	fmt.Println("===============================")

	// 处理用户输入
	client.ProcessUserInput()

	fmt.Println("客户端已退出")
}

// logVerbose 详细模式日志输出
func (client *Client) logVerbose(format string, args ...interface{}) {
	if client.verbose {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// logInfo 信息日志输出
func (client *Client) logInfo(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format+"\n", args...)
}

// RunAutoMode 自动模式，用于测试
func (client *Client) RunAutoMode() {
	client.logInfo("启动自动测试模式")

	// 设置信号处理
	client.SetupSignalHandler()

	// 连接服务器
	err := client.Connect()
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}
	defer client.Disconnect()

	// 启动接收消息的goroutine
	go client.ReceiveMessage()
	time.Sleep(500 * time.Millisecond)

	// 自动执行测试命令
	testCommands := []string{
		"help",
		"who",
		"hello everyone, this is auto test",
		"rename AutoUser_" + strconv.Itoa(int(time.Now().Unix())),
		"who",
		"exit",
	}

	for i, cmd := range testCommands {
		client.logInfo("执行测试命令 %d/%d: %s", i+1, len(testCommands), cmd)

		if cmd == "exit" {
			break
		}

		err := client.SendMessage(cmd)
		if err != nil {
			client.logInfo("命令执行失败: %v", err)
			break
		}

		time.Sleep(2 * time.Second) // 等待响应
	}

	client.logInfo("自动测试完成")
}

// main 函数
func main() {
	// 解析命令行参数
	config := parseFlags()

	// 创建客户端实例
	client := NewClient(config)

	// 显示启动信息
	if config.Verbose {
		fmt.Printf("客户端配置:\n")
		fmt.Printf("  服务器: %s:%d\n", config.ServerHost, config.ServerPort)
		fmt.Printf("  用户名: %s\n", config.Username)
		fmt.Printf("  详细模式: %v\n", config.Verbose)
		fmt.Printf("  自动模式: %v\n", config.AutoMode)
		fmt.Printf("  超时时间: %d秒\n", config.Timeout)
		fmt.Println()
	}

	// 根据模式运行
	if config.AutoMode {
		client.RunAutoMode()
	} else {
		client.Run()
	}
}
