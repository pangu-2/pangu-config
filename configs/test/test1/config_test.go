package test1

import (
	"fmt"
	"testing"

	"github.com/pangu-2/pangu-config/configs"
	"github.com/spf13/viper"
)

type config struct {
	App configs.App
	Pg  configs.Pg
}

func TestConfig(t *testing.T) {
	configs.LoadConfig()
	fmt.Println(viper.Get("app.application.name"))
	fmt.Println(configs.GetApplication())

	var config0 config
	// app = App{}
	viper.Unmarshal(&config0)
	fmt.Println(config0)
}
