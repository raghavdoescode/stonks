#!/bin/bash

go build src/main.go

# Move binary into /usr/bin
sudo mv main /usr/bin/stonks

