package schemas

import "gorm.io/gorm"

type RoleBinding struct {
	gorm.Model
	Subject  string `gorm:"index:idx_rbac"`
	Resource string `gorm:"index:idx_rbac"`
	Verb     string `gorm:"index:idx_rbac"`
}
