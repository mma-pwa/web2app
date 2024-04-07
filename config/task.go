package config

type Task struct {
	FriendTi  string `mapstructure:"friend_ti" json:"friend_ti" yaml:"friend_ti"`
	OpenTi    string `mapstructure:"open_ti" json:"open_ti" yaml:"open_ti"`
	LalonTi   string `mapstructure:"lalon_ti" json:"lalon_ti" yaml:"lalon_ti"`
	SyncEvent string `mapstructure:"sync_event" json:"sync_event" yaml:"sync_event"`
	LocalTi   string `mapstructure:"local_ti" json:"local_ti" yaml:"local_ti"`
}
