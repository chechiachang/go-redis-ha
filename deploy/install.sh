#!/usr/bin/env bash

# ./install.sh values-production.yaml
# or
# ./install.sh values-staging.yaml

#1. Connect to the Redis master pod that you can use as a client:
#
#   kubectl exec -it $(kubectl get pod -o jsonpath='{range .items[*]}{.metadata.name} {.status.containerStatuses[0].state}{"\n"}{end}' -l redis-role=master | grep running | awk '{print $1}') bash
#
#2. Connect using the Redis CLI (inside container):
#
#   redis-cli -a <REDIS-PASS-FROM-SECRET>

HELM_NAME=redis-1

# Stable: chart version: redis-ha-3.6.1	app version: 5.0.5
helm upgrade --install ${HELM_NAME} stable/redis-ha --version 3.6.1 -f values-staging.yaml
