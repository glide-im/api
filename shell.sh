#!/bin/bash

port=8089
pid=$(netstat -nlp | grep :$port | awk '{print $7}' | awk -F"/" '{ print $1 }')
#pid=$(ps -ef | grep 你的进程或端口 | grep -v grep | awk '{print $2}')
echo $pid
kill -9 $pid