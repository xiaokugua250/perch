## 常见代码段

1. 优雅关闭进程

```
    ch := make(chan os.Signal, 1)
	//call goroutine啟動http server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("SERVER GG惹:", err)
		}
	}()
	//Notify：將系統訊號轉發至channel
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//阻塞channel
	<-ch
    
    //收到關機訊號時做底下的流程
	fmt.Println("Graceful Shutdown start - 1")
    //透過context.WithTimeout產生一個新的子context，它的特性是有生命週期，這邊是設定10秒
    //只要超過10秒就會自動發出Done()的訊息
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	fmt.Println("Graceful Shutdown start - 2")
    //使用net/http的shutdown進行關閉http server，參數是上面產生的子context，會有生命週期10秒，
    //所以10秒內要把request全都消化掉，如果超時一樣會強制關閉，所以如果http server要處理的是
    //需要花n秒才能處理的request就要把timeout時間拉長一點
	if err := srv.Shutdown(c); err != nil {
		log.Println("srv.Shutdown:", err)
	}
    //使用select去阻塞主線程，當子context發出Done()的訊號才繼續向下走
	select {    
	case <-c.Done():
		fmt.Println("Graceful Shutdown start - 3")
		close(ch)
	}
	fmt.Println("Graceful Shutdown end ")

```

2. 利用pool 机制优化程序性能
2.1 sync.pool例子
```
import (
	"fmt"
	"sync"
	"time"
)

// Pool for our struct A
var pool *sync.Pool

// A dummy struct with a member 
type A struct {
	Name string
}

// Func to init pool
func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Returning new A")
			return new(A)
		},
	}
}

// Main func
func main() {
	// Initializing pool
	initPool()
	// Get hold of instance one
	one := pool.Get().(*A)
	one.Name = "first"
	fmt.Printf("one.Name = %s\n", one.Name)
	// Submit back the instance after using
	pool.Put(one)
	// Now the same instance becomes usable by another routine without allocating it again
}
```
参考
[1].https://www.akshaydeo.com/blog/2017/12/23/How-did-I-improve-latency-by-700-percent-using-syncPool/
[2].https://play.golang.org/p/64SoX7W-x1H
[3].https://github.com/rocketlaunchr/go-pool


3. gourtine 关闭与channel泄漏
3.1 利用done channel关闭管道
```
func worker(messages <-chan int, shutdown <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
            case message := <- messages:
            //Do something useful with message here
            case _ = <- shutdown:
            //We're done!
            return
        }
    }
}
```
Flushing channels 关闭

```
func producer(messages chan<- int, 
shutdown <-chan int,
wg *sync.WaitGroup){
    defer wg.Done()
    for {
        select {
        case _ = <-shutdown:
        return
        case time.After(time.Second):
        messages <- 0
        }
    }
}

func consumer(messages <-chan int,
shutdown <-chan int,
wg *sync.WaitGroup){
    defer wg.Done()
    for {
        select {
        case _ = <-shutdown:
        return
        case msg, ok := <- messages:
        if !ok {
            return //Channel closed, we're done
        }
        //Do something with message

    }
}

```

4 context超时关闭
4.1 context超时控制
```

func hardWork(job interface{}) error {
    time.Sleep(time.Second * 10)
    return nil
}

func requestWork(ctx context.Context, job interface{}) error {
    ctx, cancel := contextx.ShrinkDeadline(ctx, time.Second*2)
    defer cancel()

/**
这样就可以让 done <- hardWork(job) 不管在是否超时都能写入而不卡住 goroutine。此时可能有人会问如果这时写入一个已经没 goroutine 接收的 channel 会不会有问题，在 Go 里面 channel 不像我们常见的文件描述符一样，不是必须关闭的，只是个对象而已，close(channel) 只是用来告诉接收者没有东西要写了，没有其它用途。
*/
    done := make(chan error, 1) //
    panicChan := make(chan interface{}, 1)
    go func() {
        defer func() {
            if p := recover(); p != nil {
                panicChan <- p
            }
        }()

        done <- hardWork(job)
    }()

    select {
    case err := <-done:
        return err
    case p := <-panicChan:
        panic(p)
    case <-ctx.Done():
        return ctx.Err()
    }
}

```
参考
[1].http://www.hydrogen18.com/blog/stopping-it-all-in-go.html

