package config

import (
	"io/ioutil"

	"github.com/berkaroad/saashard/utils/yaml"
)

// Config is a global config for saashard.
type Config struct {
	BindIP      string `yaml:"bind_ip"`
	ProxyPort   int    `yaml:"proxy_port"`
	AdminPort   int    `yaml:"admin_port"`
	LogPath     string `yaml:"log_path"`
	LogLevel    string `yaml:"log_level"`
	LogSQL      string `yaml:"log_sql"`
	SlowLogTime int    `yaml:"slow_log_time"`
	AllowIps    string `yaml:"allow_ips"`
	BlsFile     string `yaml:"blacklist_sql_file"`
	Charset     string `yaml:"proxy_charset"`

	Hosts   []HostConfig   `yaml:"hosts"`
	Nodes   []NodeConfig   `yaml:"nodes"`
	Schemas []SchemaConfig `yaml:"schemas"`
}

// HostConfig is a config of data host.
type HostConfig struct {
	Name             string   `yaml:"name"`
	MaxConnNum       int      `yaml:"max_conns_limit"`
	DownAfterNoAlive int      `yaml:"down_after_noalive"`
	User             string   `yaml:"user"`
	Password         string   `yaml:"password"`
	Master           string   `yaml:"master"`
	Slaves           []string `yaml:"slaves"`
}

// NodeConfig is a config of data node.
type NodeConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

// SchemaConfig is a config of schema.
type SchemaConfig struct {
	Name      string        `yaml:"name"`
	User      string        `yaml:"user"`
	Password  string        `yaml:"password"`
	ShardKey  string        `yaml:"shard_key"`
	ShardType string        `yaml:"shard_type"`
	Nodes     []string      `yaml:"nodes"`
	Tables    []TableConfig `yaml:"tables"`
	Views     []ViewConfig  `yaml:"views"`
}

// TableConfig is a config of table
type TableConfig struct {
	Name string `yaml:"name"`
}

// ViewConfig is a config of view
type ViewConfig struct {
	Name string `yaml:"name"`
}

// ParseConfigData is to parse config data.
func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// ParseConfigFile is to parse config file.
func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}

// WriteConfigFile is to write to config file.
func WriteConfigFile(cfg *Config, fileName string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
