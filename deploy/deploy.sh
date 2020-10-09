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
TARGET_SERVICE="admin_server cloud_server sysadmin_server dataplat_server"
for SERVICE in  ${TARGET_SERVICE[@]}
    do      
        sed  "s/\${SERVICE_NAME}/${SERVICE}/g" ${RESOURCES_DIR}/Dockerfile_Golang  > ${RESOURCES_DIR}/Dockerfile_Golang_$SERVICE
done


version=`git log --date=iso --pretty=format:"%cd @%h" -1`
if [ $? -ne 0 ]; then
    version="not a git repo"
fi

compile=`date +"%F %T %z"`" by "`go version`

cat << EOF | gofmt > version.go
package version

const (
    Version = "$version"
    Compile = "$compile"
)
EOF
