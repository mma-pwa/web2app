package config

type Es struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
