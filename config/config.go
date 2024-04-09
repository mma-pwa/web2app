package config

type Local struct {
	UploadPath string `mapstructure:"uploadPath" json:"uploadPath" yaml:"uploadPath"`
	ImgUrl     string `mapstructure:"img_url" json:"img_url" yaml:"img_url"`
}

type Server struct {
	JWT                JWT                `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap                Zap                `mapstructure:"zap" json:"zap" yaml:"zap"`
	System             System             `mapstructure:"system" json:"system" yaml:"system"`
	Captcha            Captcha            `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mongodb            Mongodb            `mapstructure:"mongodb" json:"mongodb" yaml:"mongodb"`
	Es                 Es                 `mapstructure:"es" json:"es" yaml:"es"`
	Local              Local              `mapstructure:"local" json:"local" yaml:"local"`
	Mysql              Mysql              `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Cors               CORS               `mapstructure:"cors" json:"cors" yaml:"cors"`
	Runmod             Runmod             `json:"runmod" mapstructure:"runmod" yaml:"runmod"`
	PasswordManagement PasswordManagement `mapstructure:"passwordManagement" json:"passwordManagement" yaml:"passwordManagement"`
	Task               Task               `mapstructure:"task" json:"task" yaml:"task"`
	RsyncInterval      RsyncInterval      `mapstructure:"rsyncInterval" json:"rsyncInterval" yaml:"rsyncInterval"`
	RsyncTask          RsyncTask          `mapstructure:"rsync_task" json:"rsync_task" yaml:"rsync_task"`
	MessagesTask       MessagesTask       `mapstructure:"messages_task" json:"messages_task" yaml:"messages_task"`
}
