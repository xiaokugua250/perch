

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)


// 交互式认证
func interactiveAuth(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
	ans, err := client("", "", []string{"Server: "}, []bool{true})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

		ans, err = client("", "", []string{"password: "}, []bool{false})
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Password := ans[0]

	return nil, nil
}

func main(){

	var authMethods []ssh.AuthMethod
	var passwd string

	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {

		answers = make([]string,len(questions))
		for index,_ques := range questions{
			fmt.Printf("quest is %s\n",_ques)
			answers[index]=passwd
		}
		if len(questions) == 0 {
			return []string{}, nil
		}


		return answers,nil
		//return []string{"XHThIKYQ7dXG3jG5q7U="}, nil
	}
	rsaAuthMethod := func(keyfile string) (ssh.AuthMethod){
		buffer ,err := ioutil.ReadFile(keyfile)
		if err!= nil{
			return nil
		}
		key ,err :=ssh.ParsePrivateKey(buffer)
		if err!= nil{
			return nil
		}

		return ssh.PublicKeys(key)
	}
	authMethods = append(authMethods, ssh.RetryableAuthMethod(ssh.KeyboardInteractive(keyboardInteractiveChallenge),5))
	authMethods = append(authMethods, ssh.RetryableAuthMethod(ssh.Password("XHThIKYQ7dXG3jG5q7U="),5))
	authMethods= append(authMethods,rsaAuthMethod("aaa"))

	config := &ssh.ClientConfig{
		User: "liangdu",
		Auth:authMethods,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "10.127.48.13:4134", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()
	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		panic(err)
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	if err = session.Shell(); err != nil {
		fmt.Printf("Failed to start interactive shell: %s", err)

	}
	session.Wait()
}




