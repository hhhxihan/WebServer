package src

import(
	"net/http"
)

type ProcessStr struct{  //实现ServerHTTP
	action http.Handler
}

func (p *ProcessStr) ServerHTTP(w http.ResponseWriter,r *http.Request) {
	
	p.action = http.DefaultServeMux
	

	//设置响应头信息

	
	p.action.ServerHTTP(w,r)
	
}