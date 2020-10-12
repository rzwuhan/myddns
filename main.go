package main

import (
	"myddns/ali"
	"myddns/log"
	"myddns/tencent"
	"myddns/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	aliAccessKeyId      = ""
	aliAccessKeySecrete = ""
	aliDomain           = "" // example.com
	aliSubDomain        = "" // pi

	tencentAccessKeyId      = ""
	tencentAccessKeySecrete = ""
	tencentDomain           = "" // example.com
	tencentSubDomain        = "" // pi
)

func waitForTerminate() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.D("signal", s, "received")
			signal.Stop(c)
			close(c)
		default:
			log.F("unhandled signal:", s)
		}
	}
}

func updateLoop() {
	prevIp := ""
	for {
		// 获取公网IP
		ip := utils.GetIp()
		if len(ip) == 0 {
			// 获取公网IP失败, 1m后重试
			log.E("get outbound ip failed")
			time.Sleep(1 * time.Minute)
			continue
		}

		if ip != prevIp {
			// 更新阿里云域名
			if err := ali.UpdateDomainRecord(aliSubDomain, aliDomain, ip); err != nil {
				log.E("update domain record failed:", err)
			} else {
				log.I("update ali domain record with outbound ip success")
			}

			// 更新腾讯云域名
			if err := tencent.UpdateDomainRecord(tencentSubDomain, tencentDomain, ip); err != nil {
				log.E("update domain record failed:", err)
			} else {
				log.I("update tencent domain record with outbound ip success")
			}
			prevIp = ip
		}
		// 每隔1h检查一次
		time.Sleep(1 * time.Hour)
	}
}

func main() {
	// 配置API访问密钥
	ali.UseAccessKey(aliAccessKeyId, aliAccessKeySecrete)
	tencent.UseAccessKey(tencentAccessKeyId, tencentAccessKeySecrete)

	go updateLoop()

	waitForTerminate()
}
