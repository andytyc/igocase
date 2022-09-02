#!/bin/bash

# ps: 在mac电脑上的使用脚本

if [ "$#" != 1 ];then
    echo "执行错误, 必须有1个参数:\n  编译平台"
    exit 0
fi

if [ "$1" != "mac" -a "$1" != "x86" -a "$1" != "arm" -a "$1" != "win" ]
then
    echo "编译平台,选项不合法,选项:{mac,x86{linux-x86-amd},arm{linux-arm},win{windows}}"
    exit 0
fi

appname="$1-igocase"

if [ "$1" = "mac" ]
then
    export GOOS=darwin
    export GOARCH=amd64
elif [ "$1" = "x86" ]
then
    export GOOS=linux
    export GOARCH=amd64
elif [ "$1" = "arm" ]
then
    export GOOS=linux
    export GOARCH=arm64
elif [ "$1" = "win" ]
then
    export GOOS=windows
    export GOARCH=amd64
    export CGO_ENABLED=0
    appname+=".exe"
fi

go mod tidy

function build_do() {
    md5sum $appname
    go build -o $appname -ldflags "-s -w" .
    md5sum $appname
}

build_do

echo "#### build done $1 {$GOOS, $GOARCH} ####"
