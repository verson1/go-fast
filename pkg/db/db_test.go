package db

import (
	"fmt"
	"go-fast/pkg/config"
	"testing"
)

func TestNewDB(t *testing.T) {
	conf := &config.DBConfig{
		Driver:          "postgre",
		Dsn:             "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=123456 sslmode=disable TimeZone=Asia/Shanghai",
		MaxOpenConn:     1,
		MaxIdleConn:     2,
		ConnMaxLifetime: 60,
		LogLever:        4,
	}
	db, err := NewDB(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
}
