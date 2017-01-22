#!/bin/bash
PROJECT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
go build && ./aa-sbx --config $PROJECT_ROOT/etc/prod.config.yaml
