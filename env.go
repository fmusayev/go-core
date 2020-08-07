package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

const defaultFileName = "env.yaml"

func LoadEnv(profiles ...string) {
	LoadEnvFile(defaultFileName, profiles...)
}

func LoadEnvFile(filename string, profiles ...string) {
	yamlFile, err := ioutil.ReadFile(filename)

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("unable to read yaml file: %v", err)
	}

	setEnv(m["default"])
	for _, profile := range profiles {
		if profile == "default" {
			continue
		}
		setEnv(m[profile])
	}
}

func setEnv(in interface{}) {
	m, ok := in.(map[interface{}]interface{})
	if !ok {
		log.Fatalf("unable to parse file content")
	}

	for key, value := range m {
		k := fmt.Sprintf("%v", key)
		v := fmt.Sprintf("%v", value)
		os.Setenv(k, v)
	}
}
