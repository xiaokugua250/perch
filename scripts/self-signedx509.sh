# 生成自签名证书
#!/usr/bin/env bash

openssl genrsa -out server.key 2048

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 365


