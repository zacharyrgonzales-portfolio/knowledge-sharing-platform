#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{"Username": "john_doe", "Password": "mypassword"}' http://localhost:8080/login
