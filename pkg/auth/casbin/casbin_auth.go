package auth

import (
	"github.com/casbin/casbin/v2"
)


/**
初始化casbin 配置
*/
func CasbinInit(conf string )(enforcer, error){
	var (
		success false
		err error
	)
	enforcer,err := casbin.NewSyncedEnforcer

	return success,err
}


func CasibinEnforcer()

func CasbinEnforeAuth()(bool,error){
	var(
		err error
		AuthPass bool
	)
	enforcer,err:= casbin.
}