package main

import (
    "encoding/json"
)

type (
    Config struct {
        cfgFile configFile
    }

    configFile struct {
        configPath string
        outputPath string
    }
)

func (cfg *Config) checkConfigFile() {
    return true
}

func (cfg *Config) getConfig() {
}
