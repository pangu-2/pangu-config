package sources

import (
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var flag flagsPg

type flagsPg struct {
	env                                  string `short:"env" long:"env"`
	app_cloud_nacos_discovery_serverAddr string `short:"serverAddr"`
	app_cloud_nacos_discovery_namespace  string `short:"namespace"`
	data                                 map[string]interface{}
}

func LoadFlags(env string) *flagsPg {
	n := new(flagsPg)
	_, err := flags.Parse(&flag)
	if err != nil {
		log.Fatal("Parse error:", err)
	}
	return n
}

// 获取环境变量名称
func (c flagsPg) GetEnv() string {
	return c.env
}

// 获取nacos 服务器地址
func (c flagsPg) GetDiscoveryServerAddr() string {
	return c.app_cloud_nacos_discovery_serverAddr
}

// 获取 nacos 命名空间
func (c flagsPg) GetDiscoveryNamespace() string {
	return c.app_cloud_nacos_discovery_namespace
}

func (c flagsPg) GetData() map[string]interface{} {
	return c.data
}
