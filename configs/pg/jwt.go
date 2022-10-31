package pg

// jwt
type JwtConfig struct {
	//
	Secret string `json:"secret"`
	// 过期时间，单位 分钟
	Expire int `json:"expire"`
}

// 项目
type Jwt struct {
	// 后台
	Admin JwtConfig `json:"admin"`
	// 前台
	User JwtConfig `json:"user"`
}
