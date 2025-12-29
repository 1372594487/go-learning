/*
* @Author: 1372594487 1372594487@qq.com
* @Date: 2025-12-29 03:01:36
  - @LastEditors: '1372594487
  - @LastEditTime: 2025-12-29 22:19:29

* @Description: 反射语法/理论讲解
应用场景
1、框架开发
Web框架中的路由绑定和参数解析
各类注入的中间件
2、配置解析
JSON/YAML/TOML等配置文件的解析
根据字段标签动态映射
3、序列化/反序列化
JSON、XML编码解码
数据库ORM映射
4、依赖注入
自动创建对象实例
解析依赖关系
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 定义：程序在运行时检查自身结构的能力
// 核心：interface{} + `reflect`包

// 两大核心组件
// Type: 类型
// var x int = 10
// t := reflect.TypeOf(x) //获取x的类型

// Value: 值
// v := reflect.ValueOf(x) //获取x的值

// 关键概念
// Kind：基础类型（int、string、struct等）
// 可设置性：值能否被修改
// 类型断言：运行时类型检查

// 性能注意事项
// 反射比直接调用慢10-100倍
// 反射操作比直接操作慢，尽量避免使用反射

// 通用配置管理
type Config struct {
	AppName string `config:"app_name" default:"my_app"`
	Port    int    `config:"port" default:"8080"`
	Debug   bool   `config:"debug" default:"true"`
}

func (c *Config) UpdatePort(newPort int) (int, int) {
	old := c.Port
	c.Port = newPort
	return old, c.Port
}

func (c *Config) GettAppInfo() string {
	return fmt.Sprintf("App Name: %s, Port: %d, debug: %t", c.AppName, c.Port, c.Debug)
}

func (cm *ConfigManager) CallMethod(methodName string, args ...interface{}) []interface{} {
	v := reflect.ValueOf(cm.config)
	if v.Kind() != reflect.Ptr {
		v = v.Addr()
	}
	method := v.MethodByName(methodName)
	if !method.IsValid() {
		return []interface{}{fmt.Errorf("Method %s not found", methodName)}
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	out := method.Call(in)
	result := make([]interface{}, len(out))
	for i, out := range out {
		result[i] = out.Interface()
	}
	return result

}

type ConfigManager struct {
	config interface{}
}

func (cm *ConfigManager) LoadFromMap(data map[string]string) error {
	v := reflect.ValueOf(cm.config).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		configKey := fieldType.Tag.Get("config")
		defaultValue := fieldType.Tag.Get("default")
		value := data[configKey]
		if value == "" {
			value = defaultValue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
		case reflect.Int:
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			field.SetInt(int64(intValue))
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			field.SetBool(boolValue)

		}

	}
	return nil
}

type Person struct {
	Name string
	Age  int
}

func createInstance(t reflect.Type) interface{} {
	return &Person{Name: "John", Age: 30}
}

func main() {
	config := &Config{}
	configManager := &ConfigManager{config: config}
	configData := map[string]string{
		"app_name": "TestMyApp",
		"port":     "8088",
		"debug":    "false", // 这个字段不会被解析
	}
	err := configManager.LoadFromMap(configData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("app_name: %s, port: %d\n", config.AppName, config.Port)

	fmt.Println("----------------------------------------------------------")
	results := configManager.CallMethod("UpdatePort", 9090)
	for _, result := range results {
		fmt.Println(result)
	}
	results = configManager.CallMethod("GettAppInfo")
	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println("----------------------------------------------------------")
	personType := reflect.TypeOf(Person{})
	inst := createInstance(personType)
	if person, ok := inst.(*Person); ok {
		person.Name = "Tom"
		person.Age = 18
		fmt.Println("Person:", person)
	}
}
