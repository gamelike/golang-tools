package main

import (
	"fmt"
	"go-demo/src/args"
	"go-demo/src/nacos"
	"gopkg.in/yaml.v3"
)

func main() {
	inputArgs := args.Parser()
	_, configClient := nacos.NewClient(inputArgs.Host, inputArgs.Port)
	if inputArgs.Path != "" {
		fmt.Println("skip config path.")
	}
	dataId := "ums-database.yaml"
	content, err := nacos.Search(configClient, dataId)
	if err != nil {
		return
	}
	fmt.Println("item: " + content)
	configMap := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(content), content)
	if err != nil {
		panic(err)
	}
	defaultMap, ok := configMap["datasource"].(map[string]interface{})
	if !ok {
		defaultMap = make(map[string]interface{})
	}
	configMap["datasource"] = defaultMap
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
	successStatus, err := nacos.Update(configClient, dataId, bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(successStatus)
}
