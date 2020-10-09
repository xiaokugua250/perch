package system

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/**

系统信号处理函数

//todo 根据真实信号进行真实业务逻辑处理,通常在程序main方法中进行
*/
func SignaleHander() {

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		signalRecived := <-signalChan
		switch signalRecived {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed...")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has beening interruption by control+c...")

			exitChan <- 0
		case syscall.SIGTERM:
			fmt.Println("The kill sigterm has executed for process...")

			exitChan <- 0
		case syscall.SIGQUIT:
			log.Println("kill sigquit was executed for process..")

			exitChan <- 1

		}
	}()

	code := <-exitChan
	os.Exit(code)
}

func ShutDownGracefullyHander(resourceRealseHander func()) error {

	closeChan := make(chan int)
	go func() {
		for {
			select {
			case <-closeChan:
				// close  appliaction
				return
			default:
				//todo continue running

			}
		}
	}()
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	/*	go func() {
		signalRecived := <-signalChan
		switch signalRecived {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed...")
			closeChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has beening interruption by control+c...")

			closeChan  <- 0
		case syscall.SIGTERM:
			fmt.Println("The kill sigterm has executed for process...")

			closeChan  <- 0
		case syscall.SIGQUIT:
			log.Println("kill sigquit was executed for process..")

			closeChan  <- 1

		}
	}()*/

	<-signalChan
	close(closeChan)
	resourceRealseHander()

	return nil
}
