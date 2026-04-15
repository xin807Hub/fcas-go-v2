#!/bin/bash

BUILD_PATH=build
MAKE_PATH=src
VER=$(cat ./BUILD_VERSION)

function build()
{
    rm -rf $BUILD_PATH
    mkdir -p $BUILD_PATH/server
    echo "begin to build"

    cd server
    make generate
    make clean
    make release=yes
    make clean
    make
    cp ./config.yaml ./$MAKE_PATH/
    cd -

    cp -rf server/$MAKE_PATH/* $BUILD_PATH/server/
    cp -f  server/fcas_server.service $BUILD_PATH/server/
    echo "build end"
}
build
