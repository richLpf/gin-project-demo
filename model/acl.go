package model

import "time"

// Passage model
//缺少字段注释，是否必选，具体的数据类型, 怎么引用公共类型CreatedBy CreateedAt  UpdatedAt
//定义某些字段唯一
//定义合理的字段类型
type Roles struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Role      string    `json:"role"`
	Describe  string    `json:"describe"`
	Operator  string    `json:"operator"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted uint      `json:"is_deleted" omitempty`
}

//Resources model
type Resources struct {
	ID         uint      `json:"id"`
	Namespace  string    `json:"namespace"`
	Category   string    `json:"category"`
	Resource   string    `json:"resource"`
	Properties string    `json:"properties"`
	Name       string    `json:"name"`
	Describe   string    `json:"describe"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDeleted  uint      `json:"is_deleted", omitempty`
}

//RoleResources model
type RoleResources struct {
	ID         uint      `json:"id"`
	Namespace  string    `json:"namespace"`
	RoleID     uint      `json:"role_id"`
	ResourceID uint      `json:"resource_id"`
	Describe   string    `json:"describe"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsDeleted  uint      `json:"is_deleted", omitempty`
}

type RolePermission struct {
	ResourceID uint   `json:"resource_id"`
	Describe   string `json:"describe"`
}

type ReqRoleResources struct {
	Namespace      string           `json:"namespace"`
	RoleID         uint             `json:"role_id"`
	RolePermission []RolePermission `json:"role_permission"`
	CreatedBy      string           `json:"created_by"`
}

// Namespaces modal
type Namespaces struct {
	ID        uint      `json:"id"`
	Namespace string    `json:"namespace"`
	Parent    string    `json:"parent"`
	Name      string    `json:"name"`
	Describe  string    `json:"describe"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted uint      `json:"is_deleted,omitempty"`
}
