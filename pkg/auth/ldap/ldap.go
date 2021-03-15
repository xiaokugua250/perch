/**
基于LDAP的身份认证管理
*/
package ldap

import (
	"crypto/tls"
	"errors"
	model "perch/web/model/rbac"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
)

type LdapConfig struct {
	Host       string         //ip或者主机名
	Port       string         //端口
	BaseDn     string         //基础DN
	Attributes LdapAttributes //结果集
}
type LdapAttributes struct {
	UIDKey         int    //ldap中用户uid
	GIDKey         int    //ldap中用户gid
	UserNameKey    string //ldap中的用户名
	PasswdKey      string //密码
	GroupNameKey   string //ldap中用户组
	UserHomeDirKey string //用户家目录
	EmailKey       string //ldap中email的key
	TelephoneKey   int64
}

func init() {

}

func ConnectWithTLS(ldapAddr string, tlsConfig *tls.Config) (*ldap.Conn, error) {

	conn, err := ldap.DialTLS("tcp", ldapAddr, tlsConfig)
	if err != nil {
		return nil, err
	}

	conn.SetTimeout(5 * time.Second)
	return conn, nil
}

func BindWithTLS(conn *ldap.Conn, user string, passwd string) error {
	return conn.Bind(user, passwd)
}

func EntrySearch(conn *ldap.Conn, request *ldap.SearchRequest) (ldap.SearchResult, error) {

	result, err := conn.Search(request)
	if err != nil {
		return ldap.SearchResult{}, err
	}
	return *result, nil
}

func UserSearch(conn *ldap.Conn, searchRequest *ldap.SearchRequest) (model.AuthUser, error) {
	var (
		result         *ldap.SearchResult
		user           model.AuthUser
		ladpAttributes LdapAttributes
		err            error
	)

	result, err = conn.Search(searchRequest)
	if err != nil {
		return user, err
	}

	if len(result.Entries) <= 0 || len(result.Entries) > 1 {
		return user, errors.New("user In ldap is not exists")
	}

	uid, err := strconv.ParseInt(result.Entries[0].GetAttributeValue(strconv.Itoa(ladpAttributes.UIDKey)), 10, 64)
	if err != nil {
		return user, nil
	}
	gid, err := strconv.ParseInt(result.Entries[0].GetAttributeValue(strconv.Itoa(ladpAttributes.GIDKey)), 10, 64)
	if err != nil {
		return user, nil
	}
	user = model.AuthUser{
		ID:            0,
		UserUID:       uid,
		UserGID:       gid,
		UserName:      result.Entries[0].GetAttributeValue(ladpAttributes.UserNameKey),
		UserPasswd:    result.Entries[0].GetAttributeValue(ladpAttributes.PasswdKey),
		UserSalt:      "",
		UserEmail:     result.Entries[0].GetAttributeValue(ladpAttributes.EmailKey),
		UserAvatar:    "",
		UserLastLogin: 0,
		UserStatus:    0,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		Description:   "",
		UserRoles:     nil,
	}

	return user, nil

}

func Close(conn *ldap.Conn) {
	conn.Close()
}

//todo chu
func HandleAuth() {

}
