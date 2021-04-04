package config

var APP APPconfig

type APPconfig struct {
	DB_Username string
	DB_Password string
	DB_IP       string
	DB_Port     string
	DB_Name     string
}

func LoadConfig() {
	APP.DB_Password = "Neu12345"
	APP.DB_Username = "neu"
	APP.DB_IP = "rm-2zey0tuwg85789b4zbo.mysql.rds.aliyuncs.com"
	APP.DB_Port = "3306"
	APP.DB_Name = "carnet"

}
