package ctx

import "chat/pkg/env"

const (
	httpPort      = "HTTP_PORT"
	logLevel      = "LOG_LEVEL"
	railsGrpcHost = "RAILS_GRPC_HOST"
	railsGrpcPort = "RAILS_GRPC_PORT"
)

type Conf struct {
	httpPort      string
	logLevel      string
	railsGrpcHost string
	railsGrpcPort string
}

func loadConf() *Conf {
	return &Conf{
		httpPort:      env.GetEnvValueWithFallback(httpPort, "8080"),
		logLevel:      env.GetEnvValueWithFallback(logLevel, "info"),
		railsGrpcHost: env.GetEnvValueWithFallback(railsGrpcHost, "localhost"),
		railsGrpcPort: env.GetEnvValueWithFallback(railsGrpcPort, "50051"),
	}
}

func (c *Conf) LogLevel() string {
	return c.logLevel
}

func (c *Conf) HttpPort() string {
	return c.httpPort
}

func (c *Conf) RailsGrpcHost() string {
	return c.railsGrpcHost
}

func (c *Conf) RailsGrpcPort() string {
	return c.railsGrpcPort
}
