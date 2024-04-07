package config

type RsyncTask struct {
	FriendTi string `mapstructure:"friend_ti" json:"friend_ti" yaml:"friend_ti"`
	OpenTi   string `mapstructure:"open_ti" json:"open_ti" yaml:"open_ti"`
}
