package config

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const env = "COMMANDOR_CONFIG"

var defaultsExtensions = [...]string{"yaml", "yml"}
var defaultsNames = [...]string{"commands", "cmds"}

func Find(config string) (string, error) {
	if config == "" {
		config = os.Getenv(env)
	}

	if config != "" {
		if _, err := os.Stat(config); errors.Is(err, os.ErrNotExist) {
			return "", err
		}

		return config, nil
	}

	for _, name := range defaultsNames {
		for _, extension := range defaultsExtensions {
			config = name + "." + extension

			_, err := os.Stat(config)
			if err == nil {
				return config, nil
			}

			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		}
	}

	return "", errors.New("file not found")
}

func Load(fileName string) (Commands, error) {
	cmds := Commands{}

	file, err := os.ReadFile(fileName)
	if err != nil {
		return cmds, errors.Wrap(err, "load file error")
	}

	err = yaml.Unmarshal(file, &cmds)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cmds, nil
}
