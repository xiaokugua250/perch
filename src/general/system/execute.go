package system

import (
	"bytes"
	"os/exec"
)

/**
同步执行


	proc.StdoutPipe()
	proc.StdinPipe()
	proc.StderrPipe()
分别获取子进程的输入、输出、错误输出等
*/
func ExecuteCmdWithParamsSync(cmd string, params string) (error, string) {

	proc := exec.Command(cmd, params)
	result := bytes.NewBuffer([]byte{})
	proc.Stdout = result
	err := proc.Run()
	if err != nil {
		return err, ""
	}
	if proc.ProcessState.Success() {
		return nil, result.String()
	}
	return err, ""
}

/**
异步执行
*/
func ExecuteCmdWithParamsAsync(cmd string, params string) (error, string) {

	proc := exec.Command(cmd, params)
	result := bytes.NewBuffer([]byte{})
	proc.Stdout = result
	err := proc.Start()
	if err != nil {
		return err, ""
	}
	err = proc.Wait()
	if proc.ProcessState.Success() {
		return nil, result.String()
	}
	return err, ""
}
