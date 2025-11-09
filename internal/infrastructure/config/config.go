package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string    `yaml:"env" env-default:"prod"`
	DB  DBconfig  `yaml:"database" env-required:"true"`
	App AppConfig `yaml:"app" env-required:"true"`
}

type DBconfig struct {
	Username string `yaml:"username" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	DBname   string `yaml:"db_name" env-required:"true"`
	SSLmode  string `yaml:"sslmode" env-required:"true"`
	Driver   string `yaml:"driver"`
}

type AppConfig struct {
	Username    string `yaml:"http_username" env-required:"true"`
	Password    string `yaml:"http_password" env-required:"true"`
	Host        string `yaml:"host" env-required:"true"`
	Port        string `yaml:"port" env-required:"true"`
	RWTimeout   string `yaml:"rw_timeout" env-required:"true"`
	IdleTimeout string `yaml:"idle_timeout" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("Config path is empty")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file not found: " + path)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Error loading config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
