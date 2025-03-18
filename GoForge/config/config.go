package config

type Config struct {
	Port      string
	DebugMode bool
}

func LoadConfig() Config {
	return Config{
		Port:      ":8080",
		DebugMode: true, // Enable live reloading in debug mode
	}
}
