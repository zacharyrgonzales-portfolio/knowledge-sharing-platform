#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d '{"title": "My Article Title 2", "content": "This is the content of my article 2.", "category": "Technology", "tags": ["Golang", "Web Development"], "author_id": 1}' http://localhost:8080/articles