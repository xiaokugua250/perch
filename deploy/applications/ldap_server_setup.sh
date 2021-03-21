#!/bin/bash
#********************************************************************
#Author: liangdu1992@gmail.com
#website： www.z-gour.com
#Date： 2021-03-12
#FileName： ldap_server_setup.sh
#Description： Annotated script
# ref https://github.com/osixia/docker-openldap
#********************************************************************
set -e
set -u
set -o pipefail

docker run --name openldap-server  -p 389:389 -p 636:636    --restart=always -d  osixia/openldap:1.5.0
