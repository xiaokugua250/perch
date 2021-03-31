package auth

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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

// Increase the column size to 512.
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
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

/**
casbin进行权限验证
*/
func (casbinAcm *CasbinAcm) CasbinAccessWithDB(db *gorm.DB, request CasbinSpec) (bool, error) {
	var (
		dbAdapter *gormadapter.Adapter
		err       error
	)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	dbAdapter, err = gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	if err != nil {
		return false, err
	}
	casbinAcm.Enforce(dbAdapter)
	//	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", dbAdapter)

	// Load the policy from DB.
	if err = casbinAcm.Enforcer.LoadPolicy(); err != nil {
		return false, err
	}

	passed, err := casbinAcm.Enforcer.Enforce(request)
	if err != nil {
		return false, err
	}
	if err = casbinAcm.Enforcer.SavePolicy(); err != nil {
		return false, err
	}

	return passed, err
}
