/**
并行环境下error处理可参考golang.org/x/sync/errgroup包
另外，并行环境下需要考虑使用context包
*/
package concurrent

import (
	"sync"
)

/**
定义带Mutex锁的资源
对该资源对象的读写时都要记得对MUTEX的锁操作
*/
type SyncResource struct {
	lock     sync.Mutex
	resource interface{}
}

/**
定义带Mutex锁的资源,该资源只执行一次
对该资源对象的读写时都要记得对MUTEX的锁操作
执行一次时 执行once.Do()方法
*/
type SyncOnceResource struct {
	lock     sync.Mutex
	once     sync.Once
	resource interface{}
}

/**
sync 包中sync.Map数据结构来处理并行Map
*/
func CreateConcurMap() {
	//m :=sync.Map{}

}

/**
创建资源池，对资源池中的资源读写是线程安全的
使用场景针对资源初始化比较昂贵的情况
*/
func ResourcePool() sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			//todo  对资源进行POOL处理
			return nil
		},
	}
}
