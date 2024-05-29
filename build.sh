#!/bin/bash
docker build -t words-of-wisdom . && docker run -d --rm -p 12000:8080 --name w-o-w words-of-wisdom