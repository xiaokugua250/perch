.PHONY: aLl

all: chatapp p2pnetwork

chatapp:
	go build -mod=mod -race -o bin/application/chatapp ../internal/p2p/application/chatech



p2pnetwork:
	go build -mod=mod -race -o bin/application/p2pnetwork ../internal/p2p/application/p2p_network.go


#       gcc -g src/nfstranfs/nfsclient.c src/nfstranfs/fsmk.c src/nfstranfs/fsperm.c src/nfstranfs/fsmount.c  src/nfstranfs/fsutils.c    -o bin/nfstranfs -I include/   -std=c99 -D_XOPEN_SOURCE=500 -D_GNU_SOURCE -Wall
#starfsclient:
         #go build -mod=mod -race -o bin/starfsclient client/starfsclient.go

clean:
	rm  -rf bin/*

