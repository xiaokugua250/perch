package p2p

import (
	"bufio"
	"bytes"
	"errors"

	"time"

	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p-core/network"
	logger "github.com/sirupsen/logrus"
	"io"
	"os"

	"strings"
)

const ()

type StarfsDataStream struct {
	ID         int64         `json:"id"`      //ID
	Channel    string        `json:"channel"` //Channel
	FsOpt      int64         `json:"fs_opt"`  //fs opearation,like list,mv,delete ,copy
	FsParam    []string      `json:"fs_param"`
	FsHandler  string        `json:"fs_handler"`  // fs handler
	FileStream io.ReadWriter `json:"file_stream"` //file stream
	FsContent  []byte        `json:"fs_content"`  //file content
}

type NetworkStream struct {
	Stream network.Stream
}

/**
網絡流處理函數
*/
func NetStreamWithChannelHandler(stream network.Stream) {
	defer stream.Close()
}

/**
網絡流處理函數
*/
func ServerNetworkStreamDefaultHandler(stream network.Stream) {
	//defer stream.Close()
	cmd := make(chan string)
	dataChan := make(chan interface{})
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go ReadFromNetWorkStreamServerSide(rw)
	go WriteToNetWorkStreamServerSide(rw, cmd, dataChan)

}

/**
網絡流處理函數
*/
func ClientNetworkStreamDefaultHandler(stream network.Stream) {
	//	defer stream.Close()
	//	logger.Println("begin to handler network stream....")
	cmdChan := make(chan string)
	dataChan := make(chan bool)
	finishChan := make(chan bool)
	emptyChan := make(chan bool)
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go ReadFromNetWorkStreamClientSide(rw, cmdChan, dataChan, finishChan, emptyChan)
	go WriteToNetWorkStreamClientSide(rw, cmdChan, dataChan, finishChan, emptyChan)

}

//从网络流中读取数据
//todo 在處理流時，需要添加錯誤處理，將報錯信息返回給客戶端
func ReadFromNetWorkStreamServerSide(rw *bufio.ReadWriter) {

	for {
		//	done := make(chan bool)
		str, err := rw.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if str == "" {
			continue
		}

		if str != "\n" {
			fmt.Printf("\x1b[32m%s\x1b[0m >>> ", str)
			cmds := strings.Fields(str)
			var params []string
			//	fmt.Println("cmds from client is",cmds)
			logger.Printf("cmds from client is %s", cmds)
			if len(cmds) > 1 {
				params = cmds[1:]
			}
			if strings.ToLower(cmds[0]) != filesystem.FsGet { //todo support ls、pwd...
				go func() {
					result := filesystem.FileSystemOptHander(cmds[0], params...)
					/*	if result. != nil {
						panic(err)
						rw.WriteString(err.Error())
						rw.Flush()

					}*/
					resultByte, err := json.Marshal(result)
					if err != nil {
						utils.Logger.Error(err)
						panic(err)
						//return
					}
					utils.Logger.Printf("data to client is %#v", string(resultByte))
					//todo 將處理結果返回到客戶端
					rw.WriteString(string(resultByte) + "\n")

					rw.Flush()
					if err != nil {
						panic(err)
					}
				}()

			} else { //對應fsget需要進行文件傳輸  todo 讀取服務端文件將其進行解析，並寫入到文件stream流中
				//	done := make(chan struct{})
				if len(params) <= 0 {
					rw.WriteString(string("please input your command params(get command must with get filename,i.e get fileA)....") + "\n")
					rw.Flush()
				} else {
					go func() { //目前只能處理單個文件，後續需要處理文件夾和多個文件
						fileStat, err := os.Stat(params[0])
						if err != nil { //文件不存在
							if os.IsNotExist(err) {
								return
							}
						}
						if fileStat.IsDir() { //only support file now
							return
						}
						fileToSend, err := os.Open(params[0])
						if err != nil {
							return
						}
						defer func() {
							if err := fileToSend.Close(); err != nil {
								panic(err)
							}
						}()

						fileReader := bufio.NewReader(fileToSend)
						buf := make([]byte, 1024)
						for {
							n, err := fileReader.Read(buf)
							if err != nil {
								if err != io.EOF {
									panic(err)
								}
							}
							if n == 0 {
								break
							}
							//	io.CopyBuffer(rw.Writer,fileReader,buf)
							if _, err := rw.Write(buf[:n]); err != nil {
								panic(err)
							}

							//todo 需要注意写入缓冲区时的时机

							if n < 1024 {
								//	close(done)
								rw.Flush()
								break
							}
						}

						//close(done)

					}()
				}

			}
		} else {
			//rw.WriteString("\n")
			rw.Flush()
		}
		//go WriteToNetWorkStreamServerSide(rw)
	}
}

