package config

type Config struct {
	DB DB `yaml:"db"`
}
type DB struct {
	URI      string   `yaml:"url"`
	PoolSize int      `yaml:"poolSize"`
	SSL      bool     `yaml:"ssl"`
	CABundle []string `yaml:"sslCA"`
	Cert     string   `yaml:"sslCert"`
}
