package config

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger LoggerConf
	File   struct {
		Path string
	}
	HTTPServer struct {
		Host string
		Port string
	}
	GRPCServer struct {
		Host string
		Port string
	}
	Storage struct {
		DSN string
	}
	Queue struct {
		DSN string
	}
}

type LoggerConf struct {
	Level   string
	LogFile string
}

func (c *Config) NewConfig() *Config {
	return c
}
