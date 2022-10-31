package cloud

type Nacos struct {
	Discovery Discovery `json:"discovery"`
	Config    Config    `json:"config"`
}
