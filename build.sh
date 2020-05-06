#!/bin/bash

rm -rf testproject*

SOURCE_FILE_NAME=main
TARGET_FILE_NAME=testproject

build(){
    echo $GOOS $GOARCH
    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
    env GOOS=$GOOS GOARCH=$GOARCH \
    go build -o ${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT} \
    -v ${SOURCE_FILE_NAME}.go
    chmod +x ${tname}
    sh image.sh lastest
}

#mac os
#GOOS=darwin
#GOARCH=amd64
#build

#linux
GOOS=linux
GOARCH=amd64
build

#windows
#GOOS=windows
#GOARCH=amd64
#build

#windows
#GOARCH=386
#build