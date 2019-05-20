package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type config struct {
	MqUsername   string
	MqPassword   string
	MqHostname   string
	MqPort       string
	GrpcHostname string
	GrpcPort     int
	LogLevel     string
}

// mustGetEnv returns an env variable value if present and fails othwewise
func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Errorf("%s environment variable not set.", key)
		os.Exit(1)
	}
	return val
}

func getIntEnv(key string) int {
	val := getStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.WithField("err", err.Error()).Errorf("Error while reading %v environment variable", key)
		os.Exit(1)
	}
	return ret
}

func getBoolEnv(key string) bool {
	val := getStrEnv(key)
	ret, err := strconv.ParseBool(val)
	if err != nil {
		log.WithField("err", err.Error()).Errorf("Error while reading %v environment variable", key)
		os.Exit(1)
	}
	return ret
}

// Config keeps an exposed configuration structure
var Config config

// InitConfig populates config variable and supposed to be called when application started
func InitConfig() {
	s := getStrEnv
	i := getIntEnv
	Config = config{
		MqUsername:   s("MQ_USERNAME"),
		MqPassword:   s("MQ_PASSWORD"),
		MqHostname:   s("MQ_HOSTNAME"),
		MqPort:       s("MQ_PORT"),
		GrpcHostname: s("GRPC_HOSTNAME"),
		GrpcPort:     i("GRPC_PORT"),
		LogLevel:     s("LOGLEVEL"),
	}
}
