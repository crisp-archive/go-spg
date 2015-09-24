package main

import (
    "encoding/json"
)

type (
    Config struct {
        cfgFile configFile
        cfgData configData
    }

    configFile struct {
        configPath string
        outputPath string
    }

    configData struct {
        title string
        keywords []string
        description string
        contact contactData
        switches switchData
        plugins []pluginData
    }

    contactData struct {
        githubLink bool
        instagramLink bool
        facebookLink bool
        twitterLink bool
        linkedin bool
        weiboLink bool
        wechat bool
        douban bool
    }

    switchData struct {
        enableAbout bool
        enableRss bool
        enableArchive bool
    }

    pluginData struct {
    }
)

func (cfg *Config) checkConfigFile() {
    return true
}

func (cfg *Config) getPlugins() {
}

func (cfg *Config) getConfig() {
}
