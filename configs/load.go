package configs

import (
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() {
	LoadConfigPath(Options{
		Env:           "",
		RootDirectory: "",
	})
}

// 读取配置文件
func LoadConfigPath(opt Options) {

	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	if len(opt.RootDirectory) > 0 {
		viper.AddConfigPath(opt.RootDirectory)
	} else {
		viper.AddConfigPath("./config")
	}

	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
	}
	//
	LoadAppConfig()
	//
}

// 文件是否存在
func fileExist(path string) bool {
	err := syscall.Access(path, syscall.F_OK)
	return !os.IsNotExist(err)
}

// 判断哪种配置文件存在
func existConfigFile(opt Options, fileName, suffix string) string {
	str, _ := os.Getwd()
	log.Infof("Getwd=%#v", str)
	name := fileName + suffix
	if len(opt.Env) > 0 {
		name = fileName + "-" + opt.Env + suffix
	}
	file := "./config/" + name
	if len(opt.RootDirectory) > 0 {
		file = opt.RootDirectory + "/config/" + name
	}
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	if len(opt.RootDirectory) > 0 {
		file = opt.RootDirectory + "/config/" + fileName + suffix
	}
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	file = "./config/" + name
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	file = "./config/" + fileName + suffix
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	return ""
}

// 判断哪种配置文件存在
func existConfigFileDefault(opt Options, fileName, suffix string) string {
	str, _ := os.Getwd()
	log.Infof("Getwd=%#v", str)
	name := fileName + suffix
	file := "./config/" + name
	if len(opt.RootDirectory) > 0 {
		file = opt.RootDirectory + "/config/" + name
	}
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	file = "./config/" + fileName + suffix
	if fileExist(file) {
		log.Infof("config.file[true]=%#v", file)
		return file
	}
	log.Infof("config.file[false]=%#v", file)
	return ""
}

// 判断哪种配置文件存在
func existConfigFileApplication(opt Options) string {
	return existConfigFile(opt, "application", ".toml")
}

// 判断哪种配置文件存在
func existConfigFileDefaultApplication(opt Options) string {
	return existConfigFileDefault(opt, "application", ".toml")
}

// 存储到原始
func SetConfigByRegistry(maps map[string]string) {
	if confRegistrySource == nil {
		confRegistrySource = make(map[string]string)
	}
	// for k, s := range maps {
	// 	//转换为JSON
	// 	j2, err := yaml.YAMLToJSON([]byte(s))
	// 	if err == nil {
	// 		confRegistrySource[k] = string(j2)
	// 	}
	// }
}

// 重新读取配置信息
func LoadAppConfig() {
	var config0 ConfigPg
	// app = App{}
	viper.Unmarshal(&config0)
	configPg = config0
	//设置默认配置
	config0.SetDefault()
	log.Infof("APP=%#v", config0.App)
	log.Infof("config0=%#v", config0)
	//
	//设置默认配置
	config0.App.SetDefault()
}

// 重新初始化
func SetConfigByRegistryAndReload(maps map[string]string) {
	//存储到原始
	SetConfigByRegistry(maps)
	//
	// ReLoadConfigByRegistry()
	//
	LoadAppConfig()
}
