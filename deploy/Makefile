.PHONY: aLl clean



all:   gotool depend config target website

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



target:
	@for server in `ls ${webserver_dir}`; \
	do \
	echo "begin to build server" $${server} && \
	echo "build  -o ${resource_bin_dir}/admin-server  ${webserver_dir}/$${server}/$${server}.go" && \
	go build   -gcflags="all=-m" -ldflags="-X main.version=1.2.3" -o ${resource_bin_dir}/$${server}  ${webserver_dir}/$${server}/$${server}.go;\
	done



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


#####################################
# 参考https://www.linkinstar.wiki/2020/04/18/golang/basic/gobuild-x/
# 获取当前commit号
# Version=1.0.0
#CommitID=$(git rev-parse HEAD)
# 获取当前分支名称
#Branch=$(git rev-parse --abbrev-ref HEAD)
# 获取最近的tag
#Tag=$(git describe --abbrev=0 --tags)
## 打包时间
#DATE=$(date +'%Y-%m-%dT%H:%M:%m+08:00')
#go build -ldflags "-X 'main.Version=$Version' -X 'main.Branch=$Branch' -X 'main.CommitID=$CommitID' -X 'main.DATE=$DATE'"
