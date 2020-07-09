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

func PrintSlurmBasicFreeNodeInfo() (errcode int, err error) {
	var (
		node_info_msg    *C.node_info_msg_t
		node_update_time C.time_t
		node_show_flags  C.ushort
	)
	C.slurm_load_node(node_update_time, &node_info_msg, node_show_flags)
	C.slurm_print_node_info_msg(C.stdout, node_info_msg, 0)
	C.slurm_free_node_info_msg(node_info_msg)
	errcode = C.slurm_get_errno()
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return errcode, err
}

//获取Slurm Nodes 信息接口
func GetSlurmBasicFreeNodeInfo() (slurmNodes []model.CGoSlurmNode, errcode int, err error) {
	var (
		node_info_msg    *C.node_info_msg_t
		node_update_time C.time_t
		node_show_flags  C.ushort
		node_array       *C.struct_node_info
	)
	C.slurm_load_node(node_update_time, &node_info_msg, node_show_flags)

	node_array = node_info_msg.node_array

	slurm_nodes := (*[1 << 32]C.struct_node_info)(unsafe.Pointer(node_array))[:node_info_msg.record_count:node_info_msg.record_count]
	for _, node := range slurm_nodes {
		var slurm_node model.CGoSlurmNode
		slurm_node.NodeArch = C.GoString(node.arch)
		slurm_node.NodeBoards = int(node.boards)
		slurm_node.NodeBootTime = C.GoString(node.arch)
		slurm_node.NodeClusterName = C.GoString(node.cluster_name)
		slurm_node.NodeCores = int(node.cores)
		slurm_node.NodeCoreSpecCnt = int(node.core_spec_cnt)
		slurm_node.NodeCpuBind = int(node.cpu_bind)
		slurm_node.NodeCpuLoad = int(node.cpu_load)
		slurm_node.NodeFreeMem = int(node.free_mem)
		slurm_node.NodeCpus = int(node.cpus)
		slurm_node.NodeCpuSpecList = C.GoString(node.cpu_spec_list)
		slurm_node.NodeFeatures = C.GoString(node.features)
		slurm_node.NodeFeaturesAct = C.GoString(node.features_act)
		slurm_node.NodeGres = C.GoString(node.gres)
		slurm_node.NodeGresDrain = C.GoString(node.gres_drain)
		slurm_node.NodeGresUsed = C.GoString(node.gres_used)
		slurm_node.NodeMcsLabel = C.GoString(node.mcs_label)
		slurm_node.NodeMemSpecLimit = C.GoString(node.mem_spec_limit)
		slurm_node.NodeName = C.GoString(node.name)
		slurm_node.NodeNextState = C.GoString(node.next_state)
		slurm_node.NodeNodeAddr = C.GoString(node.node_addr)
		slurm_node.NodeNodeHostname = C.GoString(node.node_hostname)
		slurm_node.NodeNodeState = C.GoString(node.node_state)
		slurm_node.NodeOS = C.GoString(node.os)
		slurm_node.NodeOwner = C.GoString(node.owner)
		slurm_node.NodePartitions = C.GoString(node.partitions)
		slurm_node.NodePort = C.GoString(node.port)
		slurm_node.NodeRealMemory = C.GoString(node.real_memory)
		slurm_node.NodeReason = C.GoString(node.reason)
		slurm_node.NodeReasonTime = C.GoString(node.reason_time)
		slurm_node.NodeSockets = C.GoString(node.sockets)
		slurm_node.NodeThreads = int(node.threads)
		slurm_node.NodeTmpDisk = int(node.tmp_disk)
		slurm_node.NodeWeight = int(node.weight)
		slurm_node.NodeTresFmtStr = C.GoString(node.tres_fmt_str)
		slurm_node.NodeVersion = C.GoString(node.version)
		slurmNodes = append(slurmNodes, slurm_node)
	}
	C.slurm_free_node_info_msg(node_info_msg)
	errcode = C.slurm_get_errno()
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return slurmNodes, errcode, err
}
