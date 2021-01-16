package model

import "time"

type AuthUser struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	UserUID         int64    `json:"user_uid" gorm:"column:user_uid;type:int(10)"`
	UserGID         int64    `json:"user_gid" gorm:"column:user_gid;type:int(10)"`
	UserName        string    `json:"user_name" gorm:"column:user_name;type:varchar(128);not null"`
	UserPasswd      string    `json:"password" gorm:"column:password;type:varchar(128);not null"`
	UserSalt      string    	`json:"-" gorm:"column:salt;type:varchar(128);not null"`
	UserEmail       string    `json:"email" gorm:"column:email;type:varchar(128);not null"`
	UserAvatar      string    `json:"avatar" gorm:"column:avatar;type:varchar(128);not null"`
	UserLastLogin   int64 `json:"last_login" gorm:"column:last_login;type:bigint(12);not null"`
	UserStatus         int64    `json:"user_status" gorm:"column:user_status;type:int(10)"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	Description string    `json:"description" gorm:"column:description;type:varchar(128);not null"`
	UserRoles []AuthRBACRoles `json:"user_roles" gorm:"-" `
}

// TableName 数据库表名
func (AuthUser) TableName() string {
	return "auth_user"
}

type AuthGroups struct {
	ID int `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	GroupName string `json:"groupname" gorm:"column:groupname;type:varchar(128);not null"`
	GroupCreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	GroupUpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	GroupProfile   string    `json:"group_profile" gorm:"column:profile;type:varchar(128);not null"`
}

// TableName 数据库表名
func (AuthGroups) TableName() string {
	return "auth_group"
}

type AuthRBACRoles struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	RoleName        string    `json:"role_name" gorm:"column:role_name;type:varchar(128);not null"`
	RoleCreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	RoleUpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
}

// TableName 数据库表名
func (AuthRBACRoles) TableName() string {
	return "auth_rbac_roles"
}

type AuthRBACPermissions struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	PermissionName        string    `json:"permission_name" gorm:"column:permission_name;type:varchar(128);not null"`
	PermissionCode      int   `json:"permission_code" gorm:"column:permission_code;type:int(11);not null"`
	PermissionCreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	PermissionUpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`

}

// TableName 数据库表名
func (AuthRBACPermissions) TableName() string {
	return "auth_rbac_permissions"
}

type AuthRBACUserRoles struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	UserID        int    `json:"user_id" gorm:"column:user_id;type:int(10);not null"`
	RoleID         int    `json:"role_id" gorm:"column:role_id;type:int(10);not null"`
}

// TableName 数据库表名
func (AuthRBACUserRoles) TableName() string {
	return "auth_rbac_user_roles"
}

type AuthRBACRolePermissions struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	PermissionID         string    `json:"permission_id" gorm:"column:permission_id;type:int(10);not null"`
	RoleID         int    `json:"role_id" gorm:"column:role_id;type:int(10);not null"`
}

// TableName 数据库表名
func (AuthRBACRolePermissions) TableName() string {
	return "auth_rbac_role_permissions"
}

type AuthRBACGroupRoles struct { //用户组角色
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	GroupID        int    `json:"group_id" gorm:"column:group_id;type:int(10);not null"`
	RoleID         int    `json:"role_id" gorm:"column:role_id;type:int(10);not null"`
}

// TableName 数据库表名
func (AuthRBACGroupRoles) TableName() string {
	return "auth_rbac_group_roles"
}

type AuthUserRoles struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	UserUID         string    `json:"user_uid" gorm:"column:useruid;type:int(10);not null"`
	UserGID         string    `json:"user_gid" gorm:"column:usergid;type:int(10);not null"`
	UserName        string    `json:"user_name" gorm:"column:username;type:varchar(128);not null"`
	UserPasswd      string    `json:"password" gorm:"column:password;type:varchar(128);not null"`
	UserEmail       string    `json:"email" gorm:"column:email;type:varchar(128);not null"`
	UserAvatar      string    `json:"avatar" gorm:"column:avatar;type:varchar(128);not null"`
	UserLastLogin   time.Time `json:"last_login" gorm:"column:lastlogin;type:varchar(128);not null"`
	UserStatus      int       `json:"user_status" gorm:"column:user_status;type:tinyint(4);not null"`
	UserCreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	UserUpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	UserDescription string    `json:"user_description" gorm:"column:description;type:varchar(128);not null"`
}

// TableName 数据库表名
func (AuthUserRoles) TableName() string {
	return "auth_user_roles"
}
