package main
import (
    "fmt"
    "net/http" 
    "log"
)
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的 
	var bodyreader = r.Body;
	var bodt []byte;
	bodyreader.Read(bodt);
	fmt.Fprintf(w,	string(bodt)); 
	fmt.Printf(	string(bodt))
} 
func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}