package server

import "github.com/spf13/pflag"

type ServerConfig struct {
	APIPort         int
	CooridnatorPort int
	DNS             string
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}

func (o *ServerConfig) AddFlags(flags *pflag.FlagSet) {
	flags.IntVar(&o.APIPort, "api-port", 8090, "api port")
	flags.IntVar(&o.CooridnatorPort, "coordinator-port", 8091, "coordinator port")
	flags.StringVar(&o.DNS, "dns", "root:123456@tcp(127.0.0.1:3306)/engine?charset=utf8mb4&parseTime=True&loc=Local", "dns")
}

type WorkerConfig struct {
	WorkerPort     int
	CoordinatorURL string
}

func NewWorkerConfig() *ServerConfig {
	return &ServerConfig{}
}

func (o *WorkerConfig) AddFlags(flags *pflag.FlagSet) {
}
