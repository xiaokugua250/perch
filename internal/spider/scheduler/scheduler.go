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
