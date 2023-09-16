package main

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type Alert struct {
	MetricName string
	Expression string
	Dimensions string
	AlertName  string
}

func main() {

	//body := "lastTime=2%E5%88%86%E9%92%9F\\u0026rawMetricName=code4XX_stage\\u0026expression=%24Value+%3E%3D+5\\u0026metricName=code4XX_stage\\u0026instanceName=null\\u0026signature=DxpsDBGaSPrVmSqA3j7xdQ3IVyo%3D\\u0026transId=2e0a0647-0064-11ee-b07f-00163e2e4b9a\\u0026groupId=41263096\\u0026regionName=null\\u0026productGroupName=new-ecommerce-pd\\u0026metricProject=acs_apigateway\\u0026userId=1273560756088235\\u0026curValue=2\\u0026unit=%E4%B8%AA\\u0026alertName=API%E7%BD%91%E5%85%B3_%E8%BF%94%E5%9B%9E%E7%A0%814XX%E4%B8%AA%E6%95%B0\\u0026regionId=null\\u0026namespace=acs_apigateway\\u0026triggerLevel=OK\\u0026alertState=OK\\u0026preTriggerLevel=INFO\\u0026ruleId=alarmRule_uuid_b13a2db9f1d577826640dbd79aad4178d8adf31238a\\u0026dimensions=%7BapiStageName%3DRELEASE%2C+apiUid%3Dc6fa0ad7950644b68437d8ae10f0b95f%2C+region%3Dcn-shenzhen%2C+userId%3D1273560756088235%7D\\u0026timestamp=1685614196421"
	body := "lastTime=0%E5%88%86%E9%92%9F&rawMetricName=code5XX_stage&expression=%24Value+%3E%3D+5&metricName=code5XX_stage&instanceName=null&signature=FpaZx5JmbQaKGAKM90Wj95oZh%2Fo%3D&transId=db5f2204-00ae-11ee-a24f-00163e205386&groupId=41263096&regionName=null&productGroupName=new-ecommerce-pd&metricProject=acs_apigateway&userId=1273560756088235&curValue=9&unit=%E4%B8%AA&alertName=API%E7%BD%91%E5%85%B3_%E8%BF%94%E5%9B%9E%E7%A0%815XX%E4%B8%AA%E6%95%B0&regionId=null&namespace=acs_apigateway&triggerLevel=INFO&alertState=ALERT&preTriggerLevel=null&ruleId=alarmRule_uuid_bf585f0992efef820bdab9920a57e178d8ae545f32f&dimensions=%7BapiStageName%3DRELEASE%2C+apiUid%3D42df3c02c3124cc3becd911e9495cc25%2C+region%3Dcn-shenzhen%2C+userId%3D1273560756088235%7D&timestamp=1685646149981"
	aliApGatewayMetricsDetails(body)
}

func aliApGatewayMetricsDetails(data string) (result string, err error) {
	urlFull := "http://localhost?" + data
	urlFullDecode, err := url.QueryUnescape(urlFull)
	if err != nil {
		log.Errorf("url QueryUnescape err:%s", err.Error())
		return "", err
	}
	urlParse, err := url.Parse(urlFullDecode)
	if err != nil {
		log.Errorf("url parse err:%s", err.Error())
		return "", err
	}

	body, err := url.ParseQuery(urlParse.RawQuery)
	if err != nil {
		log.Errorf("url parseQuery err:%s", err.Error())
		return "", err
	}
	var buffer bytes.Buffer
	alertName := body["alertName"][0]
	metricName := body["metricName"][0]
	expression := body["expression"][0]
	curValue := body["curValue"][0]
	timestamp := body["timestamp"][0]
	productGroupName := body["productGroupName"][0]
	triggerLevel := body["triggerLevel"][0]
	buffer.WriteString(fmt.Sprintf("\n监控指标: %s", metricName))
	buffer.WriteString(fmt.Sprintf("\n当前值(个): %s", curValue))
	buffer.WriteString(fmt.Sprintf("\n告警条件: %s", expression))
	buffer.WriteString(fmt.Sprintf("\n告警名称: %s", alertName))
	buffer.WriteString(fmt.Sprintf("\n告警时间: %s", timestamp))
	buffer.WriteString(fmt.Sprintf("\n应用分组: %s", productGroupName))
	buffer.WriteString(fmt.Sprintf("\n告警级别: %s", triggerLevel))

	result = buffer.String()

	fmt.Println(result)

	return result, nil
}

type AliApiGateway struct {
	AlertName        string `json:"alertName"`        //报警规则
	Timestamp        string `json:"timestamp"`        //报警时间
	ProductGroupName string `json:"productGroupName"` //应用分组
	ApiUid           string `json:"apiUid"`           //API UID
	Regin            string `json:"regin"`            //regin
	MetricName       string `json:"metricName"`       //监控指标
	Expression       string `json:"expression"`       //告警条件  $Value >= 5
	CurValue         string `json:"curValue"`         //当前值
	LastTime         string `json:"lastTime"`         //持续时间分钟
	TriggerLevel     string `json:"triggerLevel"`     //告警级别
	AlertState       string `json:"alertState"`       //告警状态  ALERT OR OK
}
