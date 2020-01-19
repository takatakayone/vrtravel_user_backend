package config

import (
	"os"
)

const(
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var (
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

type DbConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host string `json:"host"`
	Schema string `json:"schema"`
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Username: username,
		Password: password,
		Host:     host,
		Schema:   schema,
	}
}