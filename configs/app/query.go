package app

import "github.com/pangu-2/pangu-config/configs/app/query"

// 查询相关
type Query struct {
	ResultCache query.ResultCache `json:"resultCache"`
}
