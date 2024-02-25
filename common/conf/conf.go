package conf

type ZkConfig struct {
}

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

type HttpServerConfig struct {
	Port int
}

type TcpServerConfig struct {
	Port int
}

type NatsConfig struct {
	Addr                   string
	ConnectTimeout         int
	MaxReconnectionRetries int
	RequestTimeout         int
}

type ChatConfigItemConfig struct {
	Channel int
	AppIds  []int
}

type ChannelConfigItem struct {
	ChatType int
	Config   []*ChatConfigItemConfig
}

type ChannelConfig struct {
	Items []*ChannelConfigItem
}
