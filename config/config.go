package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	defaultPort = 8080
)

type Config struct {
	JenkinsServer       string `json:"jenkins_server,omitempty"`
	JenkinsUser         string `json:"jenkins_user,omitempty"`
	JenkinsPassword     string `json:"jenkins_password,omitempty"`
	JenkinsCredentialId string `json:"jenkins_credential,omitempty"`
	Port                int    `json:"port,omitempty"`
}

func Read(path string) (*Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Fail to read the config file %s", path)
	}

	cfg := &Config{}
	err = json.Unmarshal(contents, cfg)
	if err != nil {
		return nil, fmt.Errorf("Fail to unmarshal a JSON object from the config file %s", path)
	}

	// Set the default config for configures not specified
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}

	return cfg, nil
}
