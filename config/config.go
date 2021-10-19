package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
)

const (
	TOML = ".toml"
	YAML = ".yaml"
	YML  = ".yal"
	JSON = ".json"
)

type DBConfig struct {
	Driver string `yaml:"driver"` //数据库类型 支持mysql,postgre,sqlite
	Dsn    string `yaml:"dsn"`
	//连接池相关配置
	MaxOpenConn     int `yaml:"max_open_conn"`     // 用于设置最大打开的连接数，默认值为0表示不限制。
	MaxIdleConn     int `yaml:"max_idle_conn"`     // 用于设置闲置的连接数。
	ConnMaxLifetime int `yaml:"conn_max_lifetime"` // 最大生存时间(s)
	LogLever        int `yaml:"log_lever"`         // 是否打印SQL语句
}

type RedisConfig struct {
	//for conn
	Network  string
	Addr     string
	Password string
	DB       int

	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	//for pool
	PoolSize int
}

type LogConfig struct {
	LogBase     string
	LogLevel    int
	Debug       bool
	LogToStdout bool
	RollSize    int
}

type Config struct {
	ListenPort string       `yaml:"listen_port"`
	DB         *DBConfig    `yaml:"db"`
	Redis      *RedisConfig `yaml:"redis"`
	Log        *LogConfig   `yaml:"log"`
}

var Conf *Config
var once sync.Once

func InitConfig(configPath string) *Config {
	once.Do(func() {
		Conf = parseConfig(configPath)
	})
	return Conf
}

func parseConfig(configPath string) *Config {
	suffix := filepath.Ext(configPath)
	if suffix == "" {
		panic("config path error")
	}

	var err error
	var conf *Config
	var data []byte

	data, err = ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	switch suffix {
	case TOML:
		if err = toml.Unmarshal(data, &conf); err != nil {
			panic(fmt.Sprintf("decode toml config file %s %s", configPath, err))
		}
	case YAML, YML:
		if err = yaml.Unmarshal(data, &conf); err != nil {
			panic(fmt.Sprintf("decode yaml config file %s %s", configPath, err))
		}
	case JSON:
		if err = json.Unmarshal(data, &conf); err != nil {
			panic(fmt.Sprintf("decode json config file %s %s", configPath, err))
		}
	default:
		panic("This config format is not supported")
	}
	return conf
}
