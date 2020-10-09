package system

import (
	"io"
	"os/user"

	"sync"
)

type SyncWriter struct {
	syncMux sync.Mutex
	Writer  io.Writer
}

func (syncWriter *SyncWriter) Write(b []byte) (n int, err error) {
	syncWriter.syncMux.Lock()
	defer syncWriter.syncMux.Unlock()
	return syncWriter.Writer.Write(b)
}

/**
多线程写文件
*/
func MutiGourFileWriter(target string) error {

	//todo use gorutings to write file
	//example
	/*
			targetfile ,err:= os.Create(target)
			if err!= nil{
				return err
			}
		wr := &SyncWriter{sync.Mutex{},targetfile}
		wg := sync.WaitGroup{}


		for _,val := range  dataToWrite {
			wg.Add(1)
			go func( value int) {
				fmt.Fprintf(wr,value)
				wg.Done()
			}(val)
		}
			wg.Wait()
	*/

	return nil

}

/**
获取用户家目录
*/
func GetUserHomeDir() (string, error) {

	user, err := user.Current()

	return user.HomeDir, err

}
