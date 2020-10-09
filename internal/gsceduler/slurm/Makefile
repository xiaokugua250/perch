
.PHONY: all cgslurm

all: bin/cgoslurm

bin/cgoslurm:
        gcc -g src/cgoslurm.c  -o bin/cgoslurm -I include/  -lslurm  -std=c99 -D_XOPEN_SOURCE=500 -D_GNU_SOURCE -Wall

clean:
        rm -rf bin/*