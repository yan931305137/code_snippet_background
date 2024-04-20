package types

// Config 配置文件结构体
type Config struct {
	MysqlDatabase DatabaseConfig `toml:"MysqlDatabase"`
	Logger        LoggerConfig   `toml:"Logger"`
	Admin         AdminConfig    `toml:"Admin"`
	RedisDatabase RedisConfig    `toml:"RedisDatabase"`
}

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	Master string `toml:"Master"`
	Slave  string `toml:"Slave"`
	Debug  bool   `toml:"Debug"`
	Log    string `toml:"Log"`
}

// LoggerConfig 日志配置结构体
type LoggerConfig struct {
	Path   string `toml:"Path"`
	Level  int    `toml:"Level"`
	Stdout bool   `toml:"Stdout"`
}

// AdminConfig 管理员配置结构体
type AdminConfig struct {
	Version string `toml:"Version"`
	Debug   bool   `toml:"Debug"`
	Image   string `toml:"Image"`
	Uploads string `toml:"Uploads"`
}

// RedisConfig Redis数据库配置结构体
type RedisConfig struct {
	Addr     string `toml:"Addr"`
	Password string `toml:"Password"`
	DB       int    `toml:"DB"`
}