/**
向網絡流中寫入數據
*/
func WriteToNetWorkStreamServerSide(rw *bufio.ReadWriter, cmdChan chan string, dataChan chan interface{}) {

	stdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}

		_, err = rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
		//	break
	}
}

//从网络流中读取数据
func ReadFromNetWorkStreamClientSide(rw *bufio.ReadWriter, cmdChan chan string, dataChan, finishChan, emptyChan chan bool) {

	for {
		isData, _ := <-dataChan

		if isData {
			cmdStr, _ := <-cmdChan
			params := strings.Fields(cmdStr)[1:]
			var err error
			filename := params[0]

			_, err = os.Stat(filename)

			if err == nil || os.IsExist(err) {

				//return  errors.New(fmt.Sprintf("file %s already exists...\n",filename))
			}
			fileObj, err := os.OpenFile(filename+time.Now().String(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err := fileObj.Close(); err != nil {
					panic(err)
				}
			}()

			//fmt.Println(rw.Re)
			w := bufio.NewWriter(fileObj)
			fmt.Println("negin to write to file...")
			buf := make([]byte, 1024)
			for {
				// read a chunk

				//io.Copy(w,rw.Reader)
				n, err := rw.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}

				// write a chunk
				if _, err := w.Write(buf[:n]); err != nil {
					//		panic(err)
				}

				if n < 1024 {
					w.Flush()
					break
				}
				w.Flush()
				//return err
			}
			if err = w.Flush(); err != nil {
				panic(err)
			}
			//ReadNetWorkStreamToFiles(*rw,params[0])

			<-emptyChan

		} else {

			empty, _ := <-emptyChan
			if !empty {
				str, err := rw.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading from buffer")
					panic(err)
				}

				if str == "" {
					return
				}
				if str != "\n" {
					// Green console colour:    \x1b[32m
					// Reset console colour:    \x1b[0m
					//	fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
					dst := &bytes.Buffer{}
					if err := json.Indent(dst, []byte(str), "", "  "); err != nil {
						panic(err)
					}
					fmt.Printf("\x1b[32m%s\x1b[0m> ", dst.String())
				}
			}

		}

	}

}

//case <-time.After(time.Second):
//default:

//fmt.Println("time out")
//default:

//println(" begin to handle file transfer,")
//	brea
//print("\033[H\033[2J")  //clear screen
//	break
// default:
//	//fmt.Println("wdwdwdwdw")
//	time.Sleep(1*time.Second)
//	break

