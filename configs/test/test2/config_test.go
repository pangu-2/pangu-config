package test1

import (
	"testing"

	"github.com/pangu-2/pangu-config/configs"
)

func TestConfig(t *testing.T) {
	configs.LoadConfig()
	//服务注册 服务配置获取,并重新初始化配置
	// registryClient.RegistryAndConfigReload()
}
