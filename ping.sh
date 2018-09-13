#!/usr/bin/env bash

while true
do
time curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{create(name: \"hello\", shortDescr: \"desc\"){ id }}"}' \
    127.0.0.1:8080/graphql
done


curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{create(name: \"hello123\", shortDescr: \"desc12222\"){ id }}"}' \
    127.0.0.1:8080/graphqlgraphql

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "query{get(id: 2){ id,viewsNumber,name,shortDescr }}"}' \
    127.0.0.1:8080/graphql

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{update(id:2, name: \"changed\"){ id,viewsNumber,name,shortDescr }}"}' \
    127.0.0.1:8080/graphql

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{incrementViewsNumber(id:2){ id,viewsNumber}}"}' \
    127.0.0.1:8080/graphql

while true
do
time curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{incrementViewsNumber(id:2){ id,viewsNumber}}"}' \
    127.0.0.1:8080/graphql
done

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{incrementViewsNumber(id:2){ id,viewsNumber}}"}' \
    127.0.0.1:8080/graphql

while true
do
curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "query{get(id: 2){ id,viewsNumber,name,shortDescr }}"}' \
    127.0.0.1:8080/graphql
done


curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "query{getAll{id,viewsNumber,name,shortDescr}}"}' \
    127.0.0.1:8080/graphql

curl -XPOST \
    -H "Content-Type: application/json" \
    -d '{"query": "mutation{delete(id: 2)}"}' \
    127.0.0.1:8080/graphql

