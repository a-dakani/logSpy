package configs

import (
	"github.com/a-dakani/logSpy/logger"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func LoadServices(srvs *Services) {

	// Just for Binary Usage get the config file path from executable path
	// Get the directory containing the program executable
	executablePath, err := os.Executable()
	if err != nil {
		logger.Fatal(err.Error())
	}
	programDir := filepath.Dir(executablePath)

	// Parse and validate services file
	yamlFile, err := os.ReadFile(filepath.Join(programDir, "config.services.yaml"))
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := yaml.Unmarshal(yamlFile, &srvs); err != nil {
		logger.Fatal(err.Error())
	}
	if srvs.IsFullyConfigured() {
		logger.Info("config.services.yaml file loaded")
	} else {
		logger.Fatal("config.services.yaml file not fully configured")
	}
}

func LoadConfig(cfg *Config) {
	// Just for Binary Usage get the config file path from executable path
	// Get the directory containing the program executable
	executablePath, err := os.Executable()
	if err != nil {
		logger.Fatal(err.Error())
	}
	programDir := filepath.Dir(executablePath)

	// Parse and validate config file
	yamlFile, err := os.ReadFile(filepath.Join(programDir, "config.yaml"))
	if err != nil {
		logger.Fatal(err.Error())
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("config.yaml file loaded")

}
