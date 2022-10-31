package configs

import (
	"context"
)

//配置参数
type Options struct {
	Env           string //环境变量  prod dev test
	RootDirectory string //开发环境配置文件补偿目录
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}
