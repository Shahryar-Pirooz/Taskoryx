package config

type Config struct {
	Server   ServerConfig `mapstructure:"http"`
	Database Database     `mapstructure:"database"`
	Redis    Redis        `mapstructure:"redis"`
	Jwt      JWT          `mapstructure:"jwt"`
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
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"poolsize"`
	DialTimeout  int    `mapstructure:"dialtimeout"`
	ReadTimeout  int    `mapstructure:"readtimeout"`
	WriteTimeout int    `mapstructure:"writetimeout"`
}

type JWT struct {
	AccessKey  string `mapstructure:"access_key"`
	RefreshKey string `mapstructure:"refresh_key"`
}
