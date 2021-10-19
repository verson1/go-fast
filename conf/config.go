package conf

type DBConfig struct {
	DBType   string //数据库类型 支持mysql,postgre,sqlite
	IP       string //ip
	Port     string //端口
	User     string //用户名
	Password string //密码
	DBName   string //数据库名
	//连接池相关配置
	MaxOpenConn     int  // 用于设置最大打开的连接数，默认值为0表示不限制。
	MaxIdleConn     int  // 用于设置闲置的连接数。
	ConnMaxLifetime int  // 最大生存时间(s)
	PrintSql        bool // 是否打印SQL语句
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
	ListenPort int
	Db         *DBConfig
	Redis      *RedisConfig
	Log        *LogConfig
}


var conf *Config
var once sync.once