package main

import (
	"IP_Geolocation_Offline_Resolution/Impl"
	"context"
	"fmt"
)

// IPInfoResult ip信息
//type IPInfoResult = Model.IPInfoResult

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name 没啥用，先留着 >O.o<
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetIpInfoService 获取ip信息
func (a *App) GetIpInfoService(ip string) string {
	IPSearchResult := Impl.GetIpInfo(ip)
	if IPSearchResult.Error != "" {
		return "错误：" + IPSearchResult.Error
	}
	return "【 " + IPSearchResult.Region + " 】" + " 耗时：" + IPSearchResult.Duration
}

// StartAPIService API启动
func (a *App) StartAPIService() {
	Impl.StartLocalAPI()
}

// StopAPIService API关闭
func (a *App) StopAPIService() {
	Impl.StopLocalAPI()
}
