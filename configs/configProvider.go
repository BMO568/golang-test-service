package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user" envconfig:"DB_USERNAME"`
		Password string `yaml:"pass" envconfig:"DB_PASSWORD"`
		Address  string `yaml:"address" envconfig:"DB_ADDRESS"`
		DBName   string `yaml:"dbname" envconfig:"DB_NAME"`
	} `yaml:"database"`
}

func GetConfig() (Config, error) {
	var config Config
	err := readYmlConfig(&config)

	return config, err
}

func readYmlConfig(cfg *Config) error {
	absPath, _ := filepath.Abs("config.yml")
	f, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func readEnvConfig(cfg *Config) error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	cfg.Database.Username = os.Getenv("DB_USERNAME")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.Address = os.Getenv("DB_ADDRESS")
	cfg.Database.DBName = os.Getenv("DB_NAME")

	cfg.Server.Host = os.Getenv("SERVER_HOST")
	cfg.Server.Port = os.Getenv("SERVER_PORT")

	return nil
}
