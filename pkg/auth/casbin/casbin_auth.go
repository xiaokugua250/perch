package auth

import (
	

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CasbinAcm struct { // access control model
	*casbin.SyncedEnforcer
}

type CasbinSpecRequest struct {
	Subject string
//	Role    string
	Domain  string
	Object  string
}

// Increase the column size to 512.
type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:32;uniqueIndex:unique_index"`
	V0    string `gorm:"size:32;uniqueIndex:unique_index"`
	V1    string `gorm:"size:32;uniqueIndex:unique_index"`
	V2    string `gorm:"size:32;uniqueIndex:unique_index"`
	V3    string `gorm:"size:32;uniqueIndex:unique_index"`
	V4    string `gorm:"size:32;uniqueIndex:unique_index"`
	V5    string `gorm:"size:32;uniqueIndex:unique_index"`
}
// TableName overrides the table name used by User to `profiles`
func (CasbinRule) TableName() string {
	return "casbin_rule"
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
func (casbinAcm *CasbinAcm) CasbinSetAdapters(adapter persist.Adapter) (bool, error) {

	casbinAcm.SyncedEnforcer.Enforcer.SetAdapter(adapter)
	return false, nil
}

/**
casbin进行权限验证
*/
func (casbinAcm *CasbinAcm) CasbinAccess(request CasbinSpecRequest) (bool, error) {
	var (
		err error
	)
	passed, err := casbinAcm.Enforcer.Enforce(request.Subject, request.Domain, request.Object)
	if err != nil {
		return false, err
	}

	return passed, err
}

/**
casbin进行权限验证
*/
func (casbinAcm *CasbinAcm) CasbinAccessWithDB(db *gorm.DB, request CasbinSpecRequest) (bool, error) {
	var (
		dbAdapter *gormadapter.Adapter
		err       error
	)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	dbAdapter, err = gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	//dbAdapter, err = gormadapter.Newa(db, &CasbinRule{})
	if err != nil {
		return false, err
	}
	casbinAcm.SyncedEnforcer.Enforcer.SetAdapter(dbAdapter)

	//	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", dbAdapter)
	/*
		filter := gormadapter.Filter{
			PType: []string{},
			V0: []string{},
			V1: []string{},
			V2: []string{},
			V3: []string{},
			V4: []string{},
			V5: []string{},
		}*/

	// Load the policy from DB.

	if err = casbinAcm.Enforcer.LoadPolicy(); err != nil {
		return false, err
	}

	passed, err := casbinAcm.Enforcer.Enforce(request.Subject, request.Domain, request.Object)
	if err != nil {
		return false, err
	}
	/*
		if err = casbinAcm.Enforcer.SavePolicy(); err != nil {
			return false, err
		}*/

	return passed, err
}

/**
casbin进行权限验证,添加特定filter过滤
/*
	filter := gormadapter.Filter{
		PType: []string{},
		V0: []string{},
		V1: []string{},
		V2: []string{},
		V3: []string{},
		V4: []string{},
		V5: []string{},
	}*/

func (casbinAcm *CasbinAcm) CasbinSpecAccessWithDB(db *gorm.DB, filter gormadapter.Filter, request CasbinSpecRequest) (bool, error) {
	var (
		dbAdapter *gormadapter.Adapter
		err       error
	)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	dbAdapter, err = gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	//dbAdapter, err = gormadapter.Newa(db, &CasbinRule{})
	if err != nil {
		return false, err
	}
	casbinAcm.SyncedEnforcer.Enforcer.SetAdapter(dbAdapter)

	//	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", dbAdapter)

	// Load the policy from DB.
	if err = casbinAcm.Enforcer.LoadFilteredPolicy(filter); err != nil {
		return false, err
	}

	passed, err := casbinAcm.Enforcer.Enforce(request.Subject, request.Domain, request.Object)
	if err != nil {
		return false, err
	}
	/*
		if err = casbinAcm.Enforcer.SavePolicy(); err != nil {
			return false, err
		}*/

	return passed, err
}
