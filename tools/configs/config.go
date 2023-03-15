package configs

import "github.com/spf13/viper"

type OldVersion struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Admin    AdminAccount   `mapstructure:"admin"`
}

type NewVersion struct {
	Clients ClientConfig `mapstructure:"clients"`
	Admin   AdminAccount `mapstructure:"admin"`
}

type Configs struct {
	OldVersion OldVersion `mapstructure:"old-version"`
	NewVersion NewVersion `mapstructure:"new-version"`
}

type PostgresConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	DB       string `mapstructure:"database"`
}

type AdminAccount struct {
	ID       string `mapstructure:"email"`
	Password string `mapstructure:"email"`
}

type ClientConfig struct {
	Url string `mapstructure:"url"`
}

func GetConfigs(file, path string) (*Configs, error) {
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Configs
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
