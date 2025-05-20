package config

type Config struct {
	Server     ServerConfig `mapstructure:"http"`
	Database   Database     `mapstructure:"database"`
	Production bool         `mapstructure:"production"`
	Jwt        JWT          `mapstructure:"jwt"`
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

type JWT struct {
	AccessKey  string `mapstructure:"access_key"`
	RefreshKey string `mapstructure:"refresh_key"`
}
