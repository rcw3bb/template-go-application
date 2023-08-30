package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	directory string
	filename  string
	extName   string
}

const (
	appConfFilename string = "application"
	appConfExtName  string = "properties"
	appConfDir      string = "./conf"
)

var (
	// Environment The environment to use.
	Environment string
	conf        = &config{
		directory: appConfDir,
		filename:  appConfFilename,
		extName:   appConfExtName,
	}
)

func (conf *config) getConfigFile() string {
	return conf.directory + "/" + conf.filename + "." + conf.extName
}

func (conf *config) getConfigFileByEnv() string {
	return conf.directory + "/" + conf.filename + "-" + Environment + "." + conf.extName
}

func (conf *config) isFilenameByEnvExists() bool {
	path := conf.getConfigFileByEnv()
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// InitConfig initialize the configuration.
func InitConfig() {
	viper.SetConfigFile(conf.getConfigFile())
	if conf.isFilenameByEnvExists() {
		viper.SetConfigFile(conf.getConfigFileByEnv())
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}
