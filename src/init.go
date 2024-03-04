package src

import(
	"net/http"
)

func Init() {
	http.HandlerFuc("/",Index)
	http.HandlerFunc("/login",Login)
	http.HandlerFunc("/signup",Login)
}

func Index(w http.ResponseWriter,r *http.Request){

}
func Login(w http.ResponseWriter,r *http.Request){
	
}
func Signup(w http.ResponseWriter,r *http.Request){

}