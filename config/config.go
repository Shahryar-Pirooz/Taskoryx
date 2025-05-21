package config

type Config struct {
	Server     ServerConfig `mapstructure:"http"`
	Database   Database     `mapstructure:"database"`
	Redis      Redis        `mapstructure:"redis"`
	Jwt        JWT          `mapstructure:"jwt"`
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

type Redis struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	PoolSize     string `mapstructure:"poolsize"`
	DialTimeout  string `mapstructure:"dialtimeout"`
	ReadTimeout  string `mapstructure:"readtimeout"`
	WriteTimeout string `mapstructure:"writetimeout"`
	IdleTimeout  string `mapstructure:"idletimeout"`
}

type JWT struct {
	AccessKey  string `mapstructure:"access_key"`
	RefreshKey string `mapstructure:"refresh_key"`
}
