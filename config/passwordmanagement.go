package config

type PasswordManagement struct {
	AdminDefaultPassWord    string `json:"adminDefaultPassWord" bson:"adminDefaultPassWord" yaml:"adminDefaultPassWord"`
	OperatorDefaultPassWord string `json:"operatorDefaultPassWord" bson:"operatorDefaultPassWord" yaml:"operatorDefaultPassWord"`
	AuditrDefaultPassWord   string `json:"auditrDefaultPassWord" bson:"auditrDefaultPassWord" yaml:"auditrDefaultPassWord"`
	ResetPassWord           string `json:"resetPassWord" bson:"resetPassWord" yaml:"resetPassWord"`
}
