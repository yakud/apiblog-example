#!/usr/bin/env bash

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "{ hello }"}' \
    127.0.0.1:8080/