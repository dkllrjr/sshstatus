#!/bin/bash

docker run -dp 127.0.0.1:9429:9429 --name sshstatus-test -v "$(pwd)/config.yml:/app/config.yml" dkllrjr/sshstatus:latest
