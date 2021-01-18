package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net"

	"golang.org/x/crypto/ssh"

	"log"
	//	"sshfortress/util"
)

func main() {

	config := &ssh.ServerConfig{

		// Remove to disable public key auth.
		/*	PublicKeyCallback: func(c ssh.ConnMetadata, pubKey ssh.PublicKey) (*ssh.Permissions, error) {
			if authorizedKeysMap[string(pubKey.Marshal())] {
				return &ssh.Permissions{
					// Record the public key used for authentication.
					Extensions: map[string]string{
						"pubkey-fp": ssh.FingerprintSHA256(pubKey),
					},
				}, nil
			}
			return nil, fmt.Errorf("unknown public key for %q", c.User())
		},*/
		ServerVersion: "SSH-2.0-OWN-SERVER",
	}
	// You can generate a keypair with 'ssh-keygen -t rsa'
	privateBytes, err := ioutil.ReadFile("cmd/gproxy/.id_rsa")

	if err != nil {
		log.Fatal("Failed to load private key (.id_rsa)")
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}
	/*	config.PasswordCallback= func(conn ssh.ConnMetadata, password []byte) (permissions *ssh.Permissions, e error) {
		fmt.Println(conn.User(),string(password))
		if conn.User()== "duliang" && string(password)=="duliang"{
			return nil,nil
		}
		return nil, fmt.Errorf("password rejected for %q", conn.User())

	}*/
	config.KeyboardInteractiveCallback = func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (permissions *ssh.Permissions, e error) {
		loginUser := conn.User()
		ans, err := client("", "", []string{"Server: ", "username:", "password"}, []bool{true, true, true})
		if err != nil {
			log.Println(err)
		}
		fmt.Println(ans, loginUser)

		return nil, nil
	}
	config.PublicKeyCallback = func(conn ssh.ConnMetadata, key ssh.PublicKey) (permissions *ssh.Permissions, e error) {
		var privatekey []byte
		signers, err := ssh.ParsePrivateKey(privatekey)
		if err != nil {
			log.Println(e)
		}
		if string(signers.PublicKey().Marshal()) == string(key.Marshal()) {

		}
		return nil, nil
	}
	config.MaxAuthTries = 5
	config.AuthLogCallback = func(conn ssh.ConnMetadata, method string, err error) {

		fmt.Println("==>", conn.User(), method, err)
	}
	config.AddHostKey(private)
	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		log.Fatalf("Failed to listen on 2022 (%s)", err)
	}
	log.Print("Listening on 2022...")
	nConn, err := listener.Accept()
	if err != nil {
		log.Fatal("failed to accept incoming connection: ", err)
	}
	conn, chans, reqs, err := ssh.NewServerConn(nConn, config)
	if err != nil {
		log.Fatal("failed to handshake: ", err)
	}
	fmt.Println(conn)
	//log.Printf("logged in with key %s", conn.Permissions.Extensions["pubkey-fp"])

	// The incoming Request channel must be serviced.
	go ssh.DiscardRequests(reqs)

	// Service the incoming Channel channel.
	for newChannel := range chans {
		// Channels have a type, depending on the application level
		// protocol intended. In the case of a shell, the type is
		// "session" and ServerShell may be used to present a simple
		// terminal interface.
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Fatalf("Could not accept channel: %v", err)
		}

		// Sessions have out-of-band requests such as "shell",
		// "pty-req" and "env".  Here we handle only the
		// "shell" request.
		go func(in <-chan *ssh.Request) {
			for req := range in {
				req.Reply(req.Type == "shell", nil)
			}
		}(requests)

		term := terminal.NewTerminal(channel, "> ")

		go func() {
			defer channel.Close()
			for {
				line, err := term.ReadLine()
				if err != nil {
					break
				}
				fmt.Println(line)
			}
		}()
	}
}
