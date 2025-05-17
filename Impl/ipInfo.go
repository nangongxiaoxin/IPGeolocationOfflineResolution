package Impl

import (
	"IP_Geolocation_Offline_Resolution/Model"
	_ "embed"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net"
	"strings"
	"time"
)

// 下边这玩意儿不是注释 >-<

//go:embed resources/ip/ip2region.xdb
var ip2regionxdb []byte

// IPInfoResult 是解析 IP 的返回结构体
type IPInfoResult = Model.IPInfoResult

func GetIpInfo(ip string) IPInfoResult {
	ip = strings.TrimSpace(ip)

	if net.ParseIP(ip) == nil {
		return IPInfoResult{
			IP:    ip,
			Error: "IP输入有误！",
		}
	}

	searcher, err := xdb.NewWithBuffer(ip2regionxdb)
	if err != nil {
		//fmt.Printf("failed to create searcher: %s\n", err.Error())
		return IPInfoResult{
			IP:    ip,
			Error: "创建searcher异常，请检查IP数据文件！",
		}

	}

	defer searcher.Close()
	tStart := time.Now()
	region, errSearch := searcher.SearchByStr(ip)
	if errSearch != nil {
		return IPInfoResult{
			IP:    ip,
			Error: "IP检索失败！",
		}

	}

	//耗时
	duration := time.Since(tStart)
	return IPInfoResult{
		IP:       ip,
		Region:   region,
		Duration: duration.String(),
	}

}
