package config

import (
	log "local/halodoc/log"
)

type Config struct {
	Users    []User     `json:"users"`
	LogTypes []log.Type `json:"log_types"`
}

type User struct {
	Name string
	Logs string
}
