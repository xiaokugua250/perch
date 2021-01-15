.PHONY: aLl clean website



all:   gotool depend commit config target website images

SHELL= /bin/bash

images_prefix=z-gour.com/perch
deploy_dir = ${shell pwd}/deploy
webserver_dir =${shell pwd}/web/server
website_dir = ${shell pwd}/website
bin_dir = ${shell pwd}/deploy/bin
dockerfiles_dir= ${shell pwd}/deploy/dockerfiles
gotool:
	@echo "格式化代码"
	go fmt ./
	go vet ./

commit:
	GIt_COMMIT=$(shell cd ${shell pwd} && git rev-parse --short HEAD)


depend:
	go mod tidy
config:
	cp  -rv ../configs/* resource/config/
#p2pnetwork:
#	go build -mod=mod -race -o bin/application/p2pnetwork ../internal/p2p/application/p2p_network.go
# go build   -gcflags="all=-m" -ldflags="-X main.version=1.2.3" -o ${bin_dir}/$${server}  ${webserver_dir}/$${server}/$${server}.go;

target:
	@for server in `ls ${webserver_dir}`; \
	do \
	echo "go build  -o ${bin_dir}/$${server}  ${webserver_dir}/$${server}/$${server}.go" && \
	go build  -o ${bin_dir}/$${server}  ${webserver_dir}/$${server}/$${server}.go; \
	done

website:
	npm --prefix ${website_dir} install 
	npm --prefix ${website_dir} run build:prod
	mv -v ${website_dir}/dist  ${deploy_dir}/resource/
website_image:
	$(eval GIT_COMMIT=$(shell cd ${shell pwd} && git rev-parse --short HEAD))
	docker build -t ${images_prefix}/website:$(GIT_COMMIT) -f ${deploy_dir}/dockerfiles/website/Dockerfile ${deploy_dir}
	

images:
	#@for server in `ls ${bin_dir}`; do echo " docker build -t github.com/perch/$${server} . " && go build  -o ${bin_dir}/$${server}  ${webserver_dir}/$${server}/$${server}.go; docker build -t done
	$(eval GIT_COMMIT=$(shell cd ${shell pwd} && git rev-parse --short HEAD))
	#docker build -t ${images_prefix}/website:$(GIT_COMMIT) -f ${deploy_dir}/dockerfiles/website/Dockerfile   ${deploy_dir}
	@for service in `ls ${bin_dir}`; \
		do  echo $${service}  && \
		sed  "s/service_bin/$${service}/g" ${dockerfiles_dir}/services/Dockerfile  > ${dockerfiles_dir}/services/Dockerfile_$${service} && \
		 docker build -t ${images_prefix}/$${service}:$(GIT_COMMIT) -f ${deploy_dir}/dockerfiles/services/Dockerfile_$${service} ${deploy_dir} && \
		 rm -rf ${deploy_dir}/dockerfiles/services/Dockerfile_$${service} ;\
	done


clean:
	rm  -rf ${deploy_dir}/bin/*
	rm  -rf ${deploy_dir}/resource/bin/*
	rm  -rf ${deploy_dir}/resource/dist/
	rm  -rf ${deploy_dir}/resource/config/configs

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
