package main

import(
	"net/http"
	"article"
)

func main() {
	// 启动服务
	server := http.Server{Addr: ":54110"}
	// 执行handle操作
	article.Articleok()
	server.ListenAndServe()
}