package cgslurm

/*
   //http://blog.sina.com.cn/s/blog_48c95a190102w2ln.html

   #cgo CFLAGS: -I ../include/
   #cgo LDFLAGS: -lslurm
   #include <stdio.h>
   #include <stdlib.h>
   #include <signal.h>
   #include <time.h>
   #include "slurm_errno.h"
   #include "slurm.h"

*/
import (
	"C"
)
import (
	"errors"
	_ "reflect"

	"unsafe"
)

//终端输出slurm分区信息
func PrintSlurmBasicPartitionInfo() (errcode int, err error) {
	var (
		partition_update_time C.time_t
		partition_info_msg    *C.struct_partition_info_msg
	)
	C.slurm_load_partitions(partition_update_time, &partition_info_msg, 1)
	C.slurm_print_partition_info_msg(C.stdout, partition_info_msg, 1)
	C.slurm_free_partition_info_msg(partition_info_msg)
	errcode = C.slurm_get_errno()
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return errcode, err
}

//获取Slurm Partition 分区信息接口
func GetSlurmBasicPartitionInfo() (slurmPartitions []model.CGoSlurmPartitions, errcode int, err error) {
	var (
		partition_update_time C.time_t
		partition_info_msg    *C.struct_partition_info_msg
		partition_array       *C.struct_partition_array
	)
	C.slurm_load_partitions(partition_update_time, &partition_info_msg, 1)
	partition_array = partition_info_msg.node_array

	slurm_partitions := (*[1 << 32]C.struct_partition_info)(unsafe.Pointer(partition_array))[:partition_info_msg.record_count:partition_info_msg.record_count]
	for _, partition := range slurm_partitions {
		var cgslurmpartition model.CGoSlurmPartitions
		cgslurmpartition.PartitionAllowAllocNode = C.GoString(partition.allow_alloc_nodes)
		cgslurmpartition.PartitionAllowAccount = C.GoString(partition.allow_accounts)
		cgslurmpartition.PartitionAllowGroups = C.GoString(partition.allow_groups)
		cgslurmpartition.PartitionAllowQos = C.GoString(partition.allow_qos)
		cgslurmpartition.PartitionAlternate = C.GoString(partition.alternate)
		cgslurmpartition.PartitionBillingWeightsStr = C.GoString(partition.billing_weights_str)
		cgslurmpartition.PartitionClusterName = C.GoString(partition.cluster_name)
		cgslurmpartition.PartitionCrType = int(partition.cr_type)
		cgslurmpartition.PartitionCpuBind = int(partition.cpu_bind)
		cgslurmpartition.PartitionDefMemPerCpu = int(partition.def_mem_per_cpu)
		cgslurmpartition.PartitionDefaultTime = int(partition.default_time)
		cgslurmpartition.PartitionDenyAccounts = C.GoString(partition.deny_accounts)
		cgslurmpartition.PartitionDenyQos = C.GoString(partition.deny_qos)
		cgslurmpartition.PartitionFlags = int(partition.flags)
		cgslurmpartition.PartitionGraceTime = int(partition.grace_time)
		cgslurmpartition.PartitionJobDefaultsStr = C.GoString(partition.job_defaults_str)
		cgslurmpartition.PartitionMaxCpusPerNode = int(partition.max_cpus_per_node)
		cgslurmpartition.PartitionMaxMemPerCpu = int(partition.max_mem_per_cpu)
		cgslurmpartition.PartitionMaxNodes = int(partition.max_nodes)
		cgslurmpartition.PartitionMaxShare = int(partition.max_share)
		cgslurmpartition.PartitionMaxTime = int(partition.max_time)
		cgslurmpartition.PartitionMinNodes = int(partition.min_nodes)
		cgslurmpartition.PartitionName = C.GoString(partition.name)
		//	cgslurmpartition.PartitionNodeInx =int(partition.node_inx)
		cgslurmpartition.PartitionNodes = C.GoString(partition.nodes)
		cgslurmpartition.PartitionOverTimeLimit = int(partition.over_time_limit)
		cgslurmpartition.PartitionPreemptMode = int(partition.preempt_mode)
		cgslurmpartition.PartitionPriorityJobFactor = int(partition.priority_job_factor)
		cgslurmpartition.PartitionPriorityTier = int(partition.priority_tier)
		cgslurmpartition.PartitionQosChar = C.GoString(partition.qos_char)
		cgslurmpartition.PartitionStateUp = int(partition.state_up)
		cgslurmpartition.PartitionTotalCpus = int(partition.total_cpus)
		cgslurmpartition.PartitionTotalNodes = int(partition.total_nodes)
		cgslurmpartition.PartitionTresFmtStr = C.GoString(partition.tres_fmt_str)

		slurmPartitions = append(slurmPartitions, cgslurmpartition)
	}
	C.slurm_free_partition_info_msg(partition_info_msg)
	errcode = C.slurm_get_errno()

	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return slurmPartitions, errcode, err
}
