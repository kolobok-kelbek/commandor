#!/usr/bin/env bash

go build -o commandor .
chmod +x ./commandor
cp $(pwd)/commandor /usr/local/bin/commandor
