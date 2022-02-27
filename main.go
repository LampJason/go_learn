package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world golang"))
	//读取当前系统的VERSION配置，并写入header
	key := "VERSION"
	os.Setenv(key, "v1")
	version := os.Getenv(key)
	w.Header().Set(key, version)

	//request的header写入response的header
	for k,v := range r.Header{
		for _,vv := range v{
			w.Header().Set(k,vv)
		}
	}

	//记录访问日志：客户端IP，HTTP返回码，输出到Server 标准输出
	//Fprintf 是往IO write里写，printf 是往标准输出写
	clientIp := getHeaderIp(r)
	log.Printf("clientIp %s",clientIp)
	log.Printf("response code %v", 200)
}

func getHeaderIp(r *http.Request) (ip string){
	ip = r.Header.Get("X-Forword-For")
	if ip != ""{
		return
	}
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return
	}
	//remoteAddr返回的 ip:port
	if ip = strings.Split(r.RemoteAddr,":")[0];ip != "" {
		return
	}
	return
}

//当不知道用指针的时候，就用指针
func main() {
	fmt.Println(22)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	//mux.HandleFunc("/healthz",healthz)
	if err:= http.ListenAndServe(":80", mux);err != nil{
		log.Fatalf("http err:%v", err)
	}
}

