package applications

const (
	Application_Status_Ready = iota + 1
	Application_Status_Private
	Application_Status_Public

	Instances_Status_Ready = iota + 1
	Instances_Status_Waiting
	Instance_Status_Running
	Instance_Status_Exiting
	Instance_Status_Finished
)

type Applications struct {
	ID                int    `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	Name              string `json:"name" gorm:"column:name;type:varchar(128);not null"`
	Type              string `json:"type" gorm:"column:type;type:tinyint(4);not null"`
	Source            string `json:"source" gorm:"column:source;type:varchar(128);not null"`
	Environment       string `json:"environment" gorm:"column:environment;type:varchar(128);not null"`
	DevelopedBy       string `json:"developed_by" gorm:"column:developed_by;type:varchar(128);not null"`
	ApplicationStatus int    `json:"application_status" gorm:"column:application_status;type:tinyint(4)"`
	CreatedAt         int64  `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	UpdatedAt         int64  `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	Description       string `json:"description" gorm:"column:description;type:varchar(128);not null"`
}

// TableName 数据库表名
func (Applications) TableName() string {
	return "application"
}

type ApplicationInstances struct {
	ID             int    `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	Name           string `json:"name" gorm:"column:name;type:varchar(128);not null"`
	Type           string `json:"type" gorm:"column:type;type:tinyint(4);not null"`
	Source         string `json:"source" gorm:"column:source;type:varchar(128);not null"`
	Environment    string `json:"environment" gorm:"column:environment;type:varchar(128);not null"`
	CreatedBy      string `json:"created_by" gorm:"column:created_by;type:varchar(128);not null"`
	InstanceStatus int    `json:"instance_status" gorm:"column:instances_status;type:tinyint(4)"`
	ApplicationID  int64  `json:"application_id" gorm:"column:application_id;type:int(11)"`
	CreatedAt      int64  `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	UpdatedAt      int64  `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	FinishedAt     int64  `json:"FinishedAt" gorm:"column:FinishedAt;type:bigint(20);not null"`
	Description    string `json:"description" gorm:"column:description;type:varchar(128);not null"`
}

// TableName 数据库表名
func (ApplicationInstances) TableName() string {
	return "application_instances"
}
