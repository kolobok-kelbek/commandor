package config

type Commands map[string]Command

type Command struct {
	Title       string
	Description string
	Command     string
	Tags        []string
	ShortCmd    string `yaml:"shortCmd"`
	Shortcut    string
}
