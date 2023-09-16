package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type CMD struct {
	Name string `json:"name"`
}

func main() {

	str := "even.song@lianwei.com.cn,601211150@qq.com"
	arr := strings.Split(str, ",")
	var buffer bytes.Buffer
	for _, v := range arr {
		fmt.Println(v)
		buffer.WriteString(fmt.Sprintf("\"%s\",", v))
	}
	list := buffer.String()
	cleanedEmail := strings.TrimRight(list, ",")
	fmt.Println(cleanedEmail)
	//cmd := `callor --service=gitlab --action=createProject --body={"architect":"architect","projectName":"projectName","username":"username","groupName":"groupName"}`
	//body := parseCommand(cmd)
	//if body == nil {
	//	fmt.Println("cmd参数错误:", cmd)
	//}
	//fmt.Println(string(body))
	//
	//ss := "abc"
	//sss := string(ss)
	//fmt.Println(sss)
}

func parseCommand(cmd string) (body []byte) {

	//命令行示例：cmd := `callor --service=gitlab --action=createProject --body={"architect":"architect","projectName":"projectName","username":"username","groupName":"groupName"}`
	service := flag.String("service", "default", "service")
	action := flag.String("action", "default", "action")
	data := flag.String("body", "default", "body")
	flags := strings.Split(cmd, " ")
	os.Args = os.Args[:0]
	for _, value := range flags {
		os.Args = append(os.Args, value)
	}
	flag.Parse()

	if *service == "gitlab" {
		switch *action {
		case "createProject":
			//
			project := Project{}
			json.Unmarshal([]byte(*data), &project)
			j, _ := json.Marshal(project)
			return j
		}
	}
	return nil
}

type Project struct {
	Architect    string `json:"architect"`
	ProjectName  string `json:"projectName"`
	Username     string `json:"username"`
	GroupName    string `json:"groupName"`
	SubgroupName string `json:"subgroupName"`
}
