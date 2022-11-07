package config

import (
	"flag"
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const env = "COMMANDOR_CONFIG"

var defaultsExtensions = [...]string{"yaml", "yml"}
var defaultsNames = [...]string{"commands", "cmds"}

func ConfigFind() (string, error) {
	config := flag.String("config", "", "path for commands file")
	flag.Parse()

	if config == nil || *config == "" {
		env := os.Getenv(env)
		config = &env
	}

	path := *config

	if path == "" {
		for _, name := range defaultsNames {
			for _, extension := range defaultsExtensions {
				path = name + "." + extension

				_, err := os.Stat(path)
				if err == nil {
					return path, nil
				}

				if !errors.Is(err, os.ErrNotExist) {
					return "", err
				}
			}
		}

		return "", errors.New("file not found")
	} else {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return "", err
		}
	}

	return path, nil
}

func ConfigLoad() (Commands, error) {
	cmds := Commands{}

	fileName, err := ConfigFind()
	if err != nil {
		return cmds, errors.Wrap(err, "find file error")
	}

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