/*	select {
	case <-dataChan:
		//	path :="c.txt"
		cmdStr := <-cmdChan
		fmt.Println("read cmdstr from chan", cmdStr)
		recivedfile := "recivered" + strings.Fields(cmdStr)[1]

		//todo check file is exists
		fileObj, err := os.OpenFile(recivedfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := fileObj.Close(); err != nil {
				panic(err)
			}
		}()

		//fmt.Println(rw.Re)
		w := bufio.NewWriter(fileObj)
		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := rw.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}

			// write a chunk
			if _, err := w.Write(buf[:n]); err != nil {
				panic(err)
			}
			if err = w.Flush(); err != nil {
				panic(err)
			}
			w.Flush()
			if n < 1024 {
				//break
			}
		}
	default:
		var (
		strResult string
			err       error
		)
		strResult, err = rw.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("====not file result===", strResult)
		strResult = ""
		//print("\033[H\033[2J")  //clear screen
		//	break
		break

		buf := make([]byte, 1024)
		for {
			// read a chunk
			n, err := rw.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			//fmt.Println("===>n",n)
			//fmt.Println(string(buf))

			strResult += string(buf)
			if n < 1024 {
				fmt.Println(strResult)
				strResult = ""
				//print("\033[H\033[2J")  //clear screen
				//	break
				break
			}
		}
	}*/

/**
向網絡流中寫入數據
*/
func WriteToNetWorkStreamClientSide(rw *bufio.ReadWriter, cmdChan chan string, dataChan, finishChan, emptyChan chan bool) {

	stdReader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("> ")

		cmdString, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error reading from stdin,%s", err)
			panic(err)
		}

		if strings.Contains(strings.ToLower(cmdString), "get") {

			dataChan <- true
			cmdChan <- cmdString

			//<-finishChan
		} else {
			dataChan <- false
		}
		if cmdString != "\n" {
			_, err = rw.WriteString(fmt.Sprintf("%s\n", cmdString))
			if err != nil {
				fmt.Errorf("Error writing to buffer,error is %s", err.Error())
				panic(err)
			}
			err = rw.Flush()
			if err != nil {
				fmt.Println("Error flushing buffer")
				panic(err)
			}
			emptyChan <- false

		} else {
			emptyChan <- true

		}

	}

}

//finish,_:= <-finishChan

/**
網絡流處理函數 for files
*/
func NetFileStreamWithChannelHandler(stream network.Stream) {
	defer stream.Close()

}

/**
網絡流處理函數,for files
*/
func NetFileStreamDefaultHandler(stream network.Stream) {
	defer stream.Close()

}
func NetStreamShutdown(stream network.Stream) error {

	return stream.Close()
}

func ReadNetWorkStreamToFiles(rw bufio.ReadWriter, filename string) (err error) {
	_, err = os.Stat(filename)

	if err == nil || os.IsExist(err) {

		return errors.New(fmt.Sprintf("file %s already exists...\n", filename))
	}
	fileObj, err := os.OpenFile(filename+time.Now().String(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fileObj.Close(); err != nil {
			panic(err)
		}
	}()

	//fmt.Println(rw.Re)
	w := bufio.NewWriter(fileObj)
	fmt.Println("negin to write to file...")
	buf := make([]byte, 1024)
	for {
		// read a chunk

		//io.Copy(w,rw.Reader)
		n, err := rw.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		io.CopyN(w, rw, int64(n))
		w.Flush()
		// write a chunk
		//	if _, err := w.Write(buf[:n]); err != nil {
		//		panic(err)
		//	}

		//return err
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}

	return err
}

/**
向網絡流中寫入數據
*/
func WriteNetWorkStreamWithFiles(rw *bufio.ReadWriter) {
	/**
		fileStat, err := os.Stat(transferTarget)
	if err != nil {
		log.Fatalf("=== file %s opens failed,err is :%s ===", transferTarget, err.Error())
	}
	fmt.Printf("=== begin to transfer file %s,file size is %d bytes,file is dir %t=== \n", fileStat.Name(), fileStat.Size(), fileStat.IsDir())
	fileObj, err := os.Open(transferTarget)
	if err != nil {
		log.Fatalf("=== file %s opens failed,err is :%s ===", transferTarget, err.Error())
	}
	defer func() {
		if err := fileObj.Close(); err != nil {
			panic(err)
		}
	}()

	fileReader := bufio.NewReader(fileObj)

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fileReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := rw.Write(buf[:n]); err != nil {
			panic(err)
		}

		//todo 需要注意写入缓冲区时的时机
		if err = rw.Flush(); err != nil {
			panic(err)
		}

	}
	*/
}

/**
對網絡流的處理操作，
主要用於處理獲取請求中的文件請求，在接受到文件請求後將文件轉換爲流進行處理
ref https://stackoverflow.com/questions/55340599/how-to-handle-buffered-read-write-streams-to-peers-in-golang-using-libp2p
*/
func ReadStreamData(rw *bufio.ReadWriter) {
	for {
		// read bytes until new line
		msg, err := rw.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			continue
		}

		// get the id
		id := int64(binary.LittleEndian.Uint64(msg[0:8]))

		// get the content, last index is len(msg)-1 to remove the new line char
		content := string(msg[8 : len(msg)-1])

		if content != "" {
			// we print [message ID] content
			fmt.Printf("[%d] %s", id, content)
		}

		// here you could parse your message
		// and prepare a response
		//	response, err := prepareResponse(content)
		if err != nil {
			fmt.Println("Err while preparing response: ", err)
			continue
		}

		//	if err := s.sendMsg(rw, id, response); err != nil {
		fmt.Println("Err while sending response: ", err)
		continue
	}

}

