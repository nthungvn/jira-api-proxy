package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

// AppConfiguration ...
type AppConfiguration struct {
	BaseURL string `json:"baseUrl"`
	PORT    string `json:"port"`
}

// ServerPort return port contain ":" in front of
func (c *AppConfiguration) ServerPort() string {
	if c.PORT == "" {
		c.PORT = "8470"
	}
	logrus.Infof("Listen on %s port", c.PORT)
	return fmt.Sprintf(":%s", c.PORT)
}

func getConfFile() string {
	dir := os.Getenv("CONFIG_DIR")
	if len(dir) == 0 {
		dir = "configs"
	}
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	fileName := []string{"conf.", env, ".json"}
	filePath := path.Join(dir, strings.Join(fileName, ""))
	return filePath
}
