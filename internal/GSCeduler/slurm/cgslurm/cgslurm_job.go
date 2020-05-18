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
	"starlight/common/model"
	"unsafe"
)

//终端输出slurm分区信息
func PrintSlurmBasicJobsInfo() (errcode int, err error) {
	var (
		jobs_update_time C.time_t
		jobs_info_msg    *C.struct_job_info_msg
	)
	C.slurm_load_jobs(jobs_update_time, &jobs_info_msg, 1)
	C.slurm_print_job_info_msg(C.stdout, jobs_info_msg, 1)
	C.slurm_free_job_info_msg(jobs_info_msg)
	errcode = int(C.slurm_get_errno())
	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return errcode, err

}

//获取Slurm Partition 分区信息接口
func GetSlurmBasicJobsInfo() (slurmJobs []model.CGoSlurmJobs, errcode int, err error) {
	var (
		jobs_update_time C.time_t
		jobs_info_msg    *C.struct_job_info_msg
		jobs_array       *C.struct_job_array
	)
	C.slurm_load_partitions(jobs_update_time, &jobs_info_msg, 1)
	jobs_array = jobs_info_msg.node_array

	slurm_jobs := (*[1 << 32]C.struct_jobs_info)(unsafe.Pointer(jobs_array))[:jobs_info_msg.record_count:jobs_info_msg.record_count]
	for _, job := range slurm_jobs {
		var cgslurmjob model.CGoSlurmJobs
		cgslurmjob.JobAccount = C.GoString(job.account)
		cgslurmjob.JobAccrueTime = C.GoString(job.accrue_time)
		cgslurmjob.JobAdminComment = C.GoString(job.admin_comment)
		cgslurmjob.JobAllocNode = C.GoString(job.alloc_node)
		cgslurmjob.JobAllocSid = C.GoString(job.alloc_sid)
		cgslurmjob.JobArrayBitmap = C.GoString(job.array_bitmap)
		cgslurmjob.JobArrayJobID = C.GoString(job.array_job_id)
		cgslurmjob.JobArrayTaskID = C.GoString(job.array_task_id)
		cgslurmjob.JobArrayMaxTasks = C.GoString(job.array_max_tasks)
		cgslurmjob.JobArrayTaskStr = C.GoString(job.array_task_str)
		cgslurmjob.JobAssocID = C.GoString(job.assoc_id)
		cgslurmjob.JobBatchFeatures = C.GoString(job.batch_features)
		cgslurmjob.JobBatchFlag = C.GoString(job.batch_flag)
		cgslurmjob.JobBatchHost = C.GoString(job.batch_host)
		cgslurmjob.JobBitflags = C.GoString(job.bitflags)
		cgslurmjob.JobBoards_per_node = C.GoString(job.boards_per_node)
		cgslurmjob.Jobburst_buffer = C.GoString(job.burst_buffer)
		cgslurmjob.Jobburst_buffer_state = C.GoString(job.burst_buffer_state)
		cgslurmjob.Jobcluster = C.GoString(job.cluster)
		cgslurmjob.Jobcluster_features = C.GoString(job.cluster_features)
		cgslurmjob.JobCommand = C.GoString(job.command)
		cgslurmjob.JobComment = C.GoString(job.comment)
		cgslurmjob.JobContiguous = C.GoString(job.contiguous)
		cgslurmjob.JobCoreSpec = C.GoString(job.core_spec)
		cgslurmjob.JobCoresPerSocket = C.GoString(job.core_per_socket)
		cgslurmjob.JobBillableTres = C.GoString(job.billable_tres)
		cgslurmjob.JobCpusPerTask = C.GoString(job.cpus_per_task)
		cgslurmjob.JobCpuFreqMin = C.GoString(job.cpu_freq_min)
		cgslurmjob.JobCpuFreqMax = C.GoString(job.cpu_freq_max)
		cgslurmjob.JobCpuFreqGov = C.GoString(job.cpu_freq_gov)
		cgslurmjob.JobCpusPerTres = C.GoString(job.cpus_per_tres)
		cgslurmjob.JobDeadline = C.GoString(job.deadline)
		cgslurmjob.JobDelayBoot = C.GoString(job.delay_boot)
		cgslurmjob.JobDependency = C.GoString(job.dependency)
		cgslurmjob.JobDerivedEC = C.GoString(job.derived_ec)
		cgslurmjob.JobEligible_time = C.GoString(job.derived_ec)
		cgslurmjob.JobEndTime = C.GoString(job.end_time)
		cgslurmjob.JobExcNodes = C.GoString(job.exc_nodes)
		cgslurmjob.JobExcNodeInx = C.GoString(job.exc_node_inx)
		cgslurmjob.JobExitCode = C.GoString(job.exit_code)
		cgslurmjob.JobFeatures = C.GoString(job.features)
		cgslurmjob.JobFedOriginStr = C.GoString(job.fed_origin_str)
		//cgslurmjob.JobFedOriginStr = C.GoString(job.fed_siblings_active)
		cgslurmjob.JobFedSiblingsActiveStr = C.GoString(job.fed_siblings_active_str)
		cgslurmjob.JobFedSiblingsViableStr = C.GoString(job.fed_siblings_viable_str)
		cgslurmjob.JobGresDetailCnt = C.GoString(job.gres_detail_cnt)
		cgslurmjob.JobGroupID = C.GoString(job.group_id)
		cgslurmjob.JobJobID = C.GoString(job.job_id)
		cgslurmjob.JobJobState = C.GoString(job.job_state)
		cgslurmjob.JobTastSchedEval = C.GoString(job.last_sched_eval)
		cgslurmjob.JobLicenses = C.GoString(job.licenses)
		cgslurmjob.JobMaxCpus = C.GoString(job.max_cpus)
		cgslurmjob.JobMaxNodes = C.GoString(job.max_nodes)
		cgslurmjob.JobMcsLabel = C.GoString(job.mac_label)
		cgslurmjob.JobMemPerTres = C.GoString(job.mem_per_tres)
		cgslurmjob.JobName = C.GoString(job.name)
		cgslurmjob.JobNetwork = C.GoString(job.network)
		cgslurmjob.JobNodes = C.GoString(job.nodes)
		cgslurmjob.JobNice = C.GoString(job.nice)
		cgslurmjob.JobNodeInx = C.GoString(job.node_inx)
		cgslurmjob.JobTasksPerCore = C.GoString(job.ntasks_per_core)
		cgslurmjob.JobTasksPerNode = C.GoString(job.ntask_per_node)
		cgslurmjob.JobTasksPerSocket = C.GoString(job.ntask_per_socket)
		cgslurmjob.JobTasksPerBoard = C.GoString(job.ntask_per_board)
		cgslurmjob.JobMumCpus = C.GoString(job.num_cpus)
		cgslurmjob.JobNumNodes = C.GoString(job.num_nodes)
		cgslurmjob.JobNumTasks = C.GoString(job.num_tasks)
		cgslurmjob.JobPackJobId = C.GoString(job.pack_job_id)
		cgslurmjob.JobPackJobIdSet = C.GoString(job.pack_job_id_set)
		cgslurmjob.JobPackJobOffset = C.GoString(job.pack_job_id_offset)
		cgslurmjob.JobPartition = C.GoString(job.partition)
		cgslurmjob.JobPnMinMemory = C.GoString(job.pn_min_memory)
		cgslurmjob.JobPnMinCpus = C.GoString(job.pn_min_cpus)
		cgslurmjob.JobPnMinTmpDisk = C.GoString(job.pn_min_tmp_disk)
		cgslurmjob.JobPowerFlags = C.GoString(job.power_flags)
		cgslurmjob.JobTreemptTime = C.GoString(job.preempt_time)
		cgslurmjob.JobTreemptableTime = C.GoString(job.preemtable_time)
		cgslurmjob.JobTreSusTime = C.GoString(job.pre_sus_time)
		cgslurmjob.JobPriority = C.GoString(job.priority)
		cgslurmjob.JobProfile = C.GoString(job.profile)
		cgslurmjob.JobQos = C.GoString(job.qos)
		cgslurmjob.JobReboot = C.GoString(job.reboot)
		cgslurmjob.JobReqNodes = C.GoString(job.req_nodes)
		cgslurmjob.JobReqNodeInx = C.GoString(job.req_node_inx)
		cgslurmjob.JobReqSwitch = C.GoString(job.req_switch)
		cgslurmjob.JobRequeue = C.GoString(job.requeue)
		cgslurmjob.JobResizeTime = C.GoString(job.resize_time)
		cgslurmjob.JobRestartCnt = C.GoString(job.restart_cnt)
		cgslurmjob.JobResvName = C.GoString(job.resv_time)
		cgslurmjob.JobSchedNodes = C.GoString(job.sched_nodes)
		cgslurmjob.JobShared = C.GoString(job.shared)
		cgslurmjob.JobShowFlags = C.GoString(job.show_flags)
		cgslurmjob.JobSiteFactor = C.GoString(job.site_factor)
		cgslurmjob.JobSocketsPerBoard = C.GoString(job.sockets_per_board)
		cgslurmjob.JobSocketsPerNode = C.GoString(job.sockets_per_node)
		cgslurmjob.JobStartTime = C.GoString(job.start_time)
		cgslurmjob.JobStartProtocolVer = C.GoString(job.start_protocol_ver)
		cgslurmjob.JobStateDesc = C.GoString(job.state_desc)
		cgslurmjob.JobStateReason = C.GoString(job.state_reason)
		cgslurmjob.JobStdErr = C.GoString(job.std_err)
		cgslurmjob.JobStdIn = C.GoString(job.std_in)
		cgslurmjob.JobStdOut = C.GoString(job.std_out)
		cgslurmjob.JobSubmitTime = C.GoString(job.submit_time)
		cgslurmjob.JobSuspendTime = C.GoString(job.suspend_time)
		cgslurmjob.JobSystemComment = C.GoString(job.system_comment)
		cgslurmjob.TimeLimit = C.GoString(job.time_limit)
		cgslurmjob.TimeMin = C.GoString(job.time_min)
		cgslurmjob.ThreadsPerCore = C.GoString(job.threads_per_core)
		cgslurmjob.JobRresBind = C.GoString(job.tres_bind)
		cgslurmjob.JobTresFreq = C.GoString(job.tres_freq)
		cgslurmjob.JobTresPerJob = C.GoString(job.tres_job)
		cgslurmjob.JobTresPerNode = C.GoString(job.tres_node)
		cgslurmjob.JobTresPerSocket = C.GoString(job.tres_per_socket)
		cgslurmjob.JobTresPerTask = C.GoString(job.tres_per_task)
		cgslurmjob.JobTresReqStr = C.GoString(job.req_str)
		cgslurmjob.JobTresAllocStr = C.GoString(job.tres_alloc_str)
		cgslurmjob.JobUserID = C.GoString(job.user_id)
		cgslurmjob.JobUserName = C.GoString(job.user_name)
		cgslurmjob.JobWait4switch = C.GoString(job.wait4switch)
		cgslurmjob.JobWckey = C.GoString(job.wckey)
		cgslurmjob.JobWorkDir = C.GoString(job.work_dir)

		slurmJobs = append(slurmJobs, cgslurmjob)
	}
	C.slurm_free_job_info_msg(jobs_info_msg)
	errcode = C.slurm_get_errno()

	err = errors.New(C.GoString(C.slurm_strerror(C.slurm_get_errno())))
	return slurmJobs, errcode, err
}
