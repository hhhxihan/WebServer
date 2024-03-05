package src

import(
	"fmt"
	"net/http"
)

type ProcessStr struct{  //实现ServerHTTP
	action http.Handler
}

func (p *ProcessStr) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Println("new connect");
	p.action=http.DefaultServeMux
	
	//设置响应头信息

	
	p.action.ServeHTTP(w,r)
	
}