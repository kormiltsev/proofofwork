#!/bin/bash
docker build -t words-of-wisdom-client -f Dockerfile.client . && docker run -it --rm --net=host --name w-o-w-c words-of-wisdom-client