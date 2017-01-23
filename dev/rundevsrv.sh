#!/bin/bash
PROJECT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
go build && ./dev --appfile $PROJECT_ROOT/ansibleapps.yaml
