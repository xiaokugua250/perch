/**
基于LDAP的身份认证管理
 */
package ldap

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func  TestConnectWithTLS(t *testing.T) {
	config:= tls.Config{InsecureSkipVerify: true}
	conn,err := ConnectWithTLS("ldap.forumsys.com:389",&config)
	if err!= nil{
		fmt.Printf("error is %s",err.Error())
		return
	}
	conn.Bind("riemann","password")
	fmt.Print(conn)
	//conn.IsClosing()
}
