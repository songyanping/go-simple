package main

import (
	"fmt"
	"regexp"
)

func regex1() {
	body := "lastTime=0%E5%88%86%E9%92%9F&rawMetricName=code5XX_stage&expression=%24Value+%3E%3D+5&metricName=code5XX_stage&instanceName=null&signature=FpaZx5JmbQaKGAKM90Wj95oZh%2Fo%3D&transId=db5f2204-00ae-11ee-a24f-00163e205386&groupId=41263096&regionName=null&productGroupName=new-ecommerce-pd&metricProject=acs_apigateway&userId=1273560756088235&curValue=9&unit=%E4%B8%AA&alertName=API%E7%BD%91%E5%85%B3_%E8%BF%94%E5%9B%9E%E7%A0%815XX%E4%B8%AA%E6%95%B0&regionId=null&namespace=acs_apigateway&triggerLevel=INFO&alertState=ALERT&preTriggerLevel=null&ruleId=alarmRule_uuid_bf585f0992efef820bdab9920a57e178d8ae545f32f&dimensions=%7BapiStageName%3DRELEASE%2C+apiUid%3D42df3c02c3124cc3becd911e9495cc25%2C+region%3Dcn-shenzhen%2C+userId%3D1273560756088235%7D&timestamp=1685646149981"
	match, _ := regexp.MatchString("metricName=code1", body)
	fmt.Println(match)

}

func main() {

	regex1()
	text := "新电商MSE网关new-bff-h5三分钟内5XX错误请求增长率告警wec-mp-bff"

	reg := regexp.MustCompile("superapp|new-bff-h5|wec-mp-bff") // 查找连续的小写字母
	l := reg.FindAllString(text, -1)                            // 输出结果["ello" "o"]
	if l != nil {
		ch := l[0]
		fmt.Println(ch)
		switch ch {
		case "superapp":
			ch = "ca"
		case "new-bff-h5":
			fmt.Println("第一个case")
			ch = "ca"
		case "wec-mp-bff":
			fmt.Println("第二个case")
			ch = "ca"
		default:
			fmt.Println("Default")
			ch = "default"
		}
	}

}
