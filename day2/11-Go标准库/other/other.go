package other

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"
)

func Demotemplate() {
	const templateText = `
	用户信息：
	=====================================
	用户名：{{.Name}}
	年龄：{{.Age}}
	邮箱：{{.Email}}
	{{if .Active}}状态： 活跃{{else}}状态： 非活跃{{end}}
	注册时间：{{.RegisteredAt.Format "2006-01-02 15:04:05"}}
	{{range .Hobbies}}爱好：{{.}}
	{{end}}
	`
	type User struct {
		Name         string
		Age          int
		Email        string
		Active       bool
		RegisteredAt time.Time
		Hobbies      []string
	}
	user := User{
		Name:         "张三",
		Age:          18,
		Email:        "zhangsan@example.com",
		Active:       true,
		RegisteredAt: time.Now(),
		Hobbies:      []string{"篮球", "足球", "游泳"},
	}

	tmpl, err := template.New("user").Parse(templateText)
	if err != nil {
		panic(err)
	}
	fmt.Println("模板渲染结果：")
	err = tmpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}

func DemoCompression() {
	originData := []byte("这是一段需要压缩的文本数据." + "重复内容" + strings.Repeat("重复", 10) +
		"结束")

	var compressed bytes.Buffer
	gzWriter := gzip.NewWriter(&compressed) // 创建gzip压缩对象
	_, err := gzWriter.Write(originData)    // 将原始数据写入压缩对象
	if err != nil {
		panic(err)
	}
	gzWriter.Close() // 关闭压缩对象，完成压缩

	fmt.Printf("原始大小：%d bytes\n", len(originData))
	fmt.Printf("压缩后大小：%d bytes\n", compressed.Len())
	fmt.Printf("压缩率:  %.1f%%\n", float64(compressed.Len())/float64(len(originData))*100)

	gzReader, err := gzip.NewReader(&compressed) // 创建gzip解压对象
	if err != nil {
		panic(err)
	}
	decompressed, err := io.ReadAll(gzReader) // 将压缩数据解压为原始数据
	if err != nil {
		panic(err)
	}
	fmt.Printf("解压后数据：%s\n", decompressed)

}

// DemoBuffIO 带缓冲的IO操作
func DemoBuffIO() {
	file, err := os.Create("test_buffer.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // 创建带缓冲的写入器
	for i := 0; i < 10; i++ {
		writer.WriteString(fmt.Sprintf("这是第%d行文本数据.\n", i+1))
	}
	writer.Flush() // 刷新缓冲区，将数据写入文件
	fmt.Println("写入完成")
	file.Close()
	file2, err := os.Open("test_buffer.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	reader := bufio.NewReader(file2) // 创建带缓冲的读取器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(line)
	}
}
