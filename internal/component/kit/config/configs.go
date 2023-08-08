package config

// Configs 配置文件映射文件
// 该文件是和 configs/app.xxx.yaml 的配置映射文件

type AppConfig struct {
	Name     string `yaml:"NAME"`
	Env      string `yaml:"ENV"`
	Debug    bool   `yaml:"DEBUG"`
	URL      string `yaml:"URL"`
	Timezone string `yaml:"TIMEZONE"`
}

type DatabaseConfig struct {
	Mysql []MysqlConfig `yaml:"Mysql"`
	Redis []RedisConfig `yaml:"Redis"`
}

type MysqlConfig struct {
	Name string `yaml:"Name"`
	DSN  string `yaml:"DSN"`
}

type RedisConfig struct {
	Name     string `yaml:"Name"`
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

type Configs struct {
	App      AppConfig      `yaml:"APP"`
	Database DatabaseConfig `yaml:"Database"`
}
