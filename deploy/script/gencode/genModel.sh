#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

# 加载 .env 文件中的环境变量
if [ -f .env ]; then
  source .env
fi
export DATABASE_HOST
export DATABASE_PORT
export DATABASE_USER
export DATABASE_PASSWORD
#生成的表名
tables=$1
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=$DATABASE_HOST
port=$DATABASE_PORT
dbname=easyGo
username=$DATABASE_USER
passwd=$DATABASE_PASSWORD

echo "开始创建库：$dbname 的表：$1"
mkdir "${tables}Model"
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero
goctl model pg datasource --url "postgres://${username}:${passwd}@${host}:${port}/${dbname}?sslmode=disable" -t "${tables}" -dir="${tables}Model" . -c --style=goZero
