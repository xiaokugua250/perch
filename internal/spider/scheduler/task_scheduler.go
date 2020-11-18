package scheduler
type Task struct {
}

TaskCh:=make(chan Task, 3)
/**
分发任务
*/
func WorkerTaskDistribute() {

}

func WorkerTaskScheduler(worker nums, tasks []Task) {

	for i := 0; i < workers; i++ {
		go TaskWorker(TaskCh)
	}

	for _,task:= tasks{
		TaskCh <- task
	}

}

func TaskWorker(ch chan Task) {
	for{
		task:=<-ch
		TaskProcess(tash)
	}

}

func TaskProcess(task Task) {

}

/**

是通过以上示例已经可以说明赢者为王这种并发模式的使用。这种并发模式适合多个协程对同一种资源的读取，更概括的讲就是做同一件事情，只要有一个协程干成了就OK了。这种模式的优点主要有两个：1.可以最大程度减少耗时；提高成功率。

func main() {
	txtResult := make(chan string, 5)
	go func() {txtResult <- getTxt("res1.flysnow.org")}()
	go func() {txtResult <- getTxt("res2.flysnow.org")}()
	go func() {txtResult <- getTxt("res3.flysnow.org")}()
	go func() {txtResult <- getTxt("res4.flysnow.org")}()
	go func() {txtResult <- getTxt("res5.flysnow.org")}()
	println(<-txtResult)
}

func getTxt(host string) string{
	//省略网络访问逻辑，直接返回模拟结果
	//http.Get(host+"/1.txt")
	return host+"：模拟结果"
}
*/
TaskResult := make(chan interface)
func WorkerTaskSchedulerWithLongTime(worker nums, tasks []Task) {
	TaskChanWithLong := make(chan Task,workers)

	for i := 0; i < workers; i++ {
		go TaskWorkerHandleLongTime(TaskChanWithLong)
	}

	for _,task:= tasks{
		TaskChanWithLong <- task
	}

	<-TaskChanWithLong

}

func TaskWorkerHandleLongTime(chan Task){
	TaskResult<-Task
}
/**
耗时比较长的任务处理
*/
func TaskProcessLongTime(task Task){

}


