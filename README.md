# Commandor

### Commandor is command management system

This system help for management commands in your project.

**Note**: *This system is not build tool like Gradle or Makefile, but you can use this tool for that.*

Utility use config files for description commands.
Default commands config file's name is `commands` or `cmds`.
Now config commands config file's extension only `yaml` or `yml`. 
Also, you can use environment variable `COMMANDOR_CONFIG` (priority) or flag `--config` for in order to 
indicate path to config file.

### Install

```BASH
git clone https://github.com/kolobok-kelbek/commandor.git
cd commandor
./install.sh
sudo chmod 755 /usr/local/bin/commandor
```

### Usage

```BASH
commandor up
```

```BASH
commandor up --config /your_project
```

### Integration tests

For this using bash script:
```BASH
./run_Integration_tests.sh
```

### Commands config example

##### commands.yaml
```YAML

up:
  tags:
    - docker-compose
    - fullstack
  description: start all containers
  command: cat commands.yaml | grep 'go to php container'
  shortCmd: u
  shortcut: u

down:
  tags:
    - docker-compose
    - fullstack
  description: start all containers
  command: echo "docker-compose up"
  shortCmd: d
  shortcut: d

php-exec:
  tags:
    - docker-compose
    - backend
  title: php execute
  description: go to php container
  command: echo "docker-compose exec php"
  shortCmd: pe
  shortcut: p

npm-build:
  tags:
    - frontend
  title: npm build
  description: build static npm files and compile modules
  command: echo "npm i && npm build"
  shortCmd: nb
  shortcut: n

```