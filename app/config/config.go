package config

import (
    "os"
)

type Config struct {
    Port        string
    Environment string
}

func Load() *Config {
    return &Config{
        Port:        os.Getenv("PORT"),
        Environment: os.Getenv("ENVIRONMENT"),
    }
}
