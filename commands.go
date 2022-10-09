package main

type commands map[string]command

type command struct {
	Title       string
	Description string
	Command     string
	Tags        []string
	ShortCmd    string `yaml:"shortCmd"`
	Shortcut    string
}
