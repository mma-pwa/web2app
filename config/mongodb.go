package config

type Mongodb struct {
	MgAddrs  string `json:"mgAddrs"`
	MgDbName string `json:"mgDbName"`
}

func (m *Mongodb) Dsn() string {
	return m.MgAddrs
}
