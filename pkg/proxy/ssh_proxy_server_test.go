/**
golang 反向代理
*/
package proxy

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"testing"
)

func TestNewSSHProxy(t *testing.T) {
	var hostKey ssh.PublicKey
	target := SSHPrxoyTarget{
		RemoteServer:"10.127.48.13:4134",
		RemoteHostKey:hostKey,
	}
	config := SSHProxyConnectConfig{
		UserName:"liangdu",
		//PassWord:"XHThIKYQ7dXG3jG5q7U=",

	}
	client,session,err :=NewSSHProxy(target,config)
	if err!= nil{
		log.Fatalln(err)
	}
	defer client.Close()
	var b bytes.Buffer

/*	err =session.Run("echo hello wolrd")
		if err!= nil{
			log.Fatalln(err)
		}*/
	session.Stdout = &b
	session.Run("ls")
	fmt.Println(b.String())

}
