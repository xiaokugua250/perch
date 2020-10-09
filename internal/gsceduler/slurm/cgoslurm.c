#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include "slurm_errno.h"
#include "slurm.h"

int main(void){

        long api_version = 0L;
        api_version = slurm_api_version();

        printf("%ld\n",api_version);

        return 0;
}
~