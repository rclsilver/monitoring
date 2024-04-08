package mqtt

import (
	"fmt"

	"github.com/ovh/configstore"
)

const (
	KeyHost = "MQTT_HOST"
	KeyPort = "MQTT_PORT"

	DefaultHost = "localhost"
	DefaultPort = 1883
)

type Config struct {
	Host string
	Port int
}

func LoadConfig() (*Config, error) {
	var cfg Config

	host, err := configstore.GetItemValue(KeyHost)
	if err != nil {
		if _, ok := err.(configstore.ErrItemNotFound); !ok {
			return nil, fmt.Errorf("unable to get the MQTT host: %w", err)
		}
		cfg.Host = DefaultHost
	} else {
		cfg.Host = host
	}

	port, err := configstore.GetItemValueInt(KeyPort)
	if err != nil {
		if _, ok := err.(configstore.ErrItemNotFound); !ok {
			return nil, fmt.Errorf("unable to get the MQTT port: %w", err)
		}
		cfg.Port = DefaultPort
	} else {
		cfg.Port = int(port)
	}

	return &cfg, nil
}
