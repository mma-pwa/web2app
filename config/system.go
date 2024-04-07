package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Addr          string `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 地址
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimitCount" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimitTime" yaml:"iplimit-time"`
}