func WriteStreamData(rw *bufio.ReadWriter, id int64, content []byte) error {
	// allocate our slice of bytes with the correct size 4 + size of the message + 1
	var err error
	msg := make([]byte, 4+len(content)+1)

	// write id
	binary.LittleEndian.PutUint64(msg, uint64(id))

	// add content to msg
	copy(msg[13:], content)

	// add new line at the end
	msg[len(msg)-1] = '\n'

	// write msg to stream
	_, err = rw.Write(msg)
	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}
	err = rw.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer")
		return err
	}
	return nil
}

//从网络流中中读取数据写入到本地文件
func ReadP2PDataFile(rw *bufio.ReadWriter, isFileTransfer bool, transferTarget string) {
	fileObj, err := os.Create(transferTarget)
	if err != nil {
		logger.Printf("===create file %s errors ,error  is %s ", transferTarget, err.Error())
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fileObj.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fileObj)
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := rw.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
			//return
		}
		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
		if err = w.Flush(); err != nil {
			panic(err)
		}
		if n < 1024 {
			break
		}
	}
	fmt.Printf("=== \x1b[32m\x1b[0m> file %s transfer reciverd finished ===\n", transferTarget)
}

//fileName := make(chan string)
// 打开本地文件写入到网络流中
func WriteP2PDataFile(rw *bufio.ReadWriter, isFileTransfer bool, transferTarget string) {
	//	transferTarget <-  fil
	fileStat, err := os.Stat(transferTarget)
	if err != nil {
		logger.Fatalf("=== file %s opens failed,err is :%s ===", transferTarget, err.Error())
	}
	fmt.Printf("=== begin to transfer file %s,file size is %d bytes,file is dir %t=== \n", fileStat.Name(), fileStat.Size(), fileStat.IsDir())
	fileObj, err := os.Open(transferTarget)
	if err != nil {
		logger.Fatalf("=== file %s opens failed,err is :%s ===", transferTarget, err.Error())
	}
	defer func() {
		if err := fileObj.Close(); err != nil {
			panic(err)
		}
	}()

	fileReader := bufio.NewReader(fileObj)

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fileReader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := rw.Write(buf[:n]); err != nil {
			panic(err)
		}

		//todo 需要注意写入缓冲区时的时机
		if err = rw.Flush(); err != nil {
			panic(err)
		}

	}
}

//文件传输Handler
func TransferNetWorkHander(stream network.Stream) {

}

//流传输handler
func StreamNetWorkHander(stream network.Stream) {
	cmdChan := make(chan string)
	dataChan := make(chan interface{})
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go ReadFromNetWorkStreamServerSide(rw)
	//io.Copy(bufio.NewWriter(stream),bufio.NewReader(stream))
	go WriteToNetWorkStreamServerSide(rw, cmdChan, dataChan)

}
