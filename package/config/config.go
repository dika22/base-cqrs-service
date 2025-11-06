package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	RateLimitThreshold        string `env:"RATE_LIMIT_THRESHOLD"`
	RatelimitMaxRetry         string `env:"RATELIMIT_MAX_RETRY"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)
	if c.RateLimitThreshold == "" {
		c.RateLimitThreshold = "100"
	}
	return c
}
