package scheduler

import (
	"os"
	"os/signal"
	"time"
)

//展示time包中ticker 定时任务
func TimeTickerTasks() {

	signchan := make(chan os.Signal, 1)
	signal.Notify(signchan) // with spec signals means recive all signals
	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)
	go func() {
		defer func() {
			stop <- true
		}()

		for {
			select {
			case <-ticker.C:
			//todo tick task
			case <-stop:
				//todo close gorouting
				return

			}
		}
	}()
	<-signchan
	ticker.Stop()
	stop <- true
	// job stopped
}
