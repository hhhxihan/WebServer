package main

import(
	"fmt"
	"net/http"
	"WebServer/src"
	"strconv"
)

func main() {
	
	src.LogConfig("configs/config.json")
	address:=src.Config.Address+strconv.Itoa(src.Config.Port)
	fmt.Println("listen address: ",address)
	server:=http.Server{
		Addr:address,
		Handler:&src.ProcessStr{},
	}
	src.LoadWeb("web/*.tmpl")
	src.Init() //初始化处理函数

	fmt.Println("start listen..")

	server.ListenAndServe()
}