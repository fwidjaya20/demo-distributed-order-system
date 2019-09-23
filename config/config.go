package config

const (
	SERVICE_NAME = "SERVICE_NAME"
)

const (
	EVENT_STORE_GRPC_ADDR = "EVENT_STORE_GRPC_ADDR"
)

var defaultConfig = map[string]string {
	SERVICE_NAME:			"Order System",
	EVENT_STORE_GRPC_ADDR:	":5000",
}

func GetEnv(key string) string {
	return defaultConfig[key]
}
