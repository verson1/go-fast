package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	conf := InitConfig("../config.yaml")
	fmt.Println(conf.DB)
}
