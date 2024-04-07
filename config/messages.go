package config

type MessagesTask struct {
	Syslog string `json:"syslog" yaml:"syslog"`
	Ntp    string `json:"ntp" yaml:"ntp"`
	Email  string `json:"email" yaml:"email"`
}
