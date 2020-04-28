#!/bin/bash -x

set -e

docker build -t gcsgo -f Dockerfile .
CONTAINER=$(docker run -d gcsgo false)
docker cp $CONTAINER:/go/src/gcsgo/bin/gcsgo gcsgo
docker rm $CONTAINER
