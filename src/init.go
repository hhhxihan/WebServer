package src

import(
	"net/http"
)
func Init() {
	http.HandlerFunc("/",Index)
	http.HandlerFunc("/login",Login)
	http.HandlerFunc("/signup",Signup)
}

func Index(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
	Templates.Lookup("Index.tmpl").Execute(w,nil)
	
}
}
func Login(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
		Templates.Lookup("Login.tmpl").Execute(w,nil)
		
	}
}
func Signup(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
	Templates.Lookup("Signup.tmpl").Execute(w,nil)
	
}
}