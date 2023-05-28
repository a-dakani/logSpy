package configs

import (
	"github.com/a-dakani/LogSpy/logger"
	"os"
	"strconv"
	"strings"
)

type Service struct {
	Name           string `yaml:"name"`
	Host           string `yaml:"host"`
	User           string `yaml:"user"`
	Port           int    `yaml:"port"`
	PrivateKeyPath string `yaml:"private_key_path"`
	Krb5ConfPath   string `yaml:"krb5_conf_path"`
	Files          []File `yaml:"files"`
}
type File struct {
	Alias string `yaml:"alias"`
	Path  string `yaml:"path"`
}
type Services struct {
	Services []Service `yaml:"services"`
}

type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

func (s *Service) IsFullyConfigured() bool {
	propsDefined := s.Name != "" && s.Host != "" && s.User != "" && s.Port != 0 && (s.PrivateKeyPath != "" || s.Krb5ConfPath != "") && len(s.Files) > 0
	if !propsDefined {
		logger.Warning("Service is not fully configured")
		return false
	} else {
		if (s.PrivateKeyPath != "" && !fileExist(s.PrivateKeyPath)) ||
			(s.Krb5ConfPath != "" && !fileExist(s.Krb5ConfPath)) {
			logger.Warning("Private key or krb5.conf file does not exist")
			return false
		}

	}
	for _, file := range s.Files {
		if file.Path == "" || file.Alias == "" {
			logger.Warning("File path or alias is not defined")
			return false

		}
	}
	return true
}

func (s *Services) IsFullyConfigured() bool {
	for _, service := range s.Services {
		if !service.IsFullyConfigured() {
			return false
		}
	}
	return true
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func ParseFiles(files string) []File {
	var parsedFiles []File
	if files != "" {
		for index, file := range strings.Split(files, ",") {
			parsedFiles = append(parsedFiles, File{Alias: strconv.Itoa(index), Path: file})
		}
	}
	return parsedFiles
}
