package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewClient(ip string, port uint64) (naming_client.INamingClient, config_client.IConfigClient) {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId("public"),
		constant.WithUsername("nacos"),
		constant.WithPassword("nacos"),
		constant.WithLogLevel("debug"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
	)
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			ip,
			port,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	// Another way of create the naming client for service discovery (recommend)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	// Another way of create the config client for dynamic configuration (recommend)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}
	return namingClient, configClient
}
