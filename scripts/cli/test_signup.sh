#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{"Username": "john_doe", "Email": "john.doe@example.com", "Password": "mypassword"}' http://localhost:8080/signup
