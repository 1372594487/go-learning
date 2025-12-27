/*
 * @Author: 1372594487 1372594487@qq.com
 * @Date: 2025-12-27 02:20:30
 * @LastEditors: '1372594487
 * @LastEditTime: 2025-12-27 15:40:06
 * @Description: File description
 */
package main

import (
	// 命令行参数解析
	"flag"
	"fmt"
	"log"

	// 用户模块
	"net-http/config"
	"net-http/user"

	// HTTP服务器
	"net/http"
)

func main() {

	var configFile string
	var port int
	flag.StringVar(&configFile, "config", "config.json", "Path to configuration file")
	flag.IntVar(&port, "port", 8080, "Port to run the server on")
	flag.Parse()

	config, err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}

	handler := user.NewUserHandler(user.NewUserStore())
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Handle GET /users
			handler.GetUsers(w, r)
		case http.MethodPost:
			// Handle POST /users
			handler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", handler.GetUser)
	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("Server is running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
