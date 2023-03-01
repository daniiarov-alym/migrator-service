package config

import (
	"fmt"
	"os"
	logger "github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

var conf Config

func Init() {
	defer func() {
		if r := recover(); r != nil {
			logger.Fatalf("Failed to read configuration: %s", r)
		}
	}()
	conf.Host = readField("PG_HOST")
	conf.Port = readField("PG_PORT")
	conf.User = readField("PG_USER")
	conf.Password = readField("PG_PASSWORD")
	conf.Database = readField("PG_DATABASE")
}

func Conf() Config {
	return conf
}

func readField(field string) string {
	val := os.Getenv(field)
	if val == "" {
		panic(fmt.Errorf("key %s is absent", field))
	}
	return val
}