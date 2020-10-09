.PHONY: aLl clean 

all:  admin-server dataplat-server cloud-server sysadmin-server 

#chatapp:
#go build -mod=mod -race -o bin/application/chatapp ../internal/p2p/application/chatech

website_dir = ../website


#p2pnetwork:
#	go build -mod=mod -race -o bin/application/p2pnetwork ../internal/p2p/application/p2p_network.go

admin-server:
	go build -o bin/application/admin-server ../web/server/admin_server/admin_server.go

dataplat-server:
	go build -o bin/application/dataplat-server ../web/server/dataplat_server/dataplat_server.go

cloud-server:
	go build -o bin/application/cloud-server ../web/server/cloud_server/cloud_server.go


sysadmin-server:
	go build -o bin/application/sysadmin-server../web/server/sysadmin_server/sysadmin_server.go
	echo $@
#       gcc -g src/nfstranfs/nfsclient.c src/nfstranfs/fsmk.c src/nfstranfs/fsperm.c src/nfstranfs/fsmount.c  src/nfstranfs/fsutils.c    -o bin/nfstranfs -I include/   -std=c99 -D_XOPEN_SOURCE=500 -D_GNU_SOURCE -Wall
#starfsclient:
         #go build -mod=mod -race -o bin/starfsclient client/starfsclient.go

website:
	 $(shell cd $(website_dir);pwd)

	npm run build:prod 

images:
	docker build -t 

clean:
	rm  -rf bin/*

