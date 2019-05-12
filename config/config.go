package config

//AppConfig that contains various settings of application
type AppConfig struct {
	DB *DBConfig
}

//DBConfig contains properties for DBConnection
type DBConfig struct {
	Dialect  string
	Host     string
	Port     uint32
	Username string
	Password string
	Name     string
	Charset  string
}

//GetConfig function gets configuration settings for application
func GetConfig() *AppConfig {
	return &AppConfig{
		&DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "root1",
			Name:     "dbtest",
			Charset:  "utf8",
		},
	}
}
