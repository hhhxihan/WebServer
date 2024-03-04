package main

import(
	"fmt"
	"net/http"
	"src"
)

func main() {
	
	var address string;
	

	server:=http.Server{
		Addr:
		Handler:
	}

	src.Init() //初始化处理函数

	server.ListenAndServer()
}