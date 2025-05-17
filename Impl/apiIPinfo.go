package Impl

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var server *http.Server

func StartLocalAPI() {
	mux := http.NewServeMux()
	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		ip := r.URL.Query().Get("ip")

		// 设置响应头为 JSON
		w.Header().Set("Content-Type", "application/json")

		//校验
		ip = strings.TrimSpace(ip)

		if net.ParseIP(ip) == nil {
			err := json.NewEncoder(w).Encode(
				IPInfoResult{
					IP:    ip,
					Error: "IP输入有误！",
				})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("异常")
			return
		}

		result := GetIpInfo(ip)

		_ = json.NewEncoder(w).Encode(result)
	})

	server = &http.Server{
		Addr:    ":45555",
		Handler: mux,
	}

	go func() {
		log.Println("本地 API 接口已启动，监听端口 45555")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("API 启动失败:", err)
		}
	}()
}

func StopLocalAPI() {
	if server != nil {
		_ = server.Close()
		log.Println("本地 API 接口已关闭")
	}
}
