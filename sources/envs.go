package sources

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type envs struct {
	env                                  string
	app_cloud_nacos_discovery_serverAddr string
	app_cloud_nacos_discovery_namespace  string
	data                                 map[string]string
}

func LoadEnvs(envDefault string) *envs {
	n := new(envs)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	if len(envDefault) > 0 {
		env = envDefault
	}
	var myEnv map[string]string
	myEnv, err := godotenv.Read(".env." + env)
	if err != nil {
		log.Fatal(err)
	}

	n.data = make(map[string]string)
	n.data = myEnv

	log.Infof("maps=%#v", myEnv)
	if len(n.env) < 1 {
		n.env = env
	}
	log.Infof("envDefault=%#v", envDefault)

	return n
}

// 获取环境变量名称
func (c envs) GetEnv() string {
	return c.env
}

// 获取nacos 服务器地址
func (c envs) GetDiscoveryServerAddr() string {
	return c.app_cloud_nacos_discovery_serverAddr
}

// 获取 nacos 命名空间
func (c envs) GetDiscoveryNamespace() string {
	return c.app_cloud_nacos_discovery_namespace
}

func (c envs) GetData() map[string]string {
	return c.data
}
