package config

type Config struct {
	Server     ServerConfig `mapstructure:"http"`
	Database   Database     `mapstructure:"database"`
	Production bool         `mapstructure:"production"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
}
