package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"goproxy/config"
	"goproxy/server"

	"golang.org/x/net/proxy"
)



func main() {
	conf, err := config.LoadConfig("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("config: %+v\n", conf)

	// 代理地址
	proxyAddr := conf.ProxyAddr
	// 建立到代理服务器的连接
	proxyDialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("create proxy: %s\n", proxyAddr)

	for _, serverInfo := range conf.ServerConfig {
		lnAddr := serverInfo.ListenAddr
		targetAddr := serverInfo.TargetAddr

		server := server.NewServer(lnAddr, targetAddr, &proxyDialer)
		go server.Run()
	}
	
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("exit")
}





