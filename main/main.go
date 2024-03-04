package main

import(
	"fmt"
	"net/http"
	"WebServer/src"
	"strconv"
)

func main() {
	
	src.LogConfig(Config)

	server:=http.Server{
		Addr:src.Config.Address+strconv.Itoa(src.Config.Port)
		Handler:&src.ProcessStr{}
	}
	src.LoadWeb()
	src.Init() //初始化处理函数

	server.ListenAndServer()
}