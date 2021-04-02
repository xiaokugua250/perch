#!/bin/bash

#任意命令执行出错则马上退出
set -e
#不允许使用未定义的变量
set -u
#显示调试信息
set -v
## 通過uname -s 獲取操作系統類型
#sysOS=$(uname -a)




GIT_VERSION=$(git rev-parse --short HEAD)
VERSION=V1.0-${GIT_VERSION}
DOCKER_HUB=github.com/markov

RESOURCES_DIR=resource
TARGET_BIN_DIR=resource/bin/application


# shellcheck disable=SC2231
for file in ${TARGET_BIN_DIR}/*;
#for fileName in `ls ${TARGET_BIN_DIR}`;
  do
    service=${file##*/} #//只取文件名
    echo "begin to build $service image ..." &&
    sed  "s/\${SERVICE_NAME}/${service}/g" ${RESOURCES_DIR}/Dockerfile  > ${RESOURCES_DIR}/Dockerfile_tmp &&
    docker build -t github.com/perch/"${service,,}":"${VERSION}" -f ${RESOURCES_DIR}/Dockerfile_tmp ./${RESOURCES_DIR} && # {service,,}將大寫改造成小寫
    rm -rf ${RESOURCES_DIR}/Dockerfile_tmp
  done

version=`git log --date=iso --pretty=format:"%cd @%h" -1`
if [ $? -ne 0 ]; then
    version="not a git repo"
fi

compile=`date +"%F %T %z"`" by "`go version`
author="liangdu"
email="liangdu1992@gmail.com"
verion_location=internal/version
cd $verion_location
cat << EOF | gofmt > version.go
package version

const (
    Version = "$version"
    Compile = "$compile"
    Author= "$author"
    Email="$email"
)
EOF

echo "begin to make project..."
#make
echo "begin to build project with docker-compose..."
#docker-compose build
echo "begin to deploy project with docker-compose..."
#docker-compose up -d

