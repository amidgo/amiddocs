package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	URL   string `yaml:"url"`
	Token string `yaml:"client_token"`
}

type PostgresConfig struct {
	LocalURL  string `yaml:"local_url"`
	DockerURL string `yaml:"docker_url"`
}

type JwtConfig struct {
	Name             string `yaml:"name"`
	Pempath          string `yaml:"pempath"`
	AccessTokenTime  uint64 `yaml:"access_token_time"`
	RefreshTokenTime uint64 `yaml:"refresh_token_time"`
}

type TLS struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type FileStorage struct {
	DepFileStorage string `yaml:"department_file_storage"`
	Root           string `yaml:"root"`
}

type Config struct {
	Server      *ServerConfig   `yaml:"server"`
	Database    *PostgresConfig `yaml:"postgres"`
	Jwt         *JwtConfig      `yaml:"jwt"`
	TLS         *TLS            `yaml:"tls"`
	FileStorage *FileStorage    `yaml:"filestorage"`
}

func (c *Config) DatabaseURL() string {
	mode := os.Getenv("MODE")
	switch mode {
	case "local":
		return c.Database.LocalURL
	case "docker":
		return c.Database.DockerURL
	default:
		return c.Database.LocalURL
	}
}

func LoadConfig() *Config {
	b, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := new(Config)
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
