#!/bin/bash
#********************************************************************
#Author: liangdu1992@gmail.com
#website： www.z-gour.com
#Date： 2021-03-12
#FileName： nsq_server_setup.sh
#Description： Annotated script
# ref https://nsq.io/deployment/docker.html
#********************************************************************
set -e
set -u
set -o pipefail

docker-compose -f nsq_compose.yaml up -d

