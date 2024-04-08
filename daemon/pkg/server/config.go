package server

import (
	"fmt"

	"github.com/ovh/configstore"
)

const (
	KeyListenHost = "HTTP_LISTEN_HOST"
	KeyListenPort = "HTTP_LISTEN_PORT"

	DefaultListenHost = "localhost"
	DefaultListenPort = 8080
)

type config struct {
	ListenHost string
	ListenPort int

	Title   string
	Version string

	Verbose bool
}

func loadConfig() (*config, error) {
	var cfg config

	host, err := configstore.GetItemValue(KeyListenHost)
	if err != nil {
		if _, ok := err.(configstore.ErrItemNotFound); !ok {
			return nil, fmt.Errorf("unable to get the HTTP listen host: %w", err)
		}
		cfg.ListenHost = DefaultListenHost
	} else {
		cfg.ListenHost = host
	}

	port, err := configstore.GetItemValueInt(KeyListenPort)
	if err != nil {
		if _, ok := err.(configstore.ErrItemNotFound); !ok {
			return nil, fmt.Errorf("unable to get the HTTP listen port: %w", err)
		}
		cfg.ListenPort = DefaultListenPort
	} else {
		cfg.ListenPort = int(port)
	}

	return &cfg, nil
}
