package config

import "github.com/spf13/viper"

type Environment struct {
	Config *Config `mapstructure:"CONFIG"`
}

func InitConfig(path string) (*Environment, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	env := &Environment{}
	err := v.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	return env, nil
}
