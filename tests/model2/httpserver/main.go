package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("hello world golang"))
	//1。 读取当前系统的VERSION配置，并写入header
	key := "VERSION"
	os.Setenv(key, "v1")
	version := os.Getenv(key)
	w.Header().Set(key, version)

	//2. request的header写入response的header
	for k,v := range r.Header{
		for _,vv := range v{
			w.Header().Set(k,vv)
		}
	}

	//2. 记录访问日志：客户端IP，HTTP返回码，输出到Server 标准输出
	//Fprintf 是往IO write里写，printf 是往标准输出写
	//3. 客户端真实ip地址获取
	clientIp := getHeaderIp(r)
	log.Printf("clientIp %s",clientIp)
	log.Printf("response code %v", 200)
}

//健康检查
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

//客户端ip获取方法
func getHeaderIp(r *http.Request) (ip string){
	//支持从nginx反向代理获取来源真实ip
	ip = strings.TrimSpace(strings.Split(r.Header.Get("X-Forword-For"),",")[0])
	if ip != ""{
		return
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-IP"))
	if ip != "" {
		return
	}
	//remoteAddr返回的 ip:port
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr));err == nil {
		return ip
	}
	return
}

//当不知道用指针的时候，就用指针
func main() {
	fmt.Println(22)
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz",healthz)
	if err:= http.ListenAndServe(":9999", mux);err != nil{
		log.Fatalf("http server err:%s", err.Error())
	}
}

