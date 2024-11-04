package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func Search(configClient config_client.IConfigClient, dataId string) (string, error) {
	config, err := configClient.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   dataId,
		Group:    "DEFAULT_GROUP",
		PageNo:   0,
		PageSize: 10,
	})
	return config.PageItems[0].Content, err
}

func Update(configClient config_client.IConfigClient, dataId string, content []byte) (bool, error) {
	successStatus, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:   dataId,
		Group:    "DEFAULT_GROUP",
		Content:  string(content),
		OnChange: nil,
	})
	return successStatus, err
}
