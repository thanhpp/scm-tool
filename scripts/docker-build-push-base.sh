#!/bin/bash

# current

PWD=$(pwd);

# build scm-base image
cd ~/go/src/github.com/thanhpp/scm-tool/build && \
docker build -f Dockerfile.scmsrv.base -t ghcr.io/thanhpp/scmsrv-base:latest .. && \
docker push ghcr.io/thanhpp/scmsrv-base:latest;

# build nft-base image

cd ~/go/src/github.com/thanhpp/scm-tool/build && \
docker build -f Dockerfile.nftsrv.base -t ghcr.io/thanhpp/nftsrv-base:latest .. && \
docker push ghcr.io/thanhpp/nftsrv-base:latest;

cd $PWD