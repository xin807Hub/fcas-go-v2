#!/bin/bash

BUILD_PATH=build
VER=$(cat ./BUILD_VERSION)

function build()
{
    rm -rf $BUILD_PATH
    mkdir -p $BUILD_PATH
    mkdir $BUILD_PATH/server
    echo "begin to build"
    cd server
    make generate
    make clean
    make release=yes
    make clean
    make
    cp ./config.yaml ./build/
    cd -
    cp -rf server/build/* $BUILD_PATH/server/
    cp -f script/service/fcas_server.service $BUILD_PATH/server/
    echo "build end"
}

build
