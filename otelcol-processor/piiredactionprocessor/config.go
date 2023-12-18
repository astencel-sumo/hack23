package piiredactionprocessor

import "go.opentelemetry.io/collector/component"

type Config struct {
	Endpoint string `mapstructure:"endpoint"`
}

func createDefaultConfig() component.Config {
	return &Config{
		Endpoint: "http://localhost:5000/",
	}
}
