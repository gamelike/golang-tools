package main

import (
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	var host string
	var port uint64
	flag.StringVar(&host, "h", "localhost", "nacos host default: localhost")
	flag.Uint64Var(&port, "p", 8848, "nacos port default: 8848")
	flag.Parse()
	fmt.Println(host)
	fmt.Println(port)
	_, configClient := NewClient(host, port)
	dataId := "ums-database.yaml"
	config, err := configClient.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   dataId,
		Group:    "DEFAULT_GROUP",
		PageNo:   0,
		PageSize: 10,
	})
	if err != nil {
		return
	}
	for _, item := range config.PageItems {
		fmt.Println("DataId" + item.DataId)
		fmt.Println("item: " + item.Content)
		content := make(map[string]interface{})
		err := yaml.Unmarshal([]byte(item.Content), content)
		if err != nil {
			panic(err)
		}
		fmt.Println("========================")
		defaultMap, ok := content["datasource"].(map[string]interface{})
		if !ok {
			defaultMap = make(map[string]interface{})
		}
		content["datasource"] = defaultMap
		datasourceMap, ok := defaultMap["default"].(map[string]interface{})
		if !ok {
			datasourceMap = make(map[string]interface{})
		}
		defaultMap["default"] = datasourceMap
		datasourceMap["url"] = "jdbc:postgresql:///postgres"
		bytes, err := yaml.Marshal(content)
		if err != nil {
			panic(err)
		}
		successStatus, err := configClient.PublishConfig(vo.ConfigParam{
			DataId:   dataId,
			Group:    "DEFAULT_GROUP",
			Content:  string(bytes),
			OnChange: nil,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(successStatus)
	}
}
func _(configClient config_client.IConfigClient) {
	filePath := "C:\\Users\\violet\\idea-projects\\bms-project\\bms-project\\release\\nacos\\DEFAULT_GROUP"
	dir, err := os.ReadDir(filePath)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		dataId := file.Name()
		content, err := os.ReadFile(filePath + "\\" + dataId)
		if err != nil {
			panic(err)
		}
		successStatus, err := configClient.PublishConfig(vo.ConfigParam{
			DataId:   dataId,
			Group:    "DEFAULT_GROUP",
			Content:  string(content),
			OnChange: nil,
		})
		if err != nil {
			panic(err)
		}
		fmt.Sprintln(successStatus)
	}
}
