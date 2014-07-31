package main

import (
	"encoding/json"
	"os"

	"github.com/citadel/citadel"
)

type Config struct {
	SSLCertificate string            `json:"ssl-cert,omitempty"`
	SSLKey         string            `json:"ssl-key,omitempty"`
	CACertificate  string            `json:"ca-cert,omitempty"`
	RedisAddr      string            `json:"redis-addr,omitempty"`
	RedisPass      string            `json:"redis-pass,omitempty"`
	ListenAddr     string            `json:"listen-addr,omitempty"`
	Dockers        []*citadel.Docker `json:"dockers,omitempty"`
}

func loadConfig() error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(&config)
}
