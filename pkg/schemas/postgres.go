package schemas

import "gorm.io/gorm"

type PostgresRoleBinding struct {
	gorm.Model
	Subject  string `gorm:"uniqueIndex:idx_rbac"`
	Resource string `gorm:"uniqueIndex:idx_rbac"`
	Verb     string `gorm:"uniqueIndex:idx_rbac"`
}

func (PostgresRoleBinding) TableName() string {
	return "role_bindings"
}
