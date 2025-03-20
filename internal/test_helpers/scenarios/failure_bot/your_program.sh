#!/bin/sh
cd $(dirname $0)
exec go run app/main.go "$@"