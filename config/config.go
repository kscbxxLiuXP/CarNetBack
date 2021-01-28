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
	APP.DB_Password = "090312"
	APP.DB_Username = "root"
	APP.DB_IP = "127.0.0.1"
	APP.DB_Port = "3306"
	APP.DB_Name = "carnet"

}
