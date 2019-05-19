package config

const (
	keysPath = "D:/Gopath/src/github.com/serafimone/InformationTagger/config/keys/"
)

//AppConfig that contains various settings of application
type AppConfig struct {
	DB             *DBConfig
	PublicKeyPath  string
	PrivateKeyPath string
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

var appConfig *AppConfig = nil

//GetConfig function gets configuration settings for application
func GetConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{
			DB: &DBConfig{
				Dialect:  "mysql",
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "root1",
				Name:     "dbtest",
				Charset:  "utf8",
			},
			PublicKeyPath:  keysPath + "public_key.pub",
			PrivateKeyPath: keysPath + "private_key",
		}
	}
	return appConfig
}
