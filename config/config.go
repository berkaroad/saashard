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

	nodes map[string]*NodeConfig
}

// GetNodes from Nodes
func (config *Config) GetNodes() map[string]*NodeConfig {
	if len(config.nodes) == 0 && len(config.Nodes) > 0 {
		config.nodes = make(map[string]*NodeConfig)
		for _, nodeConfig := range config.Nodes {
			node := nodeConfig
			if _, ok := config.nodes[node.Name]; !ok {
				config.nodes[node.Name] = &node
			}
		}
	}
	return config.nodes
}

// HostConfig is a config of data host.
type HostConfig struct {
	Name             string   `yaml:"name"`
	MaxConnNum       int      `yaml:"max_conn_num"`
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

	tables map[string]*TableConfig
}

// GetTables from Tables
func (schema *SchemaConfig) GetTables() map[string]*TableConfig {
	if len(schema.tables) == 0 && len(schema.Tables) > 0 {
		schema.tables = make(map[string]*TableConfig)
		for _, tableConfig := range schema.Tables {
			table := tableConfig
			if _, ok := schema.tables[table.Name]; !ok {
				schema.tables[table.Name] = &table
			}
		}
	}
	return schema.tables
}

// TableConfig is a config of table
type TableConfig struct {
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
