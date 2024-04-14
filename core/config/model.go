package config

type Config struct {
	SERVICENAME string `mapstructure:"SERVICENAME"`
	Issuer      string `mapstructure:"ISSUER"`
	Release     bool   `mapstructure:"RELEASE"`
	ServerPort  string `mapstructure:"SERVERPORT"`
	DEBUGPORT   string `mapstructure:"DEBUGPORT"`
	APIKey      string `mapstructure:"APIKEY"`
	LOGFORMAT   string `mapstructure:"LOGFORMAT"`
	LOGPATH     string `mapstructure:"LOGPATH"`
	Database    struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		Name     string `mapstructure:"NAME"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"DATABASE"`
}
