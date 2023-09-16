package main

import (
	"sync"
	"time"
)

type GatewayBody struct {
	AppName     string `json:"appName"`
	ApiName     string `json:"ApiName"`
	Path        string `json:"path"`
	HttpMethod  string `json:"httpMethod"`
	StatusCode  string `json:"statusCode"`
	Time        string `json:"time"`
	Timestamp   string `json:"@timestamp"`
	TraceID     string `json:"traceID"`
	ClusterID   string `json:"clusterID"`
	GatewayName string `json:"gatewayName"`
}

func workerpullSLS(id int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	// 模拟工作
	time.Sleep(time.Second)
	result := id * 2

	// 将结果发送到通道
	resultChan <- result
}

func workerPushES(id int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	// 模拟工作
	time.Sleep(time.Second)
	result := id * 2

	// 将结果发送到通道
	resultChan <- result
}

func main() {

}
