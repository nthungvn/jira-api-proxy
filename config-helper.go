package main

import (
	"os"
	"path"
	"strings"
)

// AppConfiguration ...
type AppConfiguration struct {
	BaseURL string `json:"baseUrl"`
}

func getConfFile() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	fileName := []string{"conf.", env, ".json"}
	filePath := path.Join("configs", strings.Join(fileName, ""))
	return filePath
}
