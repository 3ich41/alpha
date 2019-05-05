package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type config struct {
	MqUsername string
	MqPassword string
	MqHostname string
	MqPort     string
	LogLevel   string
}

// mustGetEnv returns an env variable value if present and fails othwewise
func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

// Config keeps an exposed configuration structure
var Config config

// InitConfig populates config variable and supposed to be called when application started
func InitConfig() {
	m := mustGetEnv
	Config = config{
		MqUsername: m("MQ_USERNAME"),
		MqPassword: m("MQ_PASSWORD"),
		MqHostname: m("MQ_HOSTNAME"),
		MqPort:     m("MQ_PORT"),
		LogLevel:   m("LOGLEVEL"),
	}
}
