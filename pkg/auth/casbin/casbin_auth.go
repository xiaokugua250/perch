package auth

import (
	"github.com/casbin/casbin/v2"
)

type CasbinAcm struct { // access control model
	*casbin.SyncedEnforcer
}

type CasbinSpec struct {
	Subject string
	Domain  string
	Object  string
	Actions []string
}

/**
初始化casbin 配置
*/
func CasbinInit(conf string) (CasbinAcm, error) {
	var (
		Acm CasbinAcm
		err error
	)
	syncEnforcer, err := casbin.NewSyncedEnforcer(conf)
	if err != nil {
		return CasbinAcm{}, err
	}
	Acm.SyncedEnforcer = syncEnforcer

	return Acm, err
}

/**
添加验证策略，策略从配置文件或者数据库中读取
*/
func (casbinAcm *CasbinAcm) CasbinAddPolicies(policies interface{}) (bool, error) {

	return false, nil
}

/**
casbin进行权限验证
*/
func (casbinAcm *CasbinAcm) CasbinAccess(request CasbinSpec) (bool, error) {
	var (
		err error
	)
	passed, err := casbinAcm.Enforcer.Enforce(request)
	if err != nil {
		return false, err
	}
	return passed, err
}
