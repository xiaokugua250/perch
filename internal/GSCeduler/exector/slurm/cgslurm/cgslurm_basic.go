package cgslurm

/*

#cgo CFLAGS: -I ../include/
#cgo LDFLAGS: -lslurm
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include "slurm_errno.h"
#include "slurm.h"

*/
import (
	"C"
	"errors"
)

//获取slurm api version
func GetSlurmBasicAPIVersion() (version int64) {

	api_version := C.slurm_api_version()

	return int64(api_version)

}

/**
ping slurm
*/
func PingSlurm(dest int) (code int, err error) {
	code = C.slurm_ping(dest)
	//errcode = C.slurm_get_errno()
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return code, err
}

//终端输出slurm Ctl Conf信息
func PrintSlurmCtlConfInfo() (errcode int, err error) {
	var (
		ctl_conf_update_time C.time_t
		ctl_conf_info_msg    *C.struct_slurm_ctl_conf
	)
	C.slurm_load_ctl_conf(ctl_conf_update_time, &ctl_conf_info_msg, 1)
	C.slurm_print_ctl_conf(C.stdout, &ctl_conf_info_msg)
	C.slurm_free_ctl_conf(ctl_conf_info_msg)
	errcode = int(C.slurm_get_errno())
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return errcode, err

}

//获取Slurm Ctl Conf配置信息接口
func GetSlurmCtlConfInfo() (slurmCtlConf model.CGoSlurmCtlConf, errcode int, err error) {
	var (
		ctl_conf_update_time C.time_t
		ctl_conf_info_msg    *C.struct_slurm_ctl_conf
	)
	C.slurm_load_ctl_conf(ctl_conf_update_time, &ctl_conf_info_msg, 1)
	//jobs_array = jobs_info_msg.node_array
	//todo 解析slurm ctl_conf_info_msg 结构体到 CGoSlurmCtlConf结构体中

	/*
		slurm_jobs := (*[1 << 32]C.struct_jobs_info)(unsafe.Pointer(jobs_array))[:jobs_info_msg.record_count:jobs_info_msg.record_count]
		for _, job := range slurm_jobs {
			var cgslurmjob model.CGoSlurmJobs
			cgslurmjob.JobAccount = C.GoString(job.account)
			cgslurmjob.JobAccrueTime = C.GoString(job.accrue_time)
			cgslurmjob.JobAdminComment = C.GoString(job.admin_comment)
			cgslurmjob.JobAllocNode = C.GoString(job.alloc_node)

			slurmJobs = append(slurmJobs, cgslurmjob)
		}
	*/
	C.slurm_free_ctl_conf(ctl_conf_info_msg)
	errcode = int(C.slurm_get_errno())
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))

	return slurmJobs, errcode, err
}
