#!/bin/bash

PORT=6001

# 停止指定端口上的进程
stop_process() {
  # 查找指定端口的进程
  PID=$(netstat -nlp | grep ":$PORT " | awk '{print $7}' | awk -F'[/]' '{print $1}')

  if [[ -z $PID ]]; then
    echo "端口 $PORT 上没有运行的进程"
  else
    echo "停止端口 $PORT 上的进程 $PID"
    kill $PID
  fi
}

# 更改文件权限
change_permission() {
  chmod 777 ./main
  echo "修改文件权限"
}

# 启动进程
start_process() {
  nohup ./main --gf.gcfg.file=config.pro.yaml >/dev/null 2>&1 &
  echo "部署完成"
}

# 停止进程
stop_process

# 更改文件权限
change_permission

# 启动进程
start_process
