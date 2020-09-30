/**
分布式爬虫调度器

分布式爬虫，可以从几个层面考虑，分别是代理层面、执行层面和存储层面。
*/
package scheduler

type Scheduler interface {
	ReadyNotifier
	ErrNotifier
	SchedulerInit() error
	SchedulerStart() error
	SchedlulerStop() error
	SchedulerStatus() error
	SchedulerSummary() error
}

//ReadyNotifier ...
type ReadyNotifier interface {
	//	WorkerReady(chan Request)
	WorkerReady(chan interface{})
}

type ErrNotifier interface {
	WorkerErrHander(chan interface{})
}

func SchedulerWithStrategy(strategy func()) {

}
