/**
golang 反向代理
ref: https://go.googlesource.com/crypto/+/master/ssh/client_auth_test.go
*/
package proxy

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	_ "perch/pkg/log"
	"time"
)

type SSHPrxoyTarget struct {
	RemoteServer string `json:"remote_server"`
	RemoteHostKey ssh.PublicKey
}

type SSHProxyConnectConfig struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	IdRSA string `json:"id_rsa"` //私钥
}




/**
针对单个URL的反向代理
 */
func NewSSHProxy(target SSHPrxoyTarget,config SSHProxyConnectConfig)(*ssh.Client,*ssh.Session,error) {
	 conf := &ssh.ClientConfig{
	 	User:config.UserName,

		/*Auth:[]ssh.AuthMethod{
	 		ssh.Password(config.PassWord),

		},*/

		 HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		 Timeout:5*time.Second,
	 }
	 if config.PassWord != ""{
	 	conf.Auth=[]ssh.AuthMethod{ssh.Password(config.PassWord)}
	 }else {
	 	conf.Auth=[]ssh.AuthMethod{ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
			answers = make([]string, len(questions))
			for n, q := range questions {
				fmt.Printf("Got question: %s\n", q)
				answers[n] = "XHThIKYQ7dXG3jG5q7U="
			}



			return answers, nil
		})}
	 }
	 client,err := ssh.Dial("tcp",target.RemoteServer,conf)

	 if err!= nil{
		 return nil,nil, err
	 }
	// defer  client.Close()
	 session ,err := client.NewSession()
	 //defer  session.Close()
	 if err!= nil{
	 	return nil, nil,err
	 }
	 return client, session,nil

}
func SshInteractive(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	answers = make([]string, len(questions))
	for n, q := range questions {
		fmt.Printf("Got question: %s\n", q)
		answers[n] = suitable_answers[pwIdx]
	}
	pwIdx++

	return answers, nil
}


var suitable_answers = []string{"bogus password", "real password"}
var pwIdx = 0
