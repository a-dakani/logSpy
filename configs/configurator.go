package configs

import (
	"errors"
	"github.com/a-dakani/logSpy/logger"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func LoadServices(srvs *Services) error {

	// Just for Binary Usage get the config file path from executable path
	// Get the directory containing the program executable
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}
	programDir := filepath.Dir(executablePath)

	// Parse and validate services file
	yamlFile, err := os.ReadFile(filepath.Join(programDir, "config.services.yaml"))
	//yamlFile, err := os.ReadFile("config.services.yaml")
	if err != nil {
		logger.Fatal("config.services.yaml file not found")
		return err
	}
	if err = yaml.Unmarshal(yamlFile, &srvs); err != nil {
		logger.Fatal("can not unmarshal config.services.yaml file")
		return err
	}

	if _, err = srvs.IsFullyConfigured(); err != nil {
		return errors.New("config.services.yaml file not fully configured")
	}

	logger.Info("config.services.yaml file loaded")
	return nil
}

func LoadConfig(cfg *Config) error {
	// Just for Binary Usage get the config file path from executable path
	// Get the directory containing the program executable
	executablePath, err := os.Executable()
	if err != nil {
		return err
	}
	programDir := filepath.Dir(executablePath)

	// Parse and validate config file
	yamlFile, err := os.ReadFile(filepath.Join(programDir, "config.yaml"))
	//yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		logger.Fatal("config.yaml file not found")
		return err
	}
	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		logger.Fatal("can not unmarshal config.yaml file")
		return err
	}
	logger.Info("config.yaml file loaded")
	return nil

}
