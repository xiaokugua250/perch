

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"net"
	_ "perch/pkg/general/log"
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
		fmt.Printf("%s",Password)

	return nil, nil
}

/**
公钥认证
 */
func SSHDServerPublicKeyAuth(conn ssh.ConnMetadata, pubKey ssh.PublicKey) (*ssh.Permissions, error){

	return nil,nil
}

/**
账户密码认证
 */
func SSHDServerPasswordAuth( conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error){
/*	fmt.Println(conn.User(),string(password))
	if conn.User()== "duliang" && string(password)=="duliang"{
		return nil,nil
	}
	return nil, fmt.Errorf("password rejected for %q", conn.User())*/
	return nil,nil
}

/**
KeyBoard方式认证
 */
func SSHDServerKeyBoardActiveAuth(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error){
/**
loginUser := conn.User()
	ans, err := client("", "", []string{"Server: ","username:","password"}, []bool{true,true,true})
	if err!= nil{
		log.Println(err)
	}
	fmt.Println(ans,loginUser)

	return nil,nil
 */
	return nil,nil
}

//ssh server 配置结构体
type SSHDServer struct {
	ServerPrivateKey string //私钥地址
	ServerHost string // 服务地址
	ServerPort int //监听地址
}

/**
处理ssh connection
 */

func (sshserver *SSHDServer) SSHDServerConnectionHandler(conn *ssh.ServerConn, chs <-chan ssh.NewChannel){
	log.Printf("user %s from  %s login,begin to handle channels",conn.User(),conn.RemoteAddr().String())
	for newChan := range chs{
		if newChan.ChannelType()!="session"{
			newChan.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type %s",newChan.ChannelType()))
			continue
		}
		go sshserver.SSHDServerConnChannelHandler(newChan)
	}
}
// At this point, we have the opportunity to reject the client's
// request for another logical connection
func (sshserver *SSHDServer) SSHDServerConnChannelHandler(chans ssh.NewChannel){
		connection,request,err := chans.Accept()
		if err!= nil{
			log.Error(err)
		}
		//todo 处理connection

}
/**
golang ssh server
 */
func (sshserver *SSHDServer) SSHDServerDaemon(){
	config := &ssh.ServerConfig{
		MaxAuthTries:5,
		NoClientAuth:false, //要求客户端进行验证
		ServerVersion: "SSH-2.0-PERCH-SERVER",
	}
	// You can generate a keypair with 'ssh-keygen -t rsa'
	privateBytes, err := ioutil.ReadFile(sshserver.ServerPrivateKey)
	if err != nil {
		log.Fatal("Failed to load private key in %s",sshserver.ServerPrivateKey)
	}
	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}
	config.AddHostKey(private)
	config.PasswordCallback  =SSHDServerPasswordAuth
	config.PublicKeyCallback=SSHDServerPublicKeyAuth
	config.KeyboardInteractiveCallback=SSHDServerKeyBoardActiveAuth
	config.MaxAuthTries  =5
	config.AuthLogCallback= func(conn ssh.ConnMetadata, method string, err error) {
		log.Println("user:",conn.User(),"login with method:",method,"login error is:",err,"timestamp is:",time.Now().Format(time.RFC3339))
	}

	listener, err := net.Listen("tcp", sshserver.ServerHost+":"+strconv.Itoa(sshserver.ServerPort))
	if err != nil {
		log.Fatalf("Failed to listen on %s,error is %s",sshserver.ServerHost+":"+strconv.Itoa(sshserver.ServerPort), err)
	}
	fmt.Printf("starting sshd daemon in %s ...\n",sshserver.ServerHost+":"+strconv.Itoa(sshserver.ServerPort))
	//log.Printf("starting sshd daemon in %s\n",sshserver.ServerHost+":"+strconv.Itoa(sshserver.ServerPort))
	//todo 处理SSH INCOMMING REQUEST
	for {
		conn,err:= listener.Accept()
		if err!= nil{
			log.Fatalf("failed to handshake with ssh,error is %s",err)
		}
		sessionConn,chans ,reqs,err := ssh.NewServerConn(conn,config)
		if err!= nil{
			log.Errorf("failed to get server connection ,error is %s",err)
			continue
		}
		go ssh.DiscardRequests(reqs)
		go sshserver.SSHDServerConnectionHandler(sessionConn,chans)
	}
}



func main(){

	sshserver := new(SSHDServer)
	sshserver.ServerPrivateKey="cmd/gproxy/.id_rsa"
	sshserver.ServerHost="0.0.0.0"
	sshserver.ServerPort=1244
	sshserver.SSHDServerDaemon()

/*
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
	session.Wait()*/
}




