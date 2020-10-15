.PHONY: aLl clean



all:   config admin-server dataplat-server cloud-server sysadmin-server

SHELL= /bin/bash

webserver_dir= ../web/server
website_dir = ../website
resource_bin_dir= resource/bin/application
resource_dist_dir=resource/

gotool:
	@echo "格式化代码"
	go fmt ./
	go vet ./

depend:
	go mod tidy
config:
	cp  -rv ../configs/* resource/config/
#p2pnetwork:
#	go build -mod=mod -race -o bin/application/p2pnetwork ../internal/p2p/application/p2p_network.go

admin-server:
	go build -gcflags="all=-m" -ldflags="-X main.version=1.2.3"    -o ${resource_bin_dir}/admin-server  ${webserver_dir}/admin_server/admin_server.go

dataplat-server:
	go build  -gcflags="all=-m"  -ldflags="-X main.version=1.2.3" -o ${resource_bin_dir}/dataplat-server ${webserver_dir}/a/dataplat_server/dataplat_server.go

cloud-server:
	go build  -gcflags="all=-m" -ldflags="-X main.version=1.2.3"  -o ${resource_bin_dir}/cloud-server ${webserver_dir}/a/cloud_server/cloud_server.go


sysadmin-server:
	go build -gcflags="all=-m" -ldflags="-X main.version=1.2.3" -o ${resource_bin_dir}/sysadmin-server ${webserver_dir}/a/sysadmin_server/sysadmin_server.go
	#echo $@
#       gcc -g src/nfstranfs/nfsclient.c src/nfstranfs/fsmk.c src/nfstranfs/fsperm.c src/nfstranfs/fsmount.c  src/nfstranfs/fsutils.c    -o bin/nfstranfs -I include/   -std=c99 -D_XOPEN_SOURCE=500 -D_GNU_SOURCE -Wall
#starfsclient:
         #go build -mod=mod -race -o bin/starfsclient client/starfsclien
website:
	npm --prefix ${website_dir} run build:prod
	cp -rv ${website_dir}/dist  ${resource_dist_dir}
	rm -rf ${website_dir}/dist

images:
	docker build -t

clean:
	rm  -rf resource/bin/*
	rm  -rf resource/dist/
	rm  -rf resource/config/configs

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
