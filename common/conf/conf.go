package conf

type DataConfig struct {
	DbConfig    *DatabaseConfig
	RedisConfig *RedisConfig
}

type DatabaseConfig struct {
	Driver string
	Source string
}

type RedisConfig struct {
	Addr         string
	DB           int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}
