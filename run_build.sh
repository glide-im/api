#!/bin/bash

if [ ! -n "$1" ]; then
  echo "不支持的操作，请执行：$0 mac|linux|fabu"
  exit 1
fi

name='minidoc'

if [ $1 == "mac" ]; then
  rm -f ${name}_mac
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${name}_mac main.go
  echo "编译mac版完成，文件：${name}_mac"
fi

if [ $1 == "linux" ]; then
  rm -f ${name}_linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${name}_linux main.go
  echo "编译linux版完成，文件：${name}_linux"
fi

if [ $1 == "fabu" ]; then
  rm -f ${name}_linux
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${name}_linux main.go
  echo "编译linux版完成，文件：${name}_linux"
  git add .
  git commit -am "Update_$(date "+%Y-%m-%d_%H%M%S")"
  git push origin master
  echo "推送git完成，请查看企业微信信息检查部署结果！"
fi
