package config

type Config struct {
	Issuer     string `mapstructure:"ISSUER"`
	Release    bool   `mapstructure:"RELEASE"`
	ServerPort string `mapstructure:"SERVERPORT"`
	DEBUGPORT  string `mapstructure:"DEBUGPORT"`
	APIKey     string `mapstructure:"APIKEY"`
	Database   struct {
		Host     string `mapstructure:"HOST"`
		Port     int    `mapstructure:"PORT"`
		Name     string `mapstructure:"NAME"`
		Username string `mapstructure:"USERNAME"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"DATABASE"`
}
