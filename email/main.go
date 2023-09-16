package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.163.com", 25, "18301300310@163.com", "TVZSOOHMFVCDHUVA")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", "18301300310@163.com")
	m.SetHeader("To", "even.song@lianwei.com.cn")
	m.SetHeader("Cc", "601211150@qq.com")
	m.SetHeader("Subject", "Amway电商基础设施告警通知!")
	body := fmt.Sprintf("<p>告警名称: k8s pod status failed</p>" +
		"<p>故障等级: P2</p>" +
		"<p>故障类型: app-component</p>" +
		"<p>故障时间: 2023-05-08T09:36:43.229Z</p>" +
		"<p>故障定位: seata-ha-server-567988d6df-fmgzc</p>" +
		"<p>所属业务: env=aliyum-sre,service=kube-state-metrics</p>" +
		"故障详情: K8S pod seata/seata-ha-server-567988d6df-fmgzc status failed")
	m.SetBody("text/html", body)
	// Send emails using d.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
